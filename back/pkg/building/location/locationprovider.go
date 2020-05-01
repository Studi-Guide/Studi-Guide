package location

import (
	"studi-guide/pkg/building/db/entityservice"
)

type LocationProvider interface {
	GetAllLocations() ([]entityservice.Location, error)
	GetLocation(name, building, campus string) (entityservice.Location, error)
	FilterLocations(name, tag, floor, building, campus string) ([]entityservice.Location, error)
}
