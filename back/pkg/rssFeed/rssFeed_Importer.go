package rssFeed

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"studi-guide/pkg/building/db/ent"
)

type Importer interface {
	RunImport() error
}

type JsonImporter struct {
	dbService Provider
	file      string
}

func NewRssFeedImporter(file string, dbService Provider) (Importer, error) {
	var i Importer = nil
	ext := filepath.Ext(file)
	if ext == ".json" {
		i = &JsonImporter{dbService: dbService, file: file}
	} else {
		return nil, errors.New("unknown extension")
	}

	return i, nil
}

func (r *JsonImporter) RunImport() error {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var items []ent.RssFeed
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		return err
	}

	for _, feed := range items {
		err = r.dbService.AddRssFeed(feed)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}
