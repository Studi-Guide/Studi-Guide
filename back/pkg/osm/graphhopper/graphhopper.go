package graphhopper

import (
	"log"
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
	httpClient utils.HttpClient
	logger     *log.Logger
}

func NewGraphHopper(env *env.Env, httpClient utils.HttpClient, logger *log.Logger) (osm.OpenStreetMapNavigationProvider, error) {

	g := GraphHopper{
		apiKey:     env.GraphHopperApiKey(),
		httpClient: httpClient,
		logger:     logger,
	}

	return &g, nil
}
