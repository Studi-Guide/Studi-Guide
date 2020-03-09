package models

import (
	"context"
	"database/sql"
	"fmt"
	fbsql "github.com/facebookincubator/ent/dialect/sql"
	"log"
	"os"
	"reflect"
	"studi-guide/ent"
	"studi-guide/pkg/env"
	"testing"
)

var testRooms []Room

func setupTestRoomDbService() (RoomServiceProvider, *sql.DB) {
	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e := env.NewEnv()

	testRooms = append(testRooms, Room{Id: 1, Name: "01", Description: "d"})
	testRooms = append(testRooms, Room{Id: 2, Name: "02", Description: "d"})
	testRooms = append(testRooms, Room{Id: 3, Name: "03", Description: "d"})

	drv, err := fbsql.Open(e.DbDriverName(), "file:"+e.DbDataSource()+"?_fk=1")
	if err != nil {
		return nil, nil
	}

	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// run the auto migration tool.
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}



	for _, room := range testRooms {
		client.Room.Create().
			SetFloor(room.Floor).
			SetName(room.Name).
			SetDescription(room.Description).
			Save(ctx)
	}

	dbService := RoomEntityService{client: client, table: "", context: ctx}

	return &dbService, drv.DB()
}

func TestNewRoomDbService(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "some_driver")
	os.Setenv("DB_DATA_SOURCE", ":some_source")

	e := env.NewEnv()

	dbService, err := NewRoomEntityService(e)
	if err == nil {
		t.Error("expected error; got: ", err)
	}
	if dbService != nil {
		t.Error("expected: ", nil, "; got: ", dbService)
	}

	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e = env.NewEnv()
	dbService, err = NewRoomEntityService(e)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
	if dbService == nil {
		t.Error("expected dbService; got: ", dbService)
	}

}

func TestGetRoomAllRooms(t *testing.T) {

	dbService, db := setupTestRoomDbService()

	getRooms, err := dbService.GetAllRooms()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []Room, b []Room) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			fmt.Println("comparing index", i)
			if !reflect.DeepEqual(a, b) {
				return false
			}
		}
		return true
	}

	if !compare(testRooms, getRooms) {
		t.Error("expected: ", testRooms, "; got: ", getRooms)
	}

	db.Exec("drop table rooms")

	getRooms, err = dbService.GetAllRooms()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareRooms []Room
	if !compare(compareRooms, getRooms) {
		t.Error("expected: ", compareRooms, "; got: ", getRooms)
	}

}

func TestGetRoom(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	room, err := dbService.GetRoom("02")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testRooms[1], room) {
		t.Error("expected: ", testRooms[1], "; got: ", room)
	}

	room, err = dbService.GetRoom("04")
	if err == nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
	var noneRoom Room
	if !reflect.DeepEqual(room, noneRoom) {
		t.Error("expected: ", noneRoom, "; got: ", room)
	}
}

func TestAddRoom(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	testRoom := Room{Id: 4, Name: "04", Description: "description"}
	err := dbService.AddRoom(testRoom)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	err = dbService.AddRoom(testRoom)
	if err == nil {
		t.Error("expected error; got nil")
	}
}

func TestAddRooms(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	var newRooms []Room
	newRooms = append(newRooms, Room{Id: 4, Name: "04", Description: "d"})
	newRooms = append(newRooms, Room{Id: 4, Name: "04", Description: "d"})
	newRooms = append(newRooms, Room{Id: 5, Name: "05", Description: "d"})

	err := dbService.AddRooms(newRooms)
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	newRooms = newRooms[:0]
	newRooms = append(newRooms, Room{Id: 6, Name: "06", Description: "d"})
	newRooms = append(newRooms, Room{Id: 7, Name: "07", Description: "d"})
	newRooms = append(newRooms, Room{Id: 8, Name: "08", Description: "d"})

	err = dbService.AddRooms(newRooms)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
}
