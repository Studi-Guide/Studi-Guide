package controllers

import (
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/navigation/services"
	"studi-guide/pkg/roomcontroller/models"
)

type NavigationMockService struct {
}

func NewNavigationMockService() services.NavigationServiceProvider {
	var rms NavigationMockService
	return &rms
}


func (n *NavigationMockService) CalculateFromString(startRoomName string, endRoomName string) (*[]navigation.PathNode, error) {

	return nil, nil
}

func (n *NavigationMockService) Calculate(startRoom models.Room, endRoom models.Room) (*[]navigation.PathNode, error) {

	return nil, nil
}

func (n *NavigationMockService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*[]navigation.PathNode, error) {
	return nil, nil
}
