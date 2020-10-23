package models

import (
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
)

type ImportDoor struct {
	Start    navigation.Coordinate
	End      navigation.Coordinate
	PathNode ImportPathNode
}

type ImportPathNode struct {
	Id                 int
	X                  int
	Y                  int
	Z                  int
	ConnectedPathNodes []int
}

type ImportMapItems struct {
	Name        string
	Description string
	Tags        []string
	Doors       []ImportDoor
	Color       string
	Floor       string
	Sections    []entitymapper.Section
	Building    string
	PathNodes   []ImportPathNode
}
