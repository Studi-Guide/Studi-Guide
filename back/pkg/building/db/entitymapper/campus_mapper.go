package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	entcampus "studi-guide/pkg/building/db/ent/campus"
)

func (r *EntityMapper) GetAllCampus() ([]*ent.Campus, error) {
	campus, err := r.client.Campus.Query().All(r.context)
	if err != nil {
		return nil, err
	}
	return campus, nil
}

func (r *EntityMapper) GetCampus(name string) (ent.Campus, error) {
	b, err := r.client.Campus.Query().
		Where(
			entcampus.Or(
				entcampus.NameEqualFold(name),
				entcampus.ShortNameEqualFold(name))).
		First(r.context)

	if err != nil {
		return ent.Campus{}, err
	}

	return *b, nil
}

func (r *EntityMapper) FilterCampus(name string) ([]*ent.Campus, error) {
	campus, err := r.client.Campus.Query().WithAddress().Where(
		entcampus.Or(
			entcampus.NameEqualFold(name),
			entcampus.ShortNameEqualFold(name))).All(r.context)

	if err != nil {
		return nil, err
	}

	return campus, nil
}
