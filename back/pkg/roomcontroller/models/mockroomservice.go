package models

import (
	"errors"
	"strconv"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
)

type RoomMockService struct {
	RoomList []entityservice.Room
}

func NewRoomMockService() *RoomMockService {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, entityservice.Room{MapItem: entityservice.MapItem{
		Floor: 1,
		},
		Location: entityservice.Location{PathNode: navigation.PathNode{
			Id:             0,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
			Name:        "RoomN01",
			Description: "Dummy",
		}})

	rms.RoomList = append(rms.RoomList, entityservice.Room{MapItem: entityservice.MapItem{
	},
	Location: entityservice.Location{PathNode:navigation.PathNode{
		Id:             3,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
		},
		Name:        "RoomN02",
		Description: "Dummy",
	}})

	rms.RoomList = append(rms.RoomList, entityservice.Room{MapItem: entityservice.MapItem{
	},
		Location: entityservice.Location{PathNode:navigation.PathNode{
		Id:             2,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
		},
		Name:        "RoomN03",
		Description: "Dummy",
		}})


rms.RoomList = append(rms.RoomList, entityservice.Room{MapItem: entityservice.MapItem{
		Floor: 1,
	},
	Location: entityservice.Location{PathNode:navigation.PathNode{
		Id:             1,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
		},
		Name:        "RoomN04",
		Description: "Dummy",
	}})

	return &rms
}

func (r *RoomMockService) GetAllRooms() ([]entityservice.Room, error) {
	if r.RoomList == nil {
		return nil, errors.New("no room list initialized")
	}
	return r.RoomList, nil
}

func (r *RoomMockService) GetRoom(name string) (entityservice.Room, error) {

	for _, room := range r.RoomList {
		if room.Name == name {
			return room, nil
		}
	}

	return entityservice.Room{}, errors.New("no such room")
}

func (r *RoomMockService) AddRoom(room entityservice.Room) error {
	r.RoomList = append(r.RoomList, room)
	return nil
}

func (r *RoomMockService) AddRooms(rooms []entityservice.Room) error {
	for _, room := range rooms {
		_ = r.AddRoom(room)
	}
	return nil
}

func (r *RoomMockService) GetAllPathNodes() ([]navigation.PathNode, error) {
	var list []navigation.PathNode
	for _, room := range r.RoomList {
		for _, node := range room.PathNodes {
			list = append(list, *node)
		}
	}

	return list, nil
}

func (r *RoomMockService) FilterRooms(floor, name, alias, room string) ([]entityservice.Room, error) {
	if r.RoomList == nil {
		return nil, errors.New("no room list initialized")
	}

	if len(floor) > 0 {
		floorInt, err := strconv.Atoi(floor)
		if err != nil {
			return nil, err
		}

		var list []entityservice.Room
		for _, room := range r.RoomList {
			if room.Floor == floorInt {
				list = append(list, room)
			}
		}

		return list, nil
	}

	return r.RoomList, nil
}