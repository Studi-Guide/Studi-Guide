package models

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"os"
	"path/filepath"
	"studi-guide/pkg/building/db/entitymapper"
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

	rooms, err2 := r.CreateMapItems(items)
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
		Rooms []entitymapper.Room `xml:"Room"`
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

func (r *RoomJsonImporter) CreateMapItems(importItems []ImportMapItems) ([]entitymapper.Room, error) {
	var rooms []entitymapper.Room

	for _, item := range importItems {

		var roomNodes []*navigation.PathNode
		var doors []entitymapper.Door
		for _, importDoor := range item.Doors {

			// extract connected nodes
			var connectedNodes []*navigation.PathNode
			for _, connectedNodeId := range importDoor.PathNode.ConnectedPathNodes {

				// create empty pothnode with ID
				connectedNodes = append(connectedNodes, &navigation.PathNode{
					Id:             connectedNodeId,
					Coordinate:     navigation.Coordinate{},
					Group:          nil,
					ConnectedNodes: nil,
				})
			}

			doors = append(doors, entitymapper.Door{
				Section: entitymapper.Section{
					Id:    0,
					Start: importDoor.Start,
					End:   importDoor.End,
				},

				PathNode: navigation.PathNode{
					Id: importDoor.PathNode.Id,
					Coordinate: navigation.Coordinate{
						X: importDoor.PathNode.X,
						Y: importDoor.PathNode.Y,
						Z: importDoor.PathNode.Z,
					},
					Group:          nil,
					ConnectedNodes: connectedNodes,
				},
			})
		}
		for _, node := range item.PathNodes {

			// extract connected nodes
			var connectedNodes []*navigation.PathNode
			for _, connectedNodeId := range node.ConnectedPathNodes {

				// create empty pothnode with ID
				connectedNodes = append(connectedNodes, &navigation.PathNode{
					Id:             connectedNodeId,
					Coordinate:     navigation.Coordinate{},
					Group:          nil,
					ConnectedNodes: nil,
				})
			}

			roomNodes = append(roomNodes, &navigation.PathNode{
				Id:             node.Id,
				Coordinate:     navigation.Coordinate{X: node.X, Y: node.Y, Z: node.Z},
				Group:          nil,
				ConnectedNodes: connectedNodes,
			})
		}

		var locationNode navigation.PathNode
		if len(item.PathNodes) < 1 {
			locationNode = navigation.PathNode{}
			log.Printf("No pathnode found for room room %s!", item.Name)
		} else {
			locationNode = *roomNodes[0]
		}

		room := entitymapper.Room{
			MapItem: entitymapper.MapItem{
				Color:     item.Color,
				Floor:     item.Floor,
				Sections:  item.Sections,
				Campus:    item.Campus,
				Building:  item.Building,
				Doors:     doors,
				PathNodes: roomNodes,
			},

			// Id should be set be DB
			Location: entitymapper.Location{
				Name:        item.Name,
				Description: item.Description,
				Tags:        item.Tags,
				Floor:       item.Floor,
				PathNode:    locationNode,
				Building:    item.Building,
			},
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}
