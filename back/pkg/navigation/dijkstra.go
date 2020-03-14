package navigation

import (
	"github.com/RyanCarrier/dijkstra"
	"log"
	"time"
)

type DijkstraNavigation struct {
	graph *dijkstra.Graph
	pathNodes map[int]PathNode
}

func NewDijkstraNavigation(pathNodes []PathNode) (*DijkstraNavigation) {

	d := DijkstraNavigation{pathNodes: make(map[int]PathNode), graph: dijkstra.NewGraph()}

	// Add vertices and fill map
	for _, node := range pathNodes {
		d.pathNodes[node.Id] = node
		d.graph.AddVertex(node.Id)
	}

	// Add arcs
	for _, source := range pathNodes {
		for _, dest := range source.ConnectedNodes {
			d.graph.AddArc(source.Id, dest.Id, int64(source.Coordinate.DistanceTo(dest.Coordinate)))
		}
	}

	return &d
}

func (d *DijkstraNavigation) GetRoute(start, end PathNode) (path []PathNode, distance int64, err error) {

	startTime := time.Now()
	defer log.Println("Route calculation from", start, "to", end, "took", time.Since(startTime))

	if start.Coordinate.Equals(end.Coordinate) {
		return nil, 0, nil
	}

	bestPath, err := d.graph.Shortest(start.Id, end.Id)

	if err != nil {
		return nil, 0, err
	}

	pathNodes := make([]PathNode, len(bestPath.Path))
	for i, id := range bestPath.Path {
		pathNodes[i] = d.pathNodes[id]
	}

	return pathNodes, bestPath.Distance, nil
}