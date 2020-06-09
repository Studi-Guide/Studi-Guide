package services

import (
	"studi-guide/pkg/navigation"
)

type NavigationServiceProvider interface {
	CalculateFromString(startName string, endName string) (*navigation.NavigationRoute, error)
	CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error)
}
