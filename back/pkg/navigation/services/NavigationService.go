package services

import (
	"studi-guide/pkg/building/db/entityservice"
	"studi-guide/pkg/navigation"
)

type NavigationService struct {
	routeCalc        navigation.RouteCalculator
	locationProvider LocationProvider
}

func NewNavigationService(routeCalculator navigation.RouteCalculator, locationProvider LocationProvider) (NavigationServiceProvider, error) {

	nodes, err := locationProvider.GetAllPathNodes()
	if err != nil {
		panic(err)
	}

	routeCalculator.Initialize(nodes)
	return &NavigationService{routeCalc: routeCalculator, locationProvider: locationProvider}, nil
}

func (n *NavigationService) CalculateFromString(startLocationName string, endLocationName string) (*navigation.NavigationRoute, error) {

	startLocation, err := n.locationProvider.GetLocation(startLocationName, "", "")
	if err != nil {
		return nil, err
	}

	endLocation, err := n.locationProvider.GetLocation(endLocationName, "", "")
	if err != nil {
		return nil, err
	}

	return n.Calculate(startLocation, endLocation)
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
