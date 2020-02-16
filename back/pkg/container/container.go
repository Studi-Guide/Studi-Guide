package container

import (
	"go.uber.org/dig"
	"log"
)

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(NewConfig)
	container.Provide(ConnectDatabase)
	return container
}
