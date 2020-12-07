package osm

import "errors"

type LatLngBounds struct {
	southWest, northEast LatLngLiteral
}

func (l *LatLngBounds) SouthWest() LatLngLiteral {
	return l.southWest
}

func (l *LatLngBounds) NorthEast() LatLngLiteral {
	return l.northEast
}

func NewLatLngBounds(southWest, northEast LatLngLiteral) (LatLngBounds, error) {

	if southWest.lat > northEast.lat || southWest.lng > northEast.lng {
		return LatLngBounds{}, errors.New("south west and north east coordinates do no specify bounds correctly")
	}

	return LatLngBounds{
		southWest: southWest,
		northEast: northEast,
	}, nil
}

func (b LatLngBounds) IncludeLiteral(l LatLngLiteral) bool {
	return b.southWest.lat <= l.lat && b.southWest.lng <= l.lng &&
		b.northEast.lat >= l.lat && b.northEast.lng >= l.lng
}
