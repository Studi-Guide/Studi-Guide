package models

type MapItem struct {
	Name        string
	Description string
	Tags        []string
	Doors       []Door
	Color       string
	Floor 		int 		`json:"-"`
	Sections    []Section
	Campus      string
	Building 	string
}
