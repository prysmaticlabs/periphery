package main

import (
	"context"
	"os"

	"github.com/prysmaticlabs/prysm/v3/proto/eth/service"
	v1 "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	prefixed "github.com/prysmaticlabs/prysm/v3/runtime/logging/logrus-prefixed-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	monitorFlags = struct {
		beaconEndpoint string
		topics         cli.StringSlice
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
		Name:  "events-monitor",
		Usage: "Ethereum beacon node events monitoring tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "grpc-endpoint",
				Destination: &monitorFlags.beaconEndpoint,
				Value:       "localhost:4000",
				Usage:       "gRPC endpoint for an Ethereum beacon node",
			},
			&cli.StringSliceFlag{
				Name:        "topics",
				Destination: &monitorFlags.topics,
				Usage:       "List of event topics to subscribe to",
				Required:    true,
			},
		},
		Action: func(cliCtx *cli.Context) error {
			return monitorEvents(cliCtx.Context)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func monitorEvents(ctx context.Context) error {
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
		case "block":
			ev := &v1.EventBlock{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("block", ev).Info("Received event")
		case "finalized_checkpoint":
			ev := &v1.EventFinalizedCheckpoint{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("finalized_checkpoint", ev).Info("Received event")
		case "reorg":
			ev := &v1.EventChainReorg{}
			if err := data.Data.UnmarshalTo(ev); err != nil {
				return err
			}
			log.WithField("chain_reorg", ev).Info("Received event")
		}
	}
}
