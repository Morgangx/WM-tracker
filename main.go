package main

import (
	wmtracker "WM_tracker/internal"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	// load any previously-saved items so that separate invocations of the
	// program can share state. if the file doesn't exist we get an empty
	// slice and continue.
	var testData []wmtracker.RequestInput
	if loaded, err := wmtracker.LoadRequests(); err != nil {
		log.Fatalf("failed to load saved requests: %v", err)
	} else if loaded != nil {
		testData = loaded
	}

	loop := true
	loopRunning := false
	refreshSeconds := 20
	exit := make(chan bool)
	cmd := &cli.Command{
		Name:  "WM_tracker",
		Usage: "tracks multiple items from warframe.market and prints out specified orders",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a item to track",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "item",
						Aliases: []string{"i"},
						Validator: func(v string) error {
							if v == "" {
								return fmt.Errorf("Item flag is required")
							}
							return nil
						},
					},
					&cli.Int8Flag{
						Name:    "rank",
						Aliases: []string{"r"},
						Value:   0,
						Validator: func(t int8) error {
							if t > 10 {
								return fmt.Errorf("Maximum rank is 10")
							}
							return nil
						},
					},
					&cli.BoolFlag{
						Name:  "WTS",
						Value: true,
					},
					&cli.BoolFlag{
						Name: "WTB",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					testData = append(testData, wmtracker.RequestInput{
						Slug: cmd.String("item"),
						Rank: cmd.Int8("rank"),
						WTS:  cmd.Bool("WTS"),
						WTB:  cmd.Bool("WTB"),
					})
					wmtracker.PrintMutipleTopOrders(&testData)
					return nil
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "starts a tracking loop",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					loopRunning = true
					go wmtracker.MainLoop(&refreshSeconds, &loop)
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "stops main loop",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					loop = false
					exit <- true
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	if loopRunning {
		<-exit
	}
}
