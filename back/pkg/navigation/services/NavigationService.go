package services

import (
	"studi-guide/pkg/navigation"
)

type NavigationService struct {
	routeCalc        navigation.RouteCalculator
	pathNodeProvider PathNodeProvider
}

func NewNavigationService(routeCalculator navigation.RouteCalculator, pathNodeProvider PathNodeProvider) (NavigationServiceProvider, error) {

	nodes, err := pathNodeProvider.GetAllPathNodes()
	if err != nil {
		panic(err)
	}

	routeCalculator.Initialize(nodes)
	return &NavigationService{routeCalc: routeCalculator, pathNodeProvider: pathNodeProvider}, nil
}

func (n *NavigationService) CalculateFromString(startLocationName string, endLocationName string) (*navigation.NavigationRoute, error) {

	start, err := n.pathNodeProvider.GetPathNode(startLocationName)
	if err != nil {
		return nil, err
	}

	end, err := n.pathNodeProvider.GetPathNode(endLocationName)
	if err != nil {
		return nil, err
	}

	nodes, distance, err := n.routeCalc.GetRoute(start, end)
	return &navigation.NavigationRoute{Route: nodes, Distance: distance}, err
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error) {

	//TODO implement
	return nil, nil
}
