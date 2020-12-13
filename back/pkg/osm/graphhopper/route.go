package graphhopper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"studi-guide/pkg/osm"
	"studi-guide/pkg/osm/latlng"
)

func (g *GraphHopper) GetRoute(start, end latlng.LatLngLiteral, locale string) ([]osm.Route, error) {

	if len(g.apiKey) == 0 {
		return nil, errors.New("no api key was provided")
	}

	query := ghRootUrl + ghRouteUrl +
		"?point=" + start.LatStr() + "," + start.LngStr() + "&point=" + end.LatStr() + "," + end.LngStr() +
		"&vehicle=" + ghVehicle + "&key=" + g.apiKey + "&locale=" + locale + "&points_encoded=false"

	uri, _ := url.Parse(query)
	resp, err := g.httpClient.Do(&http.Request{
		Method: "GET",
		URL:    uri,
	})

	if err != nil {
		return nil, err
	}

	g.logRequestStats(resp.Header)

	if resp.StatusCode == http.StatusTooManyRequests {
		g.logger.Println("graphhopper rate limit reached")
		return nil, errors.New("currently not available")
	} else if resp.StatusCode != http.StatusOK {
		return nil, errors.New("get route error on endpoint")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var route GraphHopperRoute
	if err := json.Unmarshal(body, &route); err != nil {
		return nil, err
	}

	return route.ToOsmRoute(), nil
}

func (g *GraphHopper) logRequestStats(h http.Header) {
	g.logger.Println("GraphHopper Request Stats: ",
		xRateLimitCredits+":"+h.Get(xRateLimitCredits),
		xRateLimitLimit+":"+h.Get(xRateLimitLimit),
		xRateLimitRemaining+":"+h.Get(xRateLimitRemaining),
		xRateLimitReset+":"+h.Get(xRateLimitReset),
	)
}
