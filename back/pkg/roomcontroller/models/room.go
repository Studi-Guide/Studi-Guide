package models

import "image"

type Room struct {
	Id          int             `json:"id" xml:"id" db:"ID"`
	Name        string          `json:"name" xml:"name" db:"Name"`
	Description string          `json:"description" xml:"description" db:"Description"`
	Coordinates image.Rectangle `json:"coordinates" xml:"coordinates"`
	Floor		int 			`json:"floor" xml:"floor" db:"Floor"`
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
}
