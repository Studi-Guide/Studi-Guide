package models

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path/filepath"
)

type RoomImporter interface {
	RunImport() error
}

type RoomJsonImporter struct {
	dbService RoomServiceProvider
	file      string
}

func (r *RoomJsonImporter) RunImport() error {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var rooms []Room
	err = json.NewDecoder(file).Decode(&rooms)
	if err != nil {
		return err
	}

	return r.dbService.AddRooms(rooms)
}

type RoomXmlImporter struct {
	dbService RoomServiceProvider
	file      string
}

func (r *RoomXmlImporter) RunImport() error {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var rooms struct {
		Rooms []Room `xml:"Room"`
	}
	err = xml.NewDecoder(file).Decode(&rooms)
	if err != nil {
		return err
	}

	return r.dbService.AddRooms(rooms.Rooms)
}

func NewRoomImporter(file string, dbService RoomServiceProvider) (RoomImporter, error) {
	var i RoomImporter = nil
	ext := filepath.Ext(file)
	if ext == ".xml" {
		i = &RoomXmlImporter{dbService: dbService, file: file}
	} else if ext == ".json" {
		i = &RoomJsonImporter{dbService: dbService, file: file}
	} else {
		return nil, errors.New("Unknown extension")
	}

	return i, nil
}
