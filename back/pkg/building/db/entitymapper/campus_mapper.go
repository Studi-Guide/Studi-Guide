package entitymapper

import (
	"log"
	"studi-guide/pkg/building/db/ent"
	entaddress "studi-guide/pkg/building/db/ent/address"
	entcampus "studi-guide/pkg/building/db/ent/campus"
)

func (r *EntityMapper) GetAllCampus() ([]*ent.Campus, error) {
	campus, err := r.client.Campus.Query().WithAddress().All(r.context)
	if err != nil {
		return nil, err
	}
	return campus, nil
}

func (r *EntityMapper) GetCampus(name string) (*ent.Campus, error) {
	b, err := r.client.Campus.Query().WithAddress().
		Where(
			entcampus.Or(
				entcampus.NameEqualFold(name),
				entcampus.ShortNameEqualFold(name))).
		First(r.context)

	if err != nil {
		return &ent.Campus{}, err
	}

	return b, nil
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
		log.Printf("campus %v already imported", campus.Name)
		return nil
	}

	campusCreate := r.client.Campus.
		Create().
		SetLongitude(campus.Longitude).
		SetLatitude(campus.Latitude).
		SetName(campus.Name).
		SetShortName(campus.ShortName)

	for _, address := range campus.Edges.Address {
		addressEntity, err := r.GetOrAddAddress(address)
		if err != nil {
			log.Print("Error adding address:", address, " Error:", err)
		} else {
			campusCreate.AddAddress(addressEntity)
		}
	}

	_, err := campusCreate.Save(r.context)
	if err != nil {
		log.Print("Error adding campus:", campus.Name, " Error:", err)
		return err
	}

	return nil
}

func (r *EntityMapper) GetOrAddAddress(address *ent.Address) (*ent.Address, error) {
	found, _ := r.client.Address.Query().Where(
		entaddress.And(
			entaddress.StreetEqualFold(address.Street),
			entaddress.PLZEQ(address.PLZ)),
		entaddress.NumberEqualFold(address.Street)).
		First(r.context)

	if found != nil {
		return found, nil
	}

	return r.client.Address.Create().
		SetPLZ(address.PLZ).
		SetNumber(address.Number).
		SetCountry(address.Country).
		SetCity(address.City).
		SetStreet(address.Street).Save(r.context)
}
