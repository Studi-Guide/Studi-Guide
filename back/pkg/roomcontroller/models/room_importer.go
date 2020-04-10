package models

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path/filepath"
	"studi-guide/pkg/navigation"
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

	var items []ImportMapItems
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		return err
	}

	rooms, err2 :=  r.CreateMapItems(items)
	if err2 != nil {
		return err2
	}

	err = r.dbService.AddRooms(rooms)
	return err
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

func (r *RoomJsonImporter) CreateMapItems (importItems []ImportMapItems ) ([]Room, error) {
	var rooms []Room
	for _, item := range importItems {
		room := Room{
			MapItem: MapItem{
				Name:        item.Name,
				Description: item.Description,
				Tags:        item.Tags,
				Color:       item.Color,
				Floor:       item.Floor,
				Sections:    item.Sections,
				Campus:      item.Campus,
				Building:    item.Building,
			},
			// Id should be set be DB
			PathNodes: []*navigation.PathNode{
				{
					Id: item.PathNodes[0].Id,
					Coordinate: navigation.Coordinate{
						X: item.PathNodes[0].X,
						Y: item.PathNodes[0].Y,
						Z: item.PathNodes[0].Z,
					},
					Group: nil,
				},
			},
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func getPathNodesFromImport(importModes []ImportPathNode) []navigation.PathNode {
	var nodes []navigation.PathNode
	for _, item := range importModes {

		nodes = append(nodes, navigation.PathNode{
			Id:             item.Id,
			Coordinate:    	navigation.Coordinate{
				X: item.X,
				Y: item.Y,
				Z: item.Z,
			},
			Group:          nil,
			ConnectedNodes: nil,
		})
	}

	return nodes
}