package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/prysmaticlabs/prysm/v3/io/file"
	"github.com/prysmaticlabs/prysm/v3/proto/eth/service"
	v1 "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	prefixed "github.com/prysmaticlabs/prysm/v3/runtime/logging/logrus-prefixed-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	monitorFlags = struct {
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
				Value:       "localhost:3500",
				Usage:       "HTTP standard API endpoint for an Ethereum beacon node",
			},
			&cli.StringSliceFlag{
				Name:        "topics",
				Destination: &monitorFlags.topics,
				Usage:       "List of event topics to subscribe to. Supported: head, block, finalized_checkpoint, chain_reorg",
				Required:    true,
			},
			&cli.Uint64Flag{
				Name:        "on-reorg-depth",
				Destination: &monitorFlags.reorgDepth,
				Value:       2,
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
		case "head":
			ev := &v1.EventHead{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("head", ev).Info("Received event")
			rawEvent, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "head", rawEvent); err != nil {
				log.WithError(err).Error("Could not send head event as email attachment")
			}
		case "block":
			ev := &v1.EventBlock{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("block", ev).Info("Received event")
			rawEvent, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "block", rawEvent); err != nil {
				log.WithError(err).Error("Could not send block event as email attachment")
			}
		case "finalized_checkpoint":
			ev := &v1.EventFinalizedCheckpoint{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("finalized_checkpoint", ev).Info("Received event")
			rawEvent, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "finalized_checkpoint", rawEvent); err != nil {
				log.WithError(err).Error("Could not send finalized_checkpoint event as email attachment")
			}
		case "reorg":
			ev := &v1.EventChainReorg{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("chain_reorg", ev).Info("Received event")

			// Only notify if reorg depth is >= a user-specified value.
			if ev.Depth < monitorFlags.reorgDepth {
				continue
			}

			rawEvent, err := json.Marshal(ev)
			if err != nil {
				log.WithError(err).Error("Could marshal event")
				continue
			}
			if err := sendJSONEmail(sender, "chain_reorg", rawEvent); err != nil {
				log.WithError(err).Error("Could not send chain_reorg event as email attachment")
			}
		}
	}
}

func sendJSONEmail(sender emailSender, eventName string, data []byte) error {
	timeNow := time.Now()
	m := newEmailMessage(
		fmt.Sprintf("New %s event detected", eventName),
		fmt.Sprintf("Detected %s event at %v\n", eventName, time.Now()),
	)
	m.from = monitorFlags.sendFrom
	m.to = monitorFlags.sendTo.Value()
	fileName := fmt.Sprintf("%s-%d.json", eventName, timeNow.Unix())
	m.attachFileBytes(fileName, data)
	return sender.send(m)
}
