package models

import "studi-guide/pkg/navigation"

type Room struct {
	MapItem
	Id       int
	PathNode navigation.PathNode
}

type ConnectorSpace struct {
	MapItem
	Id        int
	PathNodes []navigation.PathNode
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetAllConnectorSpaces() ([]ConnectorSpace, error)
	GetRoomsFromFloor(floor int) ([]Room, error)
	GetConnectorsFromFloor(floor int) ([]ConnectorSpace, error)
}
