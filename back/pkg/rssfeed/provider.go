package rssfeed

import "studi-guide/pkg/building/db/ent"

type Provider interface {
	GetRssFeed(name string) (*ent.RssFeed, error)
	AddRssFeed(rssFeed ent.RssFeed) error
}
