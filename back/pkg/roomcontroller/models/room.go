package models

import "image"

type Room struct {
	Id          int             `json:"id" xml:"id"`
	Name        string          `json:"name" xml:"name"`
	Description string          `json:"description" xml:"description"`
	Coordinates image.Rectangle `json:"coordinates" xml:"coordinates"`
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	QueryRooms(query string) ([]Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
}
