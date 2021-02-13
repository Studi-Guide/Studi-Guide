package campus

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"studi-guide/pkg/building/db/ent"
)

type CampusImporter interface {
	RunImport() error
}

type CampusJsonImporter struct {
	dbService CampusProvider
	file      string
}

func NewCampusImporter(file string, dbService CampusProvider) (CampusImporter, error) {
	var i CampusImporter = nil
	ext := filepath.Ext(file)
	if ext == ".json" {
		i = &CampusJsonImporter{dbService: dbService, file: file}
	} else {
		return nil, errors.New("Unknown extension")
	}

	return i, nil
}

func (r *CampusJsonImporter) RunImport() error {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var items []ent.Campus
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		return err
	}

	for _, campus := range items {
		err = r.dbService.AddCampus(campus)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}
