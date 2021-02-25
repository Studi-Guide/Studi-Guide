package main

import (
	"log"
	"net/http"
	"os"
	searchnservice "studi-guide/cmd/studi-guide-search/server"
	"studi-guide/docs"
	"studi-guide/pkg/building/campus"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/building/info"
	"studi-guide/pkg/building/location"
	"studi-guide/pkg/building/room/models"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation/services"
	"studi-guide/pkg/utils"

	"go.uber.org/dig"
)

func main() {

	log.Print("Starting service")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "StuidGuide API"
	docs.SwaggerInfo.Description = "This is the search service of studiGuide."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	container := BuildContainer()

	error := container.Invoke(func(server *searchnservice.SearchMicroService) {
		port := ":8080"
		log.Printf("Starting http listener on %s", port)
		log.Fatal(server.Start(port))
	})

	log.Fatal(error)
}

func defaultLogger() *log.Logger {
	return log.New(os.Stdout, log.Prefix(), log.Flags())
}

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(func() utils.HttpClient {
		return http.DefaultClient
	})

	container.Provide(env.NewEnv)
	container.Provide(defaultLogger)
	container.Provide(entitymapper.NewEntityMapper)

	// Register entity service for multiple interfaces
	container.Invoke(func(entityserver *entitymapper.EntityMapper) {
		container.Provide(func() services.PathNodeProvider {
			return entityserver
		})

		container.Provide(func() models.RoomServiceProvider {
			return entityserver
		})

		container.Provide(func() location.LocationProvider {
			return entityserver
		})

		container.Provide(func() info.BuildingProvider {
			return entityserver
		})

		container.Provide(func() campus.CampusProvider {
			return entityserver
		})
	})

	container.Provide(searchnservice.NewSearchService)

	// Register entity service for multiple interfaces
	container.Invoke(func(entityserver *entitymapper.EntityMapper) {
		container.Provide(func() services.PathNodeProvider {
			return entityserver
		})
	})

	return container
}
