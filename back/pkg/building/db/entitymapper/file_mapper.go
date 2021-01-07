package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	entfile "studi-guide/pkg/building/db/ent/file"
	"studi-guide/pkg/file"
)

func (e * EntityMapper) mapFile(entFile *ent.File) file.File {
	return file.File{
		Name: entFile.Name,
		Path: e.env.AssetStorage() + entFile.Path,
	}
}

func (e* EntityMapper) fileMapper(files []file.File) ([]*ent.File, error) {
	var entFiles []*ent.File
	for _, i := range files {
		var f *ent.File
		var err error
		if q := e.client.File.Query().Where(entfile.PathEQ(i.Path)); q.ExistX(e.context) {
			f, err = q.First(e.context)
			if err != nil {
				return nil, err
			}
		} else {
			f, err = e.client.File.Create().
				SetName(i.Name).
				SetPath(i.Path).
				Save(e.context)
			if err != nil {
				return nil, err
			}
		}

		entFiles = append(entFiles, f)
	}
	return entFiles, nil
}