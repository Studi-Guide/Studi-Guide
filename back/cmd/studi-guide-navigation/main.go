package main

import (
	"log"
	"net/http"
	"os"
	navigationservice "studi-guide/cmd/studi-guide-navigation/server"
	"studi-guide/docs"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/navigation/services"
	"studi-guide/pkg/utils"

	"go.uber.org/dig"
)

func main() {

	log.Print("Starting service")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "StuidGuide API"
	docs.SwaggerInfo.Description = "This is the navigation service of studiGuide."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	container := BuildContainer()

	error := container.Invoke(func(server *navigationservice.NavigationMicroService) {
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
	container.Provide(navigationservice.NewNavigationService)

	// Register entity service for multiple interfaces
	container.Invoke(func(entityserver *entitymapper.EntityMapper) {
		container.Provide(func() services.PathNodeProvider {
			return entityserver
		})
	})

	container.Provide(navigation.NewDijkstraNavigation)
	container.Provide(services.NewNavigationService)
	return container
}
