package osmnavigation

type OpenStreetMapNavigationProvider interface {
	GetRoute(start, end LatLngLiteral, locale string) ([]byte, error)
}
