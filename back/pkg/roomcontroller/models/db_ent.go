package models

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/room"
	"studi-guide/pkg/env"
)

type RoomEntityService struct {
	client    *ent.Client
	context   context.Context
	table string
}

func NewRoomEntityService(env *env.Env) (RoomServiceProvider, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx,  err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func (r *RoomEntityService) GetAllRooms() ([]*ent.Room, error) {

	rooms, err := r.client.Room.Query().All(r.context)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomEntityService) GetRoom(name string) (*ent.Room, error) {

	room, err := r.client.Room.Query().Where(room.Name(name)).First(r.context)
	if err != nil {
		return &ent.Room{}, err
	}

	return room, nil
}

func (r *RoomEntityService) AddRoom(room ent.Room) error {

	_, err :=r.client.Room.Create().
		SetFloor(room.Floor).
		SetName(room.Name).
		SetDescription(room.Description).
		Save(r.context)
	// Add other stuff like sequence etc.

	if err != nil {
		return  err
	}

	return nil
}

func (r *RoomEntityService) AddRooms(rooms []ent.Room) error {
	var errorStr []string
	for _, room := range rooms {
		if err := r.AddRoom(room); err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(room.ID))
			log.Println(err, "room:", room)
		} else {
			log.Println("add room:", room)
		}
	}

	if len(errorStr) > 0 {
		return errors.New(strings.Join(errorStr, "; "))
	}
	return nil
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
