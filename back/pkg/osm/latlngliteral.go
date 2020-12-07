package osm

import (
	"errors"
	"strconv"
)

var LatLngLiteralRegex = "([0-9]+\\.?[0-9]+,[0-9]+\\.?[0-9]+)"

type LatLngLiteral struct {
	lat, lng float64
}

func NewLatLngLiteral(lat, lng float64) (LatLngLiteral, error) {
	if lat < -90.0 || lat > 90.0 || lng < -180.0 || lng > 180.0{
		return LatLngLiteral{}, errors.New("values do not match required bounds")
	}

	return LatLngLiteral{
		lat: lat,
		lng: lng,
	}, nil
}

func (l* LatLngLiteral) Lat() float64 {
	return l.lat
}

func (l* LatLngLiteral) Lng() float64 {
	return l.lng
}

func (l *LatLngLiteral) LatStr() string {
	return strconv.FormatFloat(l.lat, 'f', -1, 64)
}

func (l *LatLngLiteral) LngStr() string {
	return strconv.FormatFloat(l.lng, 'f', -1, 64)
}

func ParseLatLngLiteral(lat, lng string) (LatLngLiteral, error) {

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return LatLngLiteral{}, err
	}

	lngF, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return LatLngLiteral{}, err
	}

	return NewLatLngLiteral(latF, lngF)
}