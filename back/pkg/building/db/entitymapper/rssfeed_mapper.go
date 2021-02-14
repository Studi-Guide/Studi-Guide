package entitymapper

import (
	"log"
	"studi-guide/pkg/building/db/ent"
	entfeed "studi-guide/pkg/building/db/ent/rssfeed"
)

//GetRssFeed returns a certain rss feed
func (r *EntityMapper) GetRssFeed(name string) (*ent.RssFeed, error) {
	b, err := r.client.RssFeed.Query().
		Where(entfeed.NameEqualFold(name)).
		First(r.context)

	if err != nil {
		return &ent.RssFeed{}, err
	}

	return b, nil
}

//AddRssFeed adds a rss feed into the db
func (r *EntityMapper) AddRssFeed(rssFeed ent.RssFeed) error {
	found, _ := r.client.RssFeed.Query().Where(entfeed.NameEqualFold(rssFeed.Name)).First(r.context)
	if found != nil {
		log.Printf("rssFeed %v already imported", rssFeed.Name)
		return nil
	}

	_, err := r.client.RssFeed.Create().
		SetName(rssFeed.Name).
		SetURL(rssFeed.URL).Save(r.context)

	if err != nil {
		log.Print("Error adding rssFeed:", rssFeed.Name, " Error:", err)
		return err
	}

	return nil
}
