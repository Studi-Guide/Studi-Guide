package generate

//go:generate go run github.com/facebook/ent/cmd/entc generate ./pkg/building/db/ent/schema
//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./cmd/studi-guide/main.go
