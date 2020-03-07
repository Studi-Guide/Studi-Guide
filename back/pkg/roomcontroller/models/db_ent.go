package models

import (
	"context"
	"errors"
	"log"
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

func NewRoomeEntityService(env *env.Env) (RoomServiceProvider, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx,  err := openDB(driverName, dataSourceName)

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources:", err)
	}

	if err != nil {
		return nil, err
	}

	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func (r *RoomEntityService) GetAllRooms() ([]Room, error) {

	rooms, err := r.client.Room.Query().All(r.context)
	if err != nil {
		return nil, err
	}

	var mRooms []Room
	for _, room := range(rooms) {
		mRooms = append(mRooms, Room{
			Id:          room.ID,
			Name:        room.Name,
			Description: room.Description,
			Floor:       room.Floor,
		})
	}

	return mRooms, nil
}

func (r *RoomEntityService) GetRoom(name string) (Room, error) {

	room, err := r.client.Room.Query().Where(room.Name(name)).First(r.context)
	if err != nil {
		return Room{}, err
	}

	return Room{
		Id:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		Floor:       room.Floor,
	}, nil
}

func (r *RoomEntityService) AddRoom(room Room) error {

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

func (r *RoomEntityService) AddRooms(rooms []Room) error {
	var errorStr []string
	for _, room := range rooms {
		if err := r.AddRoom(room); err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(room.Id))
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
	client, err := ent.Open(dbDriverName, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, ctx, err
}
