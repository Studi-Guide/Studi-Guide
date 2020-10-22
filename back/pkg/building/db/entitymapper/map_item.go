package entitymapper

import (
	"studi-guide/pkg/navigation"
)

type MapItem struct {
	Doors     []Door
	Color     string
	Sections  []Section
	Building  string
	PathNodes []*navigation.PathNode
	Floor     string
}
