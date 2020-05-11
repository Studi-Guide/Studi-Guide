package info

import "studi-guide/pkg/building/db/entitymapper"

type BuildingProvider interface {
	GetAllBuildings() ([]entitymapper.Building, error)
	GetBuilding(name string) (entitymapper.Building, error)
	FilterBuildings(name string) ([]entitymapper.Building, error)
}
