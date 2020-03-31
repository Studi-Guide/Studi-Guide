package models

import "studi-guide/pkg/navigation"

type Room struct {
	Id          int
	MapItem 	MapItem
	PathNode    navigation.PathNode
}

type ConnectorSpace struct{
	Id          int
	MapItem 	MapItem
	PathNodes   []navigation.PathNode
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
	GetAllPathNodes() ([]navigation.PathNode, error)
}
