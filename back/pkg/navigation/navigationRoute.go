package navigation

type NavigationRoute struct {
	Start         RoutePoint
	End           RoutePoint
	RouteSections []RouteSection
	Distance      int64
}
