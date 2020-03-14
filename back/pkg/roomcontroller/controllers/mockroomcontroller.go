package controllers

import (
	"errors"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomMockService struct {
	RoomList []models.Room
}

func (r *RoomMockService) GetAllPathNodes() ([]navigation.PathNode, error) {
	var list []navigation.PathNode
	for _, room := range r.RoomList {
		list = append(list, room.PathNode)
	}

	return list, nil
}

func NewRoomMockService() *RoomMockService {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN01", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN02", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN03", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN04", Description: "Dummy"})

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
		if room.Name == name {
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
