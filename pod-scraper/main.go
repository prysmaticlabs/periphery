package main

import (
	"context"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	joonix "github.com/joonix/log"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

var (
	monitorFlags = struct {
		projectId      string
		fluentd        bool
		bucketName     string
		scrapeDir      string
		scrapeInterval time.Duration
	}{}
)

func scrapePodData() error {
	if monitorFlags.fluentd {
		f := joonix.NewFormatter()
		if err := joonix.DisableTimestampFormat(f); err != nil {
			log.Fatal(err)
		}
		log.SetFormatter(f)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ticker := time.NewTicker(monitorFlags.scrapeInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := uploadFilesToBucket(ctx, client); err != nil {
				log.WithError(err).Error("Error uploading to bucket")
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func uploadFilesToBucket(ctx context.Context, client *storage.Client) error {
	files, err := os.ReadDir(monitorFlags.scrapeDir)
	if err != nil {
		return err
	}
	for _, fItem := range files {
		if fItem.IsDir() {
			continue
		}
		if err := uploadFile(ctx, client, fItem.Name()); err != nil {
			log.WithError(err).Error("Error uploading to bucket")
			continue
		}
	}
	return nil
}

func uploadFile(ctx context.Context, client *storage.Client, fName string) error {
	fPath := filepath.Join(monitorFlags.scrapeDir, fName)
	f, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.WithError(err).Error("Could not close file")
		}
	}()
	wc := client.Bucket(monitorFlags.bucketName).Object(fName).NewWriter(ctx)
	if _, err := io.Copy(wc, f); err != nil {
		return err
	}
	return wc.Close()
}

func main() {
	app := &cli.App{
		Name:  "pod-scraper",
		Usage: "Tool which scrapes all files at a directory on a pod and uploads them to a gcp bucket at specified intervals",
		Flags: []cli.Flag{
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
			&cli.StringFlag{
				Name:        "scrape-dir",
				Destination: &monitorFlags.scrapeDir,
				Value:       "",
				Usage:       "Directory to scrape",
			},
			&cli.BoolFlag{
				Name:        "fluentd",
				Destination: &monitorFlags.fluentd,
				Usage:       "Fluentd log formatting",
			},
			&cli.DurationFlag{
				Name:        "scrape-interval",
				Destination: &monitorFlags.scrapeInterval,
				Value:       time.Minute * 5,
				Usage:       "Interval to store forkchoice dumps (default 5m0s)",
			},
		},
		Action: func(cliCtx *cli.Context) error {
			return scrapePodData()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
