package graphhopper

import (
	"studi-guide/pkg/env"
	"studi-guide/pkg/osm"
	"studi-guide/pkg/utils"
)

var ghRootUrl = "https://graphhopper.com/api/1/"
var ghRouteUrl = "route"
var ghVehicle = "foot"

type GraphHopper struct {
	apiKey     string
	httpClient utils.HttpClient
}

func NewGraphHopper(env *env.Env, httpClient utils.HttpClient) (osm.OpenStreetMapNavigationProvider, error) {
	g := GraphHopper{
		apiKey:     env.GraphHopperApiKey(),
		httpClient: httpClient,
	}

	return &g, nil
}
