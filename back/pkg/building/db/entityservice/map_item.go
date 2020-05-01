package entityservice

import (
	"studi-guide/pkg/navigation"
)

type MapItem struct {
	Doors     []Door
	Color     string
	Sections  []Section
	Campus    string
	Building  string
	PathNodes []*navigation.PathNode
	Floor     string
}
