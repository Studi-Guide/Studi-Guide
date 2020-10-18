package entitymapper

import (
	"github.com/prometheus/common/log"
	"studi-guide/pkg/building/db/ent"
	entcampus "studi-guide/pkg/building/db/ent/campus"
)

func (r *EntityMapper) GetAllCampus() ([]*ent.Campus, error) {
	campus, err := r.client.Campus.Query().WithAddress().All(r.context)
	if err != nil {
		return nil, err
	}
	return campus, nil
}

func (r *EntityMapper) GetCampus(name string) (ent.Campus, error) {
	b, err := r.client.Campus.Query().WithAddress().
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

func (r *EntityMapper) AddCampus(campus ent.Campus) error {
	found, _ := r.client.Campus.Query().Where(entcampus.NameEqualFold(campus.Name)).First(r.context)
	if found != nil {
		log.Info("campus ", campus.Name, " already imported")
		return nil
	}

	_, err := r.client.Campus.
		Create().
		SetLongitude(campus.Longitude).
		SetLatitude(campus.Latitude).
		SetName(campus.Name).
		SetShortName(campus.ShortName).
		AddAddress(campus.Edges.Address...).
		Save(r.context)

	if err != nil {
		log.Fatal("Error adding campus:", campus.Name, " Error:", err)
		return err
	}

	return nil
}
