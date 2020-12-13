package latlng

import "errors"

type LatLngBounds struct {
	SouthWest, NorthEast LatLngLiteral
}

func NewLatLngBounds(southWest, northEast LatLngLiteral) (LatLngBounds, error) {

	if southWest.Lat > northEast.Lat || southWest.Lng > northEast.Lng {
		return LatLngBounds{}, errors.New("south west and north east coordinates do no specify bounds correctly")
	}

	return LatLngBounds{
		SouthWest: southWest,
		NorthEast: northEast,
	}, nil
}

func (b LatLngBounds) IncludeLiteral(l LatLngLiteral) bool {
	return b.SouthWest.Lat <= l.Lat && b.SouthWest.Lng <= l.Lng &&
		b.NorthEast.Lat >= l.Lat && b.NorthEast.Lng >= l.Lng
}
