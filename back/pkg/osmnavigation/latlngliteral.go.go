package osmnavigation

import "strconv"

type LatLngLiteral struct {
	Lat, Lng float64
}

func (l *LatLngLiteral) LatStr() string {
	return strconv.FormatFloat(l.Lat, 'f', -1, 64)
}

func (l *LatLngLiteral) LngStr() string {
	return strconv.FormatFloat(l.Lng, 'f', -1, 64)
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

	return LatLngLiteral{
		Lat: latF,
		Lng: lngF,
	}, nil
}