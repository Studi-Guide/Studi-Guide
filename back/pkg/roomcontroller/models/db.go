package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type RoomDbService struct {
	db *sql.DB
	table string
}

func NewRoomDbService(driverName, dataSourceName, table string) (*RoomDbService, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if (err != nil) {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	_, _ = db.Exec(`CREATE TABLE "rooms" (
		"ID"	INTEGER,
		"Name"	TEXT UNIQUE,
		"Description"	TEXT,
		PRIMARY KEY("ID")
		);`)

	return &RoomDbService{db: db, table: table}, nil
}

func (r *RoomDbService) GetAllRooms() ([]Room, error) {
	var rooms []Room

	stmt, err := r.db.Prepare("select * from rooms")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var room Room
		err = rows.Scan(&room.Id, &room.Name, &room.Description)
		if err != nil {

		} else {
			rooms = append(rooms, room)
		}
	}

	return rooms, nil
}

func (r *RoomDbService) GetRoom(name string) (Room, error) {
	var room Room

	stmt, err := r.db.Prepare("select ID, Name, Description from " + r.table + " where Name = ?")
	if err != nil {
		return room, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&room.Id, &room.Name, &room.Description)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r* RoomDbService) QueryRooms(query string) ([]Room, error) {
	var rooms []Room

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var room Room
		err = rows.Scan(&room.Id, &room.Name, &room.Description)
		if err != nil {

		} else {
			rooms = append(rooms, room)
		}
	}

	return rooms, nil
}

func (r* RoomDbService) AddRoom(room Room) (error) {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare("insert into rooms(ID, Name, Description) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(room.Id, room.Name, room.Description)
	if err != nil {
		return err
	}

	return nil
}