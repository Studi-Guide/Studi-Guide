package info

import "studi-guide/pkg/building/db/entityservice"

type BuildingProvider interface {
	GetAllBuildings() ([]entityservice.Building, error)
	GetBuilding(name string) (entityservice.Building, error)
	FilterBuildings(name string) ([]entityservice.Building, error)
}
