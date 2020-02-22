package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"strings"
	"studi-guide/pkg/env"
)

type RoomDbService struct {
	db    *sqlx.DB
	table string
}

var schema string = `CREATE TABLE "rooms"(
		"ID"	INTEGER,
		"Name"	TEXT UNIQUE,
		"Description"	TEXT,
		PRIMARY KEY("ID")
		);`

func NewRoomDbService(env *env.Env) (RoomServiceProvider, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	db, err := sqlx.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	_, _ = db.Exec(schema)

	return &RoomDbService{db: db, table: table}, nil
}

func (r *RoomDbService) GetAllRooms() ([]Room, error) {
	var rooms []Room

	err := r.db.Select(&rooms, "select * from "+r.table)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomDbService) GetRoom(name string) (Room, error) {
	var room Room

	err := r.db.Get(&room, "select * from "+r.table+" where Name=$1", name)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *RoomDbService) AddRoom(room Room) error {

	tx := r.db.MustBegin()
	defer tx.Commit()

	_, err := tx.NamedExec("insert into "+r.table+" (ID, Name, Description) values(:ID, :Name, :Description)", &room)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoomDbService) AddRooms(rooms []Room) error {
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
