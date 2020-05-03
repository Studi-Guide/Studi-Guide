package models

import (
	"studi-guide/pkg/building/db/entitymapper"
)

type RoomServiceProvider interface {
	GetAllRooms() ([]entitymapper.Room, error)
	GetRoom(name, building, campus string) (entitymapper.Room, error)
	AddRoom(room entitymapper.Room) error
	AddRooms(rooms []entitymapper.Room) error
	FilterRooms(floor, name, alias, room, building, campus string) ([]entitymapper.Room, error)
}
