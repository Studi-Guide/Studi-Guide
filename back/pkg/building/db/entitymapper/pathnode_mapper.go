package entitymapper

import (
	"errors"
	"log"
	"strings"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/pathnode"
	"studi-guide/pkg/navigation"
)

func (r *EntityMapper) pathNodeArrayMapper(pathNodePtr []*ent.PathNode, availableNodes []*navigation.PathNode) []*navigation.PathNode {
	var pathNodes []*navigation.PathNode
	for _, node := range pathNodePtr {
		pathNodes = append(pathNodes, r.pathNodeMapper(node, availableNodes, false))
	}
	return pathNodes
}

func (r *EntityMapper) pathNodeMapper(entPathNode *ent.PathNode, availableNodes []*navigation.PathNode, resolveConnectedNodes bool) *navigation.PathNode {

	entPathNode, err := r.client.PathNode.Query().Where(pathnode.IDEQ(entPathNode.ID)).WithLinkedTo().First(r.context)
	if err != nil {
		return &navigation.PathNode{}
	}

	for _, node := range availableNodes {
		if node.Id == entPathNode.ID {
			return node
		}
	}

	p := navigation.PathNode{
		Id:             entPathNode.ID,
		Coordinate:     navigation.Coordinate{X: entPathNode.XCoordinate, Y: entPathNode.YCoordinate, Z: entPathNode.ZCoordinate},
		Group:          nil,
		ConnectedNodes: nil,
	}

	availableNodes = append(availableNodes, &p)
	if resolveConnectedNodes {
		p.ConnectedNodes = r.pathNodeArrayMapper(entPathNode.Edges.LinkedTo, availableNodes)
	}

	return &p
}

func (r *EntityMapper) mapPathNodeArray(pathNodePtr []*navigation.PathNode) ([]*ent.PathNode, error) {

	var entPathNodes []*ent.PathNode

	var errorStrs []string

	for _, ptr := range pathNodePtr {
		p, err := r.mapPathNode(ptr)
		if err != nil {
			errorStrs = append(errorStrs, err.Error())
			continue
		}
		entPathNodes = append(entPathNodes, p)
	}

	if len(errorStrs) != 0 {
		return entPathNodes, errors.New(strings.Join(errorStrs, ","))
	}

	return entPathNodes, nil
}

func (r *EntityMapper) mapPathNode(p *navigation.PathNode) (*ent.PathNode, error) {

	if p.Id != 0 {
		node, err := r.client.PathNode.Get(r.context, p.Id)
		if node != nil {
			return node, nil
		}

		switch t := err.(type) {
		default:
			log.Fatal(t)
			return nil, err
		case *ent.NotFoundError:
			// do nothing
		}
	}

	log.Println("add path node:", p)
	return r.client.PathNode.Create().
		SetID(p.Id).
		SetXCoordinate(p.Coordinate.X).
		SetYCoordinate(p.Coordinate.Y).
		SetZCoordinate(p.Coordinate.Z).
		Save(r.context)
}

func (r *EntityMapper) GetAllPathNodes() ([]navigation.PathNode, error) {
	nodesPrt, err := r.client.PathNode.Query().WithLinkedFrom().WithLinkedTo().All(r.context)
	if err != nil {
		return nil, err
	}

	var nodes []navigation.PathNode
	var nodesCache []*navigation.PathNode
	for _, nodePtr := range nodesPrt {

		node := *r.pathNodeMapper(nodePtr, nodesCache, true)
		nodes = append(nodes, node)
		nodesCache = append(nodesCache, &node)
	}

	return nodes, nil
}

func (r *EntityMapper) GetPathNodeLocationData(node navigation.PathNode) (navigation.LocationData, error) {
	entityNode, err := r.client.PathNode.Query().WithLocation(func(locationQuery *ent.LocationQuery) {
		locationQuery.WithBuilding()
	}).WithMapitem(func(query *ent.MapItemQuery) {
		query.WithBuilding()
	}).Where(pathnode.IDEQ(node.Id)).First(r.context)

	if err != nil {
		return navigation.LocationData{}, err
	}

	// look for linked location
	location := entityNode.Edges.Location
	if location != nil && location.Edges.Building != nil {
		return navigation.LocationData{Building: location.Edges.Building.Name, Floor: location.Floor}, nil
	}

	// look for linked mapitem
	mapitem := entityNode.Edges.Mapitem
	if mapitem != nil && mapitem.Edges.Building != nil {
		return navigation.LocationData{Building: mapitem.Edges.Building.Name, Floor: mapitem.Floor}, nil
	}

	return navigation.LocationData{}, errors.New("no linked data found")
}

func (r *EntityMapper) linkPathNode(pathNode *navigation.PathNode) error {

	var connectedIDs []int

	//Get database IDs
	for _, connectedNode := range pathNode.ConnectedNodes {

		entityConnectedNode, err := r.client.PathNode.Get(r.context, connectedNode.Id)
		if err != nil {
			return err
		}

		connectedIDs = append(connectedIDs, entityConnectedNode.ID)
	}

	entityNode, err := r.client.PathNode.Get(r.context, pathNode.Id)
	if err != nil {
		return err
	}

	update := entityNode.Update()
	update.AddLinkedToIDs(connectedIDs...)
	entityNode, err = update.Save(r.context)
	return err
}
