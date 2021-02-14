package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	entfile "studi-guide/pkg/building/db/ent/file"
	"studi-guide/pkg/file"
)

func (r *EntityMapper) mapFile(entFile *ent.File) file.File {
	return file.File{
		Name: entFile.Name,
		Path: r.env.AssetStorage() + entFile.Path,
	}
}

func (r *EntityMapper) fileMapper(files []file.File) ([]*ent.File, error) {
	var entFiles []*ent.File
	for _, i := range files {
		var f *ent.File
		var err error
		if q := r.client.File.Query().Where(entfile.PathEQ(i.Path)); q.ExistX(r.context) {
			f, err = q.First(r.context)
			if err != nil {
				return nil, err
			}
		} else {
			f, err = r.client.File.Create().
				SetName(i.Name).
				SetPath(i.Path).
				Save(r.context)
			if err != nil {
				return nil, err
			}
		}

		entFiles = append(entFiles, f)
	}
	return entFiles, nil
}
