# Go HTTP Example

This go example uses mux of the gorilla webkit.

## Set Up

 - execute `go mod download`

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

