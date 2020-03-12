package navigation

type RouteCalculator interface {
	GetRoute(start, end PathNode) ([]PathNode, error)
}