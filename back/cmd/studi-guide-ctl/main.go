package main

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
	"log"
	"os"
	"studi-guide/cmd/studi-guide-ctl/ctl"
	"studi-guide/pkg/config"
	"studi-guide/pkg/env"
	"studi-guide/pkg/roomcontroller/models"
)

func main() {
	container := BuildContainer()
	container.Invoke(func(cli *cli.App) {
		err := cli.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(env.NewEnv)
	container.Provide(env.NewArgs)
	container.Provide(config.NewConfig)
	container.Provide(models.NewRoomDbService)
	container.Provide(ctl.StudiGuideCtlCli)
	return container
}
