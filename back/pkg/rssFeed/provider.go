package rssFeed

import "studi-guide/pkg/building/db/ent"

type Provider interface {
	GetRssFeed(name string) (*ent.RssFeed, error)
	AddRssFeed(campus ent.RssFeed) error
}
