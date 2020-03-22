package services

import (
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type NavigationServiceProvider interface {
	// TODO replace coordinate array with a real route object
	CalculateFromString(startRoomName string, endRoomName string) ([]navigation.PathNode, error)

	Calculate(startRoom models.Room, endRoom models.Room) ([]navigation.PathNode, error)

	CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) ([]navigation.PathNode, error)
}
