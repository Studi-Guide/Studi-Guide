package models

import (
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
)

type ImportDoor struct {
	Start	 	navigation.Coordinate
	End 		navigation.Coordinate
	PathNode 	ImportPathNode
}


type ImportPathNode struct {
	Id 				int
	X 				int
	Y 				int
	Z 				int
	ConnectedPathNodes 	[]int
}

type ImportMapItems struct {
	Name        string
	Description string
	Tags        []string
	Doors       []ImportDoor
	Color       string
	Floor 		string
	Sections    []entityservice.Section
	Campus      string
	Building 	string
	PathNodes 	[]ImportPathNode
}

