package main

import (
	"log"
	"studi-guide/cmd/studi-guide/server"
	"studi-guide/docs"
	"studi-guide/pkg/building/location"
	"studi-guide/pkg/building/room/models"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/navigation/services"

	"go.uber.org/dig"
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
	container.Provide(entityservice.NewEntityService)
	container.Provide(server.NewStudiGuideServer)

	// Register entity service for multiple interfaces
	container.Invoke(func(entityserver *entityservice.EntityService) {
		container.Provide(func() services.LocationProvider {
			return entityserver
		})

		container.Provide(func() models.RoomServiceProvider {
			return entityserver
		})

		container.Provide(func() location.LocationProvider {
			return entityserver
		})
	})

	// container.Provide(container.Provide(func() services.LocationProvider {
	// 	return entityservice.NewEntityService()
	// }))

	container.Provide(navigation.NewDijkstraNavigation)
	container.Provide(services.NewNavigationService)
	return container
}
