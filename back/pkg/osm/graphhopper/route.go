package graphhopper

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"studi-guide/pkg/osm"
)

func (g *GraphHopper) GetRoute(start, end osm.LatLngLiteral, locale string) ([]byte, error) {

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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("get route error on endpoint")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
