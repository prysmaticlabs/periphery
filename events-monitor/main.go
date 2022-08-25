package main

import (
	"log"
	"os"

	"github.com/prysmaticlabs/prysm/v3/proto/eth/service"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	monitorFlags = struct {
		beaconEndpoint string
	}{}
)

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
		},
		Action: func(*cli.Context) error {
			return monitorEvents()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func monitorEvents() error {
	conn, err := grpc.Dial(monitorFlags.beaconEndpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	eventsClient := service.NewEventsClient(conn)
	return nil
}
