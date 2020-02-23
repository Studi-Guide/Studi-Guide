package models

import (
	"github.com/jmoiron/sqlx"
	"os"
	"studi-guide/pkg/env"
	"testing"
)

var testRooms []Room

func setupTestRoomDbService() (RoomServiceProvider, *sqlx.DB) {
	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e := env.NewEnv()

	testRooms = append(testRooms, Room{Id: 1, Name: "01", Description: "d"})
	testRooms = append(testRooms, Room{Id: 2, Name: "02", Description: "d"})
	testRooms = append(testRooms, Room{Id: 3, Name: "03", Description: "d"})


	db := sqlx.MustConnect(e.DbDriverName(), e.DbDataSource())

	db.Exec(schema)

	insert := "insert into rooms (ID, Name, Description) values(:ID, :Name, :Description)"
	tx := db.MustBegin()
	for _, room := range(testRooms) {
		tx.NamedExec(insert, &room)
	}
	tx.Commit()

	dbService := RoomDbService{db: db, table:"rooms"}

	return &dbService, db
}

func TestNewRoomDbService(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "some_driver")
	os.Setenv("DB_DATA_SOURCE", ":some_source")

	e := env.NewEnv()

	dbService, err := NewRoomDbService(e)
	if err == nil {
		t.Error("expected error; got: ", err)
	}
	if dbService != nil {
		t.Error("expected: ", nil, "; got: ", dbService)
	}

	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e = env.NewEnv()
	dbService, err = NewRoomDbService(e)
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

	compare := func(a []Room, b []Room) (bool) {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range(a) {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	if !compare(testRooms, getRooms) {
		t.Error("expected: ", testRooms, "; got: ", getRooms)
	}

	db.MustExec("drop table rooms")

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

	if room != testRooms[1] {
		t.Error("expected: ", testRooms[1], "; got: ", room)
	}

	room, err = dbService.GetRoom("04")
	if err == nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
	var noneRoom Room
	if room !=  noneRoom{
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