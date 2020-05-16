package mock

import (
	"errors"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
)

type RoomMockService struct {
	RoomList []entitymapper.Room
}

func NewRoomMockService() *RoomMockService {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, entitymapper.Room{MapItem: entitymapper.MapItem{
		Floor:    "1",
		Building: "main",
	},
		Location: entitymapper.Location{PathNode: navigation.PathNode{
			Id:             0,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
			Name:        "RoomN01",
			Description: "Dummy",
			Floor:       "1",
		}})

	rms.RoomList = append(rms.RoomList, entitymapper.Room{MapItem: entitymapper.MapItem{
		Building: "main",
		Floor:    "2",
	},
		Location: entitymapper.Location{PathNode: navigation.PathNode{
			Id:             3,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
			Name:        "RoomN02",
			Description: "Dummy",
			Floor:       "2",
		}})

	rms.RoomList = append(rms.RoomList, entitymapper.Room{MapItem: entitymapper.MapItem{
		Building: "main",
		Floor:    "1",
	},
		Location: entitymapper.Location{PathNode: navigation.PathNode{
			Id:             2,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
			Name:        "RoomN03",
			Description: "Dummy",
			Floor:       "1",
		}})

	rms.RoomList = append(rms.RoomList, entitymapper.Room{MapItem: entitymapper.MapItem{
		Floor:    "2",
		Building: "main",
	},
		Location: entitymapper.Location{PathNode: navigation.PathNode{
			Id:             1,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
			Name:        "RoomN04",
			Description: "Dummy",
			Floor:       "2",
		}})

	return &rms
}

func (r *RoomMockService) GetAllRooms() ([]entitymapper.Room, error) {
	if r.RoomList == nil {
		return nil, errors.New("no room list initialized")
	}
	return r.RoomList, nil
}

func (r *RoomMockService) GetRoom(name, building, campus string) (entitymapper.Room, error) {

	for _, room := range r.RoomList {
		if room.Name == name {
			return room, nil
		}
	}

	return entitymapper.Room{}, errors.New("no such room")
}

func (r *RoomMockService) AddRoom(room entitymapper.Room) error {
	r.RoomList = append(r.RoomList, room)
	return nil
}

func (r *RoomMockService) AddRooms(rooms []entitymapper.Room) error {
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

func (r *RoomMockService) FilterRooms(floor, name, alias, room, building, campus string) ([]entitymapper.Room, error) {
	if r.RoomList == nil {
		return nil, errors.New("no room list initialized")
	}

	var list []entitymapper.Room
	if len(floor) > 0 {
		for _, room := range r.RoomList {
			if room.Location.Floor == floor && room.MapItem.Floor == floor {
				list = append(list, room)
			}
		}
	} else {
		list = r.RoomList
	}

	if len(building) > 0 {
		var buildingFiltered []entitymapper.Room
		for _, room := range list {
			if room.Location.Building == building {
				buildingFiltered = append(buildingFiltered, room)
			}
		}

		list = buildingFiltered
	}

	return list, nil
}
