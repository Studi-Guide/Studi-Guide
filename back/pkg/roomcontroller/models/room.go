package models

import "studi-guide/ent"

type Room struct {
	Id          int               `db:"Id"`
	Name        string            `db:"Name"`
	Description string            `db:"Description"`
	Alias       []string          `db:"Alias"`
	Doors       []Door            `db:"Doors"`
	Color       string            `db:"Color"`
	Sections    []SectionProvider `db:"Sections"`
	Floor		int 			`json:"floor" xml:"floor" db:"Floor"`
}

type RoomServiceProvider interface {
	GetAllRooms() ([]*ent.Room, error)
	GetRoom(name string) (*ent.Room, error)
	AddRoom(room ent.Room) error
	AddRooms(rooms []ent.Room) error
}
