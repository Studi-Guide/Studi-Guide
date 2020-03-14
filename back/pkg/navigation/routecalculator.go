package navigation

type RouteCalculator interface {
	GetRoute(start, end PathNode) (path []PathNode, distance int64, err error)
}