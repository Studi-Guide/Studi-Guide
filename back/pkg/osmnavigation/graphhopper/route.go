package graphhopper

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"studi-guide/pkg/osmnavigation"
)

func (g *GraphHopper) GetRoute(start, end osmnavigation.LatLngLiteral, locale string) (error, []byte) {

	if len(g.apiKey) == 0 {
		return errors.New("no api key was provided"), nil
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
		return err, nil
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("get route error on endpoint"), nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	return nil, body
}
