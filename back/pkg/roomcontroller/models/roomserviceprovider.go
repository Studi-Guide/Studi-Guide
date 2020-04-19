package models

import (
	"studi-guide/pkg/entityservice"
)

type RoomServiceProvider interface {
	GetAllRooms() ([]entityservice.Room, error)
	GetRoom(name string) (entityservice.Room, error)
	AddRoom(room entityservice.Room) error
	AddRooms(rooms []entityservice.Room) error
	FilterRooms(floor, name, alias, room string) ([]entityservice.Room, error)
}
