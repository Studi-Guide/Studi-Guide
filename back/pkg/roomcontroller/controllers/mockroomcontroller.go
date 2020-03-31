package controllers

import (
	"errors"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomMockService struct {
	RoomList []models.Room
	ConnectorList []models.ConnectorSpace
}

func NewRoomMockService() *RoomMockService {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, models.Room{MapItem:models.MapItem{
		Name:        "RoomN01",
		Description: "Dummy",
	}})

	rms.RoomList = append(rms.RoomList, models.Room{MapItem:models.MapItem{
		Name:        "RoomN02",
		Description: "Dummy",
	}})

	rms.RoomList = append(rms.RoomList, models.Room{MapItem:models.MapItem{
		Name:        "RoomN03",
		Description: "Dummy",
	}})

	rms.RoomList = append(rms.RoomList, models.Room{MapItem:models.MapItem{
		Name:        "RoomN04",
		Description: "Dummy",
	}})

	rms.ConnectorList = append(rms.ConnectorList, models.ConnectorSpace{MapItem:models.MapItem{
		Name:        "CorridorN01",
		Description: "Dummy",
	}})

	rms.ConnectorList = append(rms.ConnectorList, models.ConnectorSpace{MapItem:models.MapItem{
		Name:        "CorridorN02",
		Description: "Dummy",
	}})

	rms.ConnectorList = append(rms.ConnectorList, models.ConnectorSpace{MapItem:models.MapItem{
		Name:        "CorridorN03",
		Description: "Dummy",
	}})

	rms.ConnectorList = append(rms.ConnectorList, models.ConnectorSpace{MapItem:models.MapItem{
		Name:        "CorridorN04",
		Description: "Dummy",
	}})

	return &rms
}

func (r *RoomMockService) GetAllRooms() ([]models.Room, error) {
	if r.RoomList == nil {
		return nil, errors.New("no room list initialized")
	}
	return r.RoomList, nil
}

func (r *RoomMockService) GetRoom(name string) (models.Room, error) {

	for _, room := range r.RoomList {
		if room.MapItem.Name == name {
			return room, nil
		}
	}

	return models.Room{}, errors.New("no such room")
}

func (r *RoomMockService) AddRoom(room models.Room) error {
	r.RoomList = append(r.RoomList, room)
	return nil
}

func (r *RoomMockService) AddRooms(rooms []models.Room) error {
	for _, room := range rooms {
		_ = r.AddRoom(room)
	}
	return nil
}

func (r *RoomMockService) GetAllConnectorSpaces() ([]models.ConnectorSpace, error) {
	panic("implement me")
}

func (r *RoomMockService) GetRoomsFromFloor(floor int) ([]models.Room, error) {
	var list []models.Room
	for _, room := range r.RoomList {
		if room.MapItem.Floor == floor {
			list = append(list, room)
		}
	}

	return list, nil
}

func (r *RoomMockService) GetConnectorsFromFloor(floor int) ([]models.ConnectorSpace, error) {
	panic("implement me")
}

func (r *RoomMockService) GetAllPathNodes() ([]navigation.PathNode, error) {
	var list []navigation.PathNode
	for _, room := range r.RoomList {
		list = append(list, room.PathNode)
	}

	return list, nil
}

