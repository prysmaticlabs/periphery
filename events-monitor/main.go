package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	types "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/io/file"
	"github.com/prysmaticlabs/prysm/v3/proto/eth/service"
	v1 "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	prefixed "github.com/prysmaticlabs/prysm/v3/runtime/logging/logrus-prefixed-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

const emailSlotsPerReorg = 32 // If we receive a reorg event, for the next 32 slots, email a forkchoice dump.

var (
	forkchoiceDebugMethod = "/eth/v1/debug/forkchoice"
	monitorFlags          = struct {
		beaconEndpoint   string
		httpEndpoint     string
		reorgDepth       uint64
		useSendgrid      bool
		sendTo           cli.StringSlice
		sendFrom         string
		smtpHost         string
		smtpPort         string
		smtpPasswordFile string
		smtpUsername     string
		topics           cli.StringSlice
	}{}
)

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
}

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
			&cli.StringSliceFlag{
				Name:        "topics",
				Destination: &monitorFlags.topics,
				Usage:       "List of event topics to subscribe to. Supported: head, chain_reorg",
				Required:    true,
			},
			&cli.Uint64Flag{
				Name:        "on-reorg-depth",
				Destination: &monitorFlags.reorgDepth,
				Value:       3,
				Usage:       "Notify via email only when a chain reorg of a specified depth is detected",
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
		},
		Action: func(cliCtx *cli.Context) error {
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

	// If we receive a reorg, we keep track of it in a special struct, and for the next EMAIL_SLOTS_PER_REORG
	// slots, we then we send an email with a forkchoice dump. Then, we reset this struct to default values.
	reorgDetected := &reorgDetectedMetadata{}
	_ = reorgDetected

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

			// Only notify if user-specified value > reorg depth.
			if monitorFlags.reorgDepth > ev.Depth {
				continue
			}

			reorgDetected = &reorgDetectedMetadata{
				hadReorg:  true,
				reorgSlot: ev.Slot,
			}

			forkchoiceDump, err := getForkchoiceDump(monitorFlags.httpEndpoint)
			if err != nil {
				log.WithError(err).Error("Could not get forkchoice dump data")
				continue
			}
			evData, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "chain_reorg", evData, forkchoiceDump); err != nil {
				log.WithError(err).Error("Could not send chain_reorg event as email attachment")
			}
		}
	}
}

func getForkchoiceDump(endpoint string) ([]byte, error) {
	var forkchoiceDump map[string]interface{}
	resp, err := http.Get(endpoint + forkchoiceDebugMethod)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&forkchoiceDump); err != nil {
		return nil, err
	}
	return json.Marshal(forkchoiceDump)
}

func sendJSONEmail(sender emailSender, eventName string, eventJSON []byte, extraData []byte) error {
	timeNow := time.Now()
	m := newEmailMessage(
		fmt.Sprintf("New %s event detected", eventName),
		fmt.Sprintf("Detected %s event at %v with data %s\n", eventName, time.Now(), string(eventJSON)),
	)
	m.from = monitorFlags.sendFrom
	m.to = monitorFlags.sendTo.Value()
	if extraData != nil {
		fileName := fmt.Sprintf("%s-%d.json", eventName, timeNow.Unix())
		m.attachFileBytes(fileName, extraData)
	}
	return sender.send(m)
}
