package main

import (
	"go.uber.org/dig"
	"log"
	"studi-guide/cmd/studi-guide/server"
	"studi-guide/docs"
	"studi-guide/pkg/config"
	"studi-guide/pkg/env"
	"studi-guide/pkg/roomcontroller/models"
)

func main() {

	log.Print("Starting service")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "StuidGuide API"
	docs.SwaggerInfo.Description = "This is a sample server StudiGuide server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	container := BuildContainer()

	error := container.Invoke(func(server *server.StudiGuideServer) {
		port := ":8080"
		log.Printf("Starting http listener on %s", port)
		log.Fatal(server.Start(port))
	})

	log.Fatal(error)
}

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(env.NewEnv)
	container.Provide(env.NewArgs)
	container.Provide(config.NewConfig)
	container.Provide(models.NewRoomDbService)
	container.Provide(server.NewStudiGuideServer)
	return container
}
