package models

import "image"

type Room struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Coordinates image.Rectangle `json:"coordinates"`
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	QueryRooms(query string) ([]Room, error)
	AddRoom(room Room) error
}
