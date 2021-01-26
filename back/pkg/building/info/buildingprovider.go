package info

import (
	"studi-guide/pkg/building/db/ent"
)

type BuildingProvider interface {
	GetAllBuildings() ([]*ent.Building, error)
	GetBuilding(name string) (*ent.Building, error)
	FilterBuildings(name string) ([]*ent.Building, error)
}
