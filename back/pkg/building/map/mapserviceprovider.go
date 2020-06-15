package maps

import "studi-guide/pkg/building/db/entitymapper"

type MapServiceProvider interface {
	GetAllMapItems() ([]entitymapper.MapItem, error)
	GetMapItemByPathNodeID(pathNodeID int) (entitymapper.MapItem, error)
	FilterMapItems(floor, building, campus string) ([]entitymapper.MapItem, error)
}
