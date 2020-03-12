package services

import (
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
)

type NavigationService struct {
	routeCalc    navigation.RouteCalculator
	roomProvider models.RoomServiceProvider
}

func NewNavigationService(routeCalculator navigation.RouteCalculator, roomProvider models.RoomServiceProvider) (NavigationServiceProvider, error) {
	return &NavigationService{routeCalc: routeCalculator, roomProvider: roomProvider}, nil
}

func (n *NavigationService) CalculateFromString(startRoomName string, endRoomName string) (*[]navigation.PathNode, error) {

	startRoom, err := n.roomProvider.GetRoom(startRoomName)
	if err != nil {
		return nil, err
	}

	endRoom, err := n.roomProvider.GetRoom(endRoomName)
	if err != nil {
		return nil, err
	}

	return n.Calculate(startRoom, endRoom)
}

func (n *NavigationService) Calculate(startRoom models.Room, endRoom models.Room) (*[]navigation.PathNode, error) {
	nodes, err := n.routeCalc.GetRoute(startRoom.PathNode, endRoom.PathNode)
	return &nodes, err
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*[]navigation.PathNode, error) {

	//TODO implement
	return nil, nil
}
