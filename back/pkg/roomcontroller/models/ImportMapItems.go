package models

type ImportDoor struct {
	Section 	Section
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
	Floor 		int
	Sections    []Section
	Campus      string
	Building 	string
	PathNodes 	[]ImportPathNode
}

