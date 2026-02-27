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
					fmt.Println("function call: add")
					// function here
					return nil
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "starts a tracking loop",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("function call: start")
					// function here
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "stops main loop",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("function call: stop")
					// function here
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	var testData []wmtracker.RequestInput

	testData = append(testData, wmtracker.RequestInput{Slug: "vauban_prime_set", WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "protea_prime_set", WTB: true, WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "wisp_prime_set", WTB: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "arcane_hot_shot", Rank: 5, WTB: true, WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "archon_vitality", Rank: 10, WTS: true})
	wmtracker.PrintMutipleTopOrders(&testData)
}
