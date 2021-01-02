package location

import (
	"studi-guide/pkg/building/db/entitymapper"
)

type LocationProvider interface {
	GetAllLocations() ([]entitymapper.Location, error)
	GetLocation(name, building, campus string) (entitymapper.Location, error)
	FilterLocations(searchQuery, tag, floor, building, campus string) ([]entitymapper.Location, error)
	AddLocation(l entitymapper.Location) error
}
