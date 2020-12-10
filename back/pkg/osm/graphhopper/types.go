package graphhopper

import (
	"fmt"
	"studi-guide/pkg/osm"
	"studi-guide/pkg/osm/latlng"
)

type GraphHopperRouteInstruction struct {
	Distance   float64   `json:"distance"`
	Heading    float64   `json:"heading"`
	Interval   []float64 `json:"interval"`
	StreetName string    `json:"street_name"`
	Text       string    `json:"text"`
	Time       float64   `json:"time"`
}

type GraphHopperRoutePoints struct {
	Coordinates []interface{} `json:"coordinates"`
	Type        string        `json:"type"`
}

type GraphHopperPath struct {
	Distance     float64                       `json:"distance"`
	Time         float64                       `json:"time"`
	Points       GraphHopperRoutePoints         `json:"points"`
	Instructions []GraphHopperRouteInstruction `json:"instructions"`
}

type GraphHopperRoute struct {
	Info  interface{}       `json:"info"`
	Paths []GraphHopperPath `json:"paths"`
}

func (r *GraphHopperRoute) ToOsmRoute() []osm.Route {
	var routes []osm.Route

	for _, path := range r.Paths {
		r := osm.Route{
			Distance:     path.Distance,
			Time:         path.Time,
			Points:       osm.RoutePoints{
				Coordinates: []latlng.LatLngLiteral{},
				Type:        path.Points.Type,
			},
			Instructions: nil,
		}

		for _, coordinate := range path.Points.Coordinates {
			//var points []float64
			//for _, c := range coordinate.([]interface{})[0].(float64) {
			//
			//}
			
			r.Points.Coordinates = append(r.Points.Coordinates, latlng.LatLngLiteral{
				Lat: coordinate.([]interface{})[1].(float64),
				Lng: coordinate.([]interface{})[0].(float64),
			})
			fmt.Println(coordinate)
		}

		for _, instruction := range path.Instructions {
			r.Instructions = append(r.Instructions, osm.RouteInstruction{
				Distance:   instruction.Distance,
				Heading:    instruction.Heading,
				Interval:   instruction.Interval,
				StreetName: instruction.StreetName,
				Text:       instruction.Text,
				Time:       instruction.Time,
			})
		}

		routes = append(routes, r)
	}

	return routes
}