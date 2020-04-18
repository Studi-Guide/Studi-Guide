package services

import (
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
)

type NavigationServiceProvider interface {
	// TODO replace coordinate array with a real route object
	CalculateFromString(startRoomName string, endRoomName string) ([]navigation.PathNode, error)

	Calculate(startRoom entityservice.Location, endRoom entityservice.Location) ([]navigation.PathNode, error)

	CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) ([]navigation.PathNode, error)
}
