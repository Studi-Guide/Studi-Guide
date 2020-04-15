package location

import (
	"studi-guide/pkg/entityservice"
)

type LocationProvider interface {
	GetAllLocations() ([]entityservice.Location, error)
	GetLocation(name string) (entityservice.Location, error)
	FilterLocations(name, tag, floor, building, campus string) ([]entityservice.Location, error)
}
