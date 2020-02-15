package ctl

import(
	"encoding/json"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"studi-guide/pkg/env"
	"studi-guide/pkg/roomcontroller/models"
)

//StudiGuideCtlCli
// get the CLI to run control commands on the studi guide server
func StudiGuideCtlCli() (*cli.App) {

	app := cli.App{
		Name: 	"studi-guide-ctl",
		Usage: 	"Manage studi-guide server",
		Commands: []*cli.Command {
			{
				Name: "migrate",
				Usage: "run database migrations",
				Subcommands: []*cli.Command{
					{
						Name: "import",
						Usage: "import data from json text file",
						Subcommands: []*cli.Command{
							{
								Name: "rooms",
								Usage: "import room data",
								Action: func(context *cli.Context) error {
									file, err := os.Open(context.Args().First())
									if err != nil {
										return err
									}

									var rooms[] models.Room
									err = json.NewDecoder(file).Decode(&rooms)
									if err != nil {
										return err
									}

									dbService, err := models.NewRoomDbService(env.GetEnv().DbDriverName(), env.GetEnv().DbDataSource(), "rooms")
									if err != nil {
										return nil
									}

									for _, room := range(rooms) {
										if err = dbService.AddRoom(room); err != nil {
											log.Println(err, "room:", room)
										} else {
											log.Println("add room:", room)
										}
									}

									return nil
								},
							},
						},
					},
					{
						Name: "init",
						Usage: "initialize database",
						Action: func(context *cli.Context) error {
							return nil
						},
					},
				},
			},
		},
	}

	return &app
}