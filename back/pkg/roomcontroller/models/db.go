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