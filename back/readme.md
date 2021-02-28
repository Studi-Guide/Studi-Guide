# Studi Guide Go Backend

This is the backend of the Studi Guide server.

## Set Up

 - execute `go mod download`
 - run `go generate ./...` to generate everything or run the steps below
 - create database schema `go generate ./ent`
 - generate swagger docs `go run github.com/swaggo/swag/cmd/swag init -g ./cmd/studi-guide/main.go`
 - optionally generate mocks `go generate ./pkg/map ./pkg/navigation/services ./pkg/location`

## Run
 - import data (rooms) `go run ./cmd/studi-guide-ctl migrate import rooms internal/rooms.json`
 - import data (campus) `go run ./cmd/studi-guide-ctl migrate import campus internal/campus.json`
 - import data (rss feeds) `go run ./cmd/studi-guide-ctl migrate import rssfeed ./internal/rssfeeds.json`
 - verify that no other process runs on port 8080
 - run server `go run ./cmd/studi-guide` with the correct environment variables set
 - or simply run `bash run.sh`

## Environment Variables
The following environment variables must be provided to let the backend do its job correctly. All environment variables except the Graphhopper API key can be set by executing `source env.sh`
 - `GRAPHHOPPER_API_KEY` API key for Graphhopper to calculate routes in Openstreetmap
 - `OPENSTREETMAP_BOUNDS=lat,lng;lat,lng` latitude/longitude bounds for southwest and northeast. This is needed to limit route requests and set the map bounds in the fontend
 - `DEVELOP=TRUE` to enable CORS requests for developing
 - `FRONTEND_PATH=/path/to/frontend` if frontend assets should also be served and are not located in the default location `./ionic`

## Try

Hit http://localhost:8080/swagger/index.html to open the swagger api page

## Swagger API infos:
https://github.com/swaggo/swag#getting-started

## Frontend integration
   - download latest frontend binaries from https://github.com/Studi-Guide/Studi-Guide/actions?query=workflow%3AGo
   - copy files into build outputfolder `./ionic`
   - execute  `go run ./cmd/studi-guide`
   - or build frontend in `front/` with the command `ionic build --engine=browser --prod`
   - execute `FRONTEND_PATH=../front/www go run ./cmd/studi-guide`
   - hit http://localhost:8080
   
## Run Docker
  - to create the docker: `docker build --rm -f Dockerfile -t studiguide/studiguide_appservice .`
  - to run the docker:  `docker run -it --rm -p 8080:8080 studiguide/studiguide_appservice:latest`

