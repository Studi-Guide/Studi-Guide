package navigation

type RouteCalculator interface {
	Initialize(pathNodes []PathNode)
	GetRoute(start, end PathNode) (path []PathNode, distance int64, err error)
}
