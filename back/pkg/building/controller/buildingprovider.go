package controller

import "studi-guide/pkg/building/model"

type BuildingProvider interface {
	GetAllBuildings() ([]model.Building, error)
	GetBuilding(name string) (model.Building, error)
	FilterBuildings(name string) ([]model.Building, error)
}
