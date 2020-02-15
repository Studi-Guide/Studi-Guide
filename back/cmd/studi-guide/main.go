package main

import (
	"log"
	"studi-guide/cmd"
	"studi-guide/cmd/studi-guide/server"
	"studi-guide/docs"
)

func main() {

	log.Print("Starting service")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "StuidGuide API"
	docs.SwaggerInfo.Description = "This is a sample server StudiGuide server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "studiguide.swagger.io"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	env := cmd.GetEnv()
	log.Println(env)

	log.Fatal(server.StudiGuideServer(env))
}
