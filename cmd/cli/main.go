package main

import (
	"log"
	"os"

	"github.com/on3dd/wot-nicknames-scrapper/cmd/cli/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "wot-nicknames-scrapper",
		Commands: []*cli.Command{
			{
				Name:   "parse",
				Usage:  "Parse nicknames from website and write them to \"/assets/data.csv\"",
				Action: actions.Parse,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Value:   "./assets/data.csv",
						Usage:   "Path to the dataset file",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
