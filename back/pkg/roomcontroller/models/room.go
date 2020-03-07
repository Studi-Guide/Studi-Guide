package models

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
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
}
