package services

import (
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
)

type NavigationServiceProvider interface {
	// TODO replace coordinate array with a real route object
	CalculateFromString(startRoomName string, endRoomName string) (*navigation.NavigationRoute, error)

	Calculate(startRoom entityservice.Location, endRoom entityservice.Location) (*navigation.NavigationRoute, error)

	CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error)
}
