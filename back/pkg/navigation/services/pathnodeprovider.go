package services

import (
	"studi-guide/pkg/navigation"
)

type PathNodeProvider interface {
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetRoutePoint(name string) (navigation.RoutePoint, error)

	GetPathNodeLocationData(node navigation.PathNode) (navigation.LocationData, error)
}
