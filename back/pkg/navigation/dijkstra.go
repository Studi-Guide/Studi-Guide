package navigation

import "github.com/RyanCarrier/dijkstra"

type DijkstraNavigation struct {
	graph dijkstra.Graph
	pathNodes []PathNode
}

func NewDijkstraNavigation(pathNodes []PathNode) (*DijkstraNavigation) {
	return &DijkstraNavigation{pathNodes:pathNodes}
}

func (d *DijkstraNavigation) GetRoute(start, end PathNode) ([]PathNode) {
	return []PathNode{}
}