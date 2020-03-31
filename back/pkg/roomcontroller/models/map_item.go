package models

type MapItem struct {
	Name        string
	Description string
	Alias       []string
	Doors       []Door
	Color       string
	Floor 		int 		`json:"-"`
	Sections    []Section
}
