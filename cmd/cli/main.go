package main

import (
	"log"
	"os"

	"github.com/on3dd/wot-nicknames-scrapper/cmd/cli/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "./assets/data.csv",
				Usage:   "Path to the dataset file",
			},
		},
		Name: "wot-nicknames-scrapper",
		Commands: []*cli.Command{
			{
				Name:   "parse",
				Usage:  "Parse nicknames from website and write them to provided file",
				Action: actions.Parse,
			},
			{
				Name:   "generate",
				Usage:  "Generate nicknames from dataset from provided file",
				Action: actions.Generate,
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "count",
						Aliases: []string{"c"},
						Value:   10,
						Usage:   "Number of nicknames to generate",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
