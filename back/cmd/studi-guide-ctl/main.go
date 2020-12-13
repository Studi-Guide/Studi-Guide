package main

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
	"log"
	"os"
	"studi-guide/cmd/studi-guide-ctl/ctl"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/env"
)

func main() {
	container := BuildContainer()
	container.Invoke(func(cli *cli.App) {
		err := cli.Run(os.Args)
		if err != nil {
			log.Fatal("cli.Run error", err)
		}
	})
}

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(env.NewEnv)
	container.Provide(entitymapper.NewEntityMapper)
	container.Provide(ctl.StudiGuideCtlCli)
	return container
}
