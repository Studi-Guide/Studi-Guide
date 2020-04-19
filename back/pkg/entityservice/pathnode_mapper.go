package entityservice

import (
	"errors"
	"log"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/pathnode"
	"studi-guide/pkg/navigation"
)

func (r *EntityService) pathNodeArrayMapper(pathNodePtr []*ent.PathNode, availableNodes []*navigation.PathNode) []*navigation.PathNode {
	var pathNodes []*navigation.PathNode
	for _, node := range pathNodePtr {
		pathNodes = append(pathNodes, r.pathNodeMapper(node, availableNodes, false))
	}
	return pathNodes
}

func (r *EntityService) pathNodeMapper(entPathNode *ent.PathNode, availableNodes []*navigation.PathNode, resolveConnectedNodes bool) *navigation.PathNode {

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

func (r *EntityService) mapPathNodeArray(pathNodePtr []*navigation.PathNode) ([]*ent.PathNode, error) {

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

func (r *EntityService) mapPathNode(p *navigation.PathNode) (*ent.PathNode, error) {

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

