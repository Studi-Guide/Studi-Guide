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

func (n *NavigationService) CalculateFromString(startLocationName string, endLocationName string) (*navigation.NavigationRoute, error) {
	startLocation, err := n.locationProvider.GetLocation(startLocationName)
	if err != nil {
		return nil, err
	}

	endRoom, err := n.roomProvider.GetRoom(endRoomName)
	if err != nil {
		return nil, err
	}

	return n.Calculate(startRoom, endRoom)
}

func (n *NavigationService) Calculate(startLocation entityservice.Location, endLocation entityservice.Location) (*navigation.NavigationRoute, error) {
	startNode := startLocation.PathNode
	endNode := endLocation.PathNode
	nodes, distance, err := n.routeCalc.GetRoute(startNode, endNode)
	return &navigation.NavigationRoute{Route: nodes, Distance: distance}, err
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error) {

	//TODO implement
	return nil, nil
}
