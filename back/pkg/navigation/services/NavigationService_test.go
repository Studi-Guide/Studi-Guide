package services

import (
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type NavigationMockService struct {
}

func NewNavigationMockService() NavigationServiceProvider {
	var rms NavigationMockService
	return &rms
}


func (n *NavigationMockService) CalculateFromString(startRoomName string, endRoomName string) (*[]navigation.Coordinate, error) {

	return nil, nil
}

func (n *NavigationMockService) Calculate(startRoom models.Room, endRoom models.Room) (*[]navigation.Coordinate, error) {

	return nil, nil
}

func (n *NavigationMockService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*[]navigation.Coordinate, error) {
	return nil, nil
}
