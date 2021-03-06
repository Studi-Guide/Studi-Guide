package ctl

import (
	"github.com/urfave/cli/v2"
	"studi-guide/pkg/building/campus"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/building/location"
	"studi-guide/pkg/building/room/models"
	"studi-guide/pkg/rssfeed"
)

//StudiGuideCtlCli
// get the CLI to run control commands on the studi guide server
func StudiGuideCtlCli(dbService *entitymapper.EntityMapper) *cli.App {

	app := cli.App{
		Name:  "studi-guide-ctl",
		Usage: "Manage studi-guide server",
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "run database migrations",
				Subcommands: []*cli.Command{
					{
						Name:  "import",
						Usage: "import data from json or xml text file",
						Subcommands: []*cli.Command{
							{
								Name:  "rooms",
								Usage: "import room data",
								Action: func(context *cli.Context) error {
									importer, err := models.NewRoomImporter(context.Args().First(), dbService)
									if err != nil {
										return err
									}
									return importer.RunImport()
								},
							},
							{
								Name:  "campus",
								Usage: "import campus data",
								Action: func(context *cli.Context) error {
									importer, err := campus.NewCampusImporter(context.Args().First(), dbService)
									if err != nil {
										return err
									}
									return importer.RunImport()
								},
							},
							{
								Name:  "rssfeed",
								Usage: "import rssfeed data",
								Action: func(context *cli.Context) error {
									importer, err := rssfeed.NewRssFeedImporter(context.Args().First(), dbService)
									if err != nil {
										return err
									}
									return importer.RunImport()
								},
							},
							{
								Name:  "location",
								Usage: "import location data",
								Action: func(context *cli.Context) error {
									importer, err := location.NewLocationImporter(context.Args().First(), dbService)
									if err != nil {
										return err
									}
									return importer.RunImport()
								},
							},
						},
					},
				},
			},
		},
	}

	return &app
}
