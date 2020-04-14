package maps

import "studi-guide/pkg/entityservice"

type MapServiceProvider interface {
	GetAllMapItems() ([]entityservice.MapItem, error)
	FilterMapItems(floor, building, campus string) ([]entityservice.MapItem, error)
}
