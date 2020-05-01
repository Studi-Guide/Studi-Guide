package maps

import "studi-guide/pkg/building/db/entityservice"

type MapServiceProvider interface {
	GetAllMapItems() ([]entityservice.MapItem, error)
	FilterMapItems(floor, building, campus string) ([]entityservice.MapItem, error)
}
