package services

import (
	"studi-guide/pkg/building/db/entityservice"
	"studi-guide/pkg/navigation"
)

type LocationProvider interface {
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetLocation(name, building, campus string) (entityservice.Location, error)
}