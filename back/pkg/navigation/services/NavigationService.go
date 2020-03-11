package services

import (
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type NavigationService struct {
}

func NewNavigationService() (NavigationServiceProvider, error) {

	return &NavigationService{}, nil
}

func (n *NavigationService) CalculateFromString(startRoomName string, endRoomName string) (*[]navigation.Coordinate, error) {

	return nil, nil
}

func (n *NavigationService) Calculate(startRoom models.Room, endRoom models.Room) (*[]navigation.Coordinate, error) {

	return nil, nil
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*[]navigation.Coordinate, error) {
	return nil, nil
}
