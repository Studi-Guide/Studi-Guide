package models

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type RoomImporter interface {
	RunImport() (error)
}

type RoomJsonImporter struct {
	dbService RoomServiceProvider
	file string
}

func (r *RoomJsonImporter) RunImport() (error) {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var rooms[] Room
	err = json.NewDecoder(file).Decode(&rooms)
	if err != nil {
		return err
	}

	for _, room := range(rooms) {
		if err = r.dbService.AddRoom(room); err != nil {
			log.Println(err, "room:", room)
		} else {
			log.Println("add room:", room)
		}
	}

	return nil
}

func NewRoomImporter(file string, dbService RoomServiceProvider) (RoomImporter, error) {
	ext := filepath.Ext(file)
	if ext == ".xml" {

	} else if ext == ".json" {
		i := RoomJsonImporter{dbService: dbService, file: file}
		return &i, nil
	}

	return nil, errors.New("Unknown extension")
}