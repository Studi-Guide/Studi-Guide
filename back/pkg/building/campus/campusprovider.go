package campus

import "studi-guide/pkg/building/db/ent"

type CampusProvider interface {
	GetAllCampus() ([]*ent.Campus, error)
	GetCampus(name string) (*ent.Campus, error)
	FilterCampus(name string) ([]*ent.Campus, error)
	AddCampus(campus ent.Campus) error
}
