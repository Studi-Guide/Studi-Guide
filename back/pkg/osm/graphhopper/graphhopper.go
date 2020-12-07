package graphhopper

import (
	"log"
	"strings"
	"studi-guide/pkg/env"
	"studi-guide/pkg/osm"
	"studi-guide/pkg/utils"
)

var ghRootUrl = "https://graphhopper.com/api/1/"
var ghRouteUrl = "route"
var ghVehicle = "foot"

var xRateLimitCredits = "x-ratelimit-credits"
var xRateLimitLimit = "x-ratelimit-limit"
var xRateLimitRemaining = "x-ratelimit-remaining"
var xRateLimitReset = "x-ratelimit-reset"

type GraphHopper struct {
	apiKey     string
	bounds     osm.LatLngBounds
	httpClient utils.HttpClient
	logger     *log.Logger
}

func NewGraphHopper(env *env.Env, httpClient utils.HttpClient, logger *log.Logger) (osm.OpenStreetMapNavigationProvider, error) {

	southWest , _ := osm.NewLatLngLiteral(0, 0)
	northEast, _ := osm.NewLatLngLiteral(0, 0)
	boundLiteral, _ := osm.NewLatLngBounds(southWest, northEast)

	if len(env.OpenStreetMapBounds()) != 0 {

		bounds := strings.Split(env.OpenStreetMapBounds(), ";")
		a := strings.Split(bounds[0], ",")
		b := strings.Split(bounds[1], ",")

		southWest, err := osm.ParseLatLngLiteral(a[0], a[1])
		if err != nil {
			return nil, err
		}

		northEast, err := osm.ParseLatLngLiteral(b[0], b[1])
		if err != nil {
			return nil, err
		}

		boundLiteral, err = osm.NewLatLngBounds(southWest, northEast)
		if err != nil {
			return nil, err
		}
	}


	g := GraphHopper{
		apiKey:     env.GraphHopperApiKey(),
		bounds:     boundLiteral,
		httpClient: httpClient,
		logger:     logger,
	}

	return &g, nil
}
