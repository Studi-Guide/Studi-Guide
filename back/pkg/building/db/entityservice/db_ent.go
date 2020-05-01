package entityservice

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/env"
)

type EntityService struct {
	client  *ent.Client
	context context.Context
	table   string
}

func newEntityService(env *env.Env) (*EntityService, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx, err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	roomCount, _ := client.Room.Query().Count(ctx)
	log.Println("Found number of rooms:", roomCount)
	return &EntityService{client: client, table: table, context: ctx}, nil
}

func NewEntityService(env *env.Env) (*EntityService, error) {
	return newEntityService(env)
}

func openDB(dbDriverName string, dbSourceName string) (*ent.Client, context.Context, error) {
	client, err := ent.Open(dbDriverName, "file:"+dbSourceName+"?cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}
	//defer client.Close()
	// run the auto migration tool.
	ctx := context.Background()

	// SQLite was developed only for testing, and it does not support the incremental updates for tables.
	// https://entgo.io/docs/dialects/#sqlite
	if _, err := os.Stat(dbSourceName); dbDriverName != "sqlite3" || (dbDriverName == "sqlite3" && os.IsNotExist(err)) {
		log.Println("running one time migration")
		if err := client.Schema.Create(ctx); err != nil {
			return nil, nil, err
		}
	}

	return client, ctx, err
}

