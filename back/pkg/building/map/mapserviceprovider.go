package maps

import "studi-guide/pkg/building/db/entitymapper"

type MapServiceProvider interface {
	GetAllMapItems() ([]entitymapper.MapItem, error)
	FilterMapItems(floor, building, campus string) ([]entitymapper.MapItem, error)
}
