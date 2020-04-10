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

	nodes, err := roomProvider.GetAllPathNodes()
	if err != nil {
		panic(err)
	}

	routeCalculator.Initialize(nodes)
	return &NavigationService{routeCalc: routeCalculator, roomProvider: roomProvider}, nil
}

func (n *NavigationService) CalculateFromString(startRoomName string, endRoomName string) ([]navigation.PathNode, error) {

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

func (n *NavigationService) Calculate(startRoom models.Room, endRoom models.Room) ([]navigation.PathNode, error) {
	startNode := *startRoom.PathNodes[0]
	endNode := *endRoom.PathNodes[0]
	nodes, _, err := n.routeCalc.GetRoute(startNode, endNode)
	return nodes, err
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) ([]navigation.PathNode, error) {

	//TODO implement
	return nil, nil
}
