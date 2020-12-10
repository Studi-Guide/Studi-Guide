package osm

import "studi-guide/pkg/osm/latlng"

type RouteInstruction struct {
	Distance   float64
	Heading    float64
	Interval   []float64
	StreetName string
	Text       string
	Time       float64
}

type RoutePoints struct {
	Coordinates []latlng.LatLngLiteral
	Type        string
}

type Route struct {
	Distance     float64
	Time         float64
	Points       RoutePoints
	Instructions []RouteInstruction
}
