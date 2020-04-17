# Go HTTP Example

This go example uses mux of the gorilla webkit.

## Set Up

 - execute `go mod download`
 - create database schema `go generate ./ent`
 - generate docs `go run github.com/swaggo/swag/cmd/swag init -g ./cmd/studi-guide/main.go`
 - optionally generate mocks `go generate ./pkg/map ./pkg/navigation/services ./pkg/location`

## Run

 - verify that no other process runs on port 8080
 - execute `go run ./cmd/studi-guide`

## Try

Hit http://localhost:8080/shoppinglist/index

## Swagger

   - execute `swag init -g cmd/studi-guide/main.go`
   - run the application
   - Hit http://localhost:8080/swagger/index.html to open the swagger api page

API infos:
https://github.com/swaggo/swag#getting-started

## Frontend integration
   - download latest frontend binaries from https://github.com/Studi-Guide/Studi-Guide/actions?query=workflow%3AGo
   - copy files into build outputfolder `./ionic`
   - execute  `go run ./cmd/studi-guide`
   - hit http://localhost:8080
   
 ## Run Docker
  - to create the docker: `docker build --rm -f Dockerfile -t studiguide/studiguide_appservice .`
  - to run the docker:  `docker run -it --rm -p 8080:8080 studiguide/studiguide_appservice:latest`