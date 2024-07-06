package main

import (
	"github.com/urfave/cli/v2"
	"jobsearch/db"
	"jobsearch/routes"
	"log"
	"os"
)

func main() {
	defer db.GetDB().Close()
	app := &cli.App{
		Name:  "Job Search",
		Usage: "TODO",
		Commands: []*cli.Command{
			{
				Name:    "dev",
				Aliases: []string{"d"},
				Usage:   "Runs the server",
				Action: func(ctx *cli.Context) error {
					routes.Init()
					return nil
				},
			},
			{
				Name:  "installdb",
				Usage: "Installs the database",
				Action: func(ctx *cli.Context) error {
					db.Install()
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
