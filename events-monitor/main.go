package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	joonix "github.com/joonix/log"
	types "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/io/file"
	"github.com/prysmaticlabs/prysm/v3/proto/eth/service"
	v1 "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

const emailSlotsPerReorg = 32 // If we receive a reorg event, for the next 32 slots, email a forkchoice dump.

var (
	forkchoiceDebugMethod = "/eth/v1/debug/forkchoice"
	monitorFlags          = struct {
		beaconEndpoint     string
		fluentd            bool
		httpEndpoint       string
		storeDumpsInterval time.Duration
		purgeDumpsInterval time.Duration
		useSendgrid        bool
		sendTo             cli.StringSlice
		sendFrom           string
		smtpHost           string
		smtpPort           string
		smtpPasswordFile   string
		smtpUsername       string
		topics             cli.StringSlice
		projectId          string
		bucketName         string
	}{}
)

func main() {
	app := &cli.App{
		Name: "events-monitor",
		Usage: "Ethereum beacon node events monitoring tool with the ability to send emails " +
			"with the event data as JSON attachments",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "grpc-endpoint",
				Destination: &monitorFlags.beaconEndpoint,
				Value:       "localhost:4000",
				Usage:       "gRPC endpoint for an Prysm beacon node",
			},
			&cli.StringFlag{
				Name:        "http-endpoint",
				Destination: &monitorFlags.httpEndpoint,
				Value:       "http://localhost:3500",
				Usage:       "HTTP standard API endpoint for an Ethereum beacon node",
			},
			&cli.DurationFlag{
				Name:        "store-dumps-interval",
				Destination: &monitorFlags.storeDumpsInterval,
				Value:       time.Minute * 5,
				Usage:       "Interval to store forkchoice dumps (default 5m)",
			},
			&cli.DurationFlag{
				Name:        "purge-dumps-after",
				Destination: &monitorFlags.purgeDumpsInterval,
				Value:       time.Hour * 48,
				Usage:       "Interval to purge forkchoice dumps (default 48h)",
			},
			&cli.StringSliceFlag{
				Name:        "topics",
				Destination: &monitorFlags.topics,
				Usage:       "List of event topics to subscribe to. Supported: head, chain_reorg",
				Required:    true,
			},
			&cli.StringSliceFlag{
				Name:        "send-to",
				Destination: &monitorFlags.sendTo,
				Usage:       "Recipient email address for events",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "send-from",
				Destination: &monitorFlags.sendFrom,
				Usage:       "Sender email address for events",
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "sendgrid",
				Destination: &monitorFlags.useSendgrid,
				Usage:       "Whether or not to use sendgrid to send emails. Requires SENDGRID_API_KEY env var set",
			},
			&cli.StringFlag{
				Name:        "smtp-host",
				Destination: &monitorFlags.smtpHost,
				Usage:       "Smtp host for sending emails",
			},
			&cli.StringFlag{
				Name:        "smtp-username",
				Destination: &monitorFlags.smtpUsername,
				Usage:       "Smtp username for sending emails",
			},
			&cli.StringFlag{
				Name:        "smtp-port",
				Destination: &monitorFlags.smtpPort,
				Usage:       "Smtp port for sending emails",
			},
			&cli.StringFlag{
				Name:        "smtp-password-file",
				Destination: &monitorFlags.smtpPasswordFile,
				Usage:       "File path to an smtp password for sending emails",
			},
			&cli.StringFlag{
				Name:        "project-id",
				Destination: &monitorFlags.projectId,
				Value:       "",
				Usage:       "Project id on gcp",
			},
			&cli.StringFlag{
				Name:        "bucket-name",
				Destination: &monitorFlags.bucketName,
				Value:       "",
				Usage:       "Bucket name for gcp uploads",
			},
			&cli.BoolFlag{
				Name:        "fluentd",
				Destination: &monitorFlags.fluentd,
				Usage:       "Fluentd log formatting",
			},
		},
		Action: func(cliCtx *cli.Context) error {
			if monitorFlags.fluentd {
				f := joonix.NewFormatter()
				if err := joonix.DisableTimestampFormat(f); err != nil {
					log.Fatal(err)
				}
				log.SetFormatter(f)
			}
			var sender emailSender
			if monitorFlags.useSendgrid {
				sender = newSendgridSender(os.Getenv("SENDGRID_API_KEY"))
			} else {
				smtpPassword, err := file.ReadFileAsBytes(monitorFlags.smtpPasswordFile)
				if err != nil {
					return err
				}
				pw := strings.TrimSpace(string(smtpPassword))
				auth := smtp.PlainAuth(
					"",
					monitorFlags.smtpUsername,
					pw,
					monitorFlags.smtpHost,
				)
				sender = newBasicSmtpSender(auth, monitorFlags.smtpHost, monitorFlags.smtpPort)
			}
			return monitorEvents(cliCtx.Context, sender)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type reorgDetectedMetadata struct {
	hadReorg  bool
	reorgSlot types.Slot
}

func monitorEvents(ctx context.Context, sender emailSender) error {
	log.Info("Starting reorg monitor")
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	conn, err := grpc.Dial(monitorFlags.beaconEndpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	eventsClient := service.NewEventsClient(conn)
	recv, err := eventsClient.StreamEvents(ctx, &v1.StreamEventsRequest{
		Topics: monitorFlags.topics.Value(),
	})
	if err != nil {
		return err
	}

	go storeForkchoiceDumps(ctx, storageClient)

	for {
		data, err := recv.Recv()
		if err != nil {
			return err
		}
		if data == nil {
			log.Error("Received nil event data from beacon node")
			continue
		}
		switch data.Event {
		case "chain_reorg":
			ev := &v1.EventChainReorg{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("chain_reorg", ev).Info("Received event")
			evData, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "chain_reorg", evData); err != nil {
				log.WithError(err).Error("Could not send chain_reorg event as email attachment")
			}
		}
	}
}

func storeForkchoiceDumps(ctx context.Context, storageClient *storage.Client) {
	log.Info("Now starting goroutine for forkchoice dumps")
	timer := time.NewTicker(monitorFlags.storeDumpsInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			if err := writeForkchoiceDump(ctx, storageClient); err != nil {
				log.WithError(err).Error("Could not write forkchoice dump")
			}
		case <-ctx.Done():
			return
		}
	}
}

func writeForkchoiceDump(ctx context.Context, storageClient *storage.Client) error {
	log.Info("Attempting to write forkchoice dump")
	var forkchoiceDump map[string]interface{}
	resp, err := http.Get(monitorFlags.httpEndpoint + forkchoiceDebugMethod)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.WithError(err).Error("Could not close body")
		}
	}()
	enc, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Infof("Got %s response from debug endpoint", enc)
	if err := json.Unmarshal(enc, &forkchoiceDump); err != nil {
		return err
	}
	fileName := forkchoiceFileName()
	log.WithField("bucket", monitorFlags.bucketName).Infof("Attempting to write %s to cloud bucket", fileName)
	wc := storageClient.Bucket(monitorFlags.bucketName).Object(fileName).NewWriter(ctx)
	defer func() {
		if err := wc.Close(); err != nil {
			log.WithError(err).Error("Could not close")
		}
	}()
	return json.NewEncoder(wc).Encode(forkchoiceDump)
}

func sendJSONEmail(sender emailSender, eventName string, eventJSON []byte) error {
	m := newEmailMessage(
		fmt.Sprintf("New %s event detected", eventName),
		fmt.Sprintf("Detected %s event at %v with data %s\n", eventName, time.Now(), string(eventJSON)),
	)
	m.from = monitorFlags.sendFrom
	m.to = monitorFlags.sendTo.Value()
	return sender.send(m)
}

func forkchoiceFileName() string {
	// Use layout string for time format.
	const layout = "2006-01-02-03:04:05-pm"
	// Place now in the string.
	t := time.Now()
	return "forkchoicedump-" + t.Format(layout) + ".json"
}
