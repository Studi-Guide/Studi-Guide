package services

import (
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
)

type LocationProvider interface {
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetLocation(name string) (entityservice.Location, error)
}