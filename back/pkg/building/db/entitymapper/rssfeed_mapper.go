package entitymapper

import (
	"log"
	"studi-guide/pkg/building/db/ent"
	entfeed "studi-guide/pkg/building/db/ent/rssfeed"
)

func (r *EntityMapper) GetRssFeed(name string) (*ent.RssFeed, error) {
	b, err := r.client.RssFeed.Query().
		Where(entfeed.NameEqualFold(name)).
		First(r.context)

	if err != nil {
		return &ent.RssFeed{}, err
	}

	return b, nil
}

func (r *EntityMapper) AddRssFeed(rssfeed ent.RssFeed) error {
	found, _ := r.client.RssFeed.Query().Where(entfeed.NameEqualFold(rssfeed.Name)).First(r.context)
	if found != nil {
		log.Printf("rssfeed %v already imported", rssfeed.Name)
		return nil
	}

	_, err := r.client.RssFeed.Create().
		SetName(rssfeed.Name).
		SetURL(rssfeed.URL).Save(r.context)

	if err != nil {
		log.Print("Error adding rssfeed:", rssfeed.Name, " Error:", err)
		return err
	}

	return nil
}
