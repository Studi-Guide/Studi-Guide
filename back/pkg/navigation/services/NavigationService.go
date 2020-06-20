package services

import (
	"fmt"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/utils"
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

	start, err := n.pathNodeProvider.GetRoutePoint(startLocationName)
	if err != nil {
		return nil, &utils.QueryError{
			Query: fmt.Sprintf("Startlocation %v: not found!", startLocationName),
			Err:   err,
		}
	}

	end, err := n.pathNodeProvider.GetRoutePoint(endLocationName)
	if err != nil {
		return nil, &utils.QueryError{
			Query: fmt.Sprintf("Endlocation %v: not found!", endLocationName),
			Err:   err,
		}
	}

	nodes, distance, err := n.routeCalc.GetRoute(start.Node, end.Node)
	route := GenerateNavigationRoute(nodes, distance, n.pathNodeProvider)
	route.End = end
	route.Start = start
	return &route, err
}

func (n *NavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error) {

	//TODO implement
	return nil, nil
}

func GenerateNavigationRoute(nodes []navigation.PathNode, distance int64, provider PathNodeProvider) navigation.NavigationRoute {

	// Create array with at least one element
	var routeSections = []navigation.RouteSection{{}}
	var routeSectionCnt = 0
	for idx, node := range nodes {

		var routeSection = &routeSections[routeSectionCnt]

		// Try to get the linked building and floor
		locationData, error := provider.GetPathNodeLocationData(node)
		// go here for initialization
		if idx == 0 {
			routeSection.Floor = locationData.Floor
			routeSection.Building = locationData.Building
			routeSection.Route = append(routeSection.Route, node)
		} else {

			// go here when node fits the route section or no information is found
			if (routeSection.Building == locationData.Building && routeSection.Floor == locationData.Floor) || error != nil {
				routeSection.Route = append(routeSection.Route, node)

				// add distance to last coordinate
				routeSection.Distance += int64(node.Coordinate.DistanceTo(routeSection.Route[len(routeSection.Route)-2].Coordinate))
			} else {
				// create a new route section and add it
				var newRouteSection = navigation.RouteSection{
					Route:       []navigation.PathNode{node},
					Description: "",
					Distance:    0,
					Building:    locationData.Building,
					Floor:       locationData.Floor,
				}

				routeSections = append(routeSections, newRouteSection)
				routeSectionCnt++
			}
		}
	}

	return navigation.NavigationRoute{
		RouteSections: routeSections,
		Distance:      distance,
	}
}
