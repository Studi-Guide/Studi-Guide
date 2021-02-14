package osm

import "studi-guide/pkg/osm/latlng"

//OpenStreetMapNavigationProvider runs the route calculator via open street map
type OpenStreetMapNavigationProvider interface {
	GetRoute(start, end latlng.LatLngLiteral, locale string) ([]Route, error)
}
