package models

type Room struct {
	Id          int               `json:"id" xml:"id" db:"ID"`
	Name        string            `json:"name" xml:"name" db:"Name"`
	Description string            `json:"description" xml:"description" db:"Description"`
	Alias       []string          `json:"alias" xml:"alias" db:"alias"`
	Doors       []Door            `json:"doors" xml:"doors" db:"doors"`
	Color       string            `json:"color" xml:"color" db:"color"`
	Sections    []SectionProvider `json:"sections" xml:"sections" db:"sections"`
}

type RoomServiceProvider interface {
	GetAllRooms() ([]Room, error)
	GetRoom(name string) (Room, error)
	AddRoom(room Room) error
	AddRooms(rooms []Room) error
}
