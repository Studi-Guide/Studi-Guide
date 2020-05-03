package services

import (
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
)

type LocationProvider interface {
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetLocation(name, building, campus string) (entitymapper.Location, error)
}