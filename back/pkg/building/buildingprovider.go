package building

import "studi-guide/pkg/entityservice"

type BuildingProvider interface {
	GetAllBuildings() ([]entityservice.Building, error)
	GetBuilding(name string) (entityservice.Building, error)
	FilterBuildings(name string) ([]entityservice.Building, error)
}
