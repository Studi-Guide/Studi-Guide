package ctl

import (
	"github.com/urfave/cli/v2"
	"studi-guide/pkg/roomcontroller/models"
)

//StudiGuideCtlCli
// get the CLI to run control commands on the studi guide server
func StudiGuideCtlCli(dbService models.RoomServiceProvider) *cli.App {

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
						},
					},
				},
			},
		},
	}

	return &app
}
