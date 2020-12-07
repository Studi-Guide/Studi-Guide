package osm

import "studi-guide/pkg/osm/latlng"

type OpenStreetMapNavigationProvider interface {
	GetRoute(start, end latlng.LatLngLiteral, locale string) ([]byte, error)
}
