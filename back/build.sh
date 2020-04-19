go mod download
go generate ./ent
go run github.com/swaggo/swag/cmd/swag init -g ./cmd/studi-guide/main.go
go generate ./pkg/map ./pkg/navigation/services ./pkg/location
go build ./cmd/...
