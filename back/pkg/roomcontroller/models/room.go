package models

import "studi-guide/pkg/navigation"

type Room struct {
	MapItem
	Location
	Id int
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
	GetAllPathNodes() ([]navigation.PathNode, error)
	FilterRooms(floor, name, alias, room string) ([]Room, error)
}
