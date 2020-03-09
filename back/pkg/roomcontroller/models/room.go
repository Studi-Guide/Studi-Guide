package models

type Room struct {
	Id          int
	Name        string
	Description string
	Alias       []string
	Doors       []Door
	Color       string
	Sections    []Sequence
	Floor		int
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
}
