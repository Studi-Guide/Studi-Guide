package models

import "studi-guide/pkg/navigation"

type MapItem struct {
	Doors       []Door
	Color       string
	Floor 		int
	Sections    []Section
	Campus      string
	Building 	string
	PathNodes   []*navigation.PathNode
}
