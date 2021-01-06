package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/file"
)

func (e * EntityMapper) mapFile(entFile *ent.File) file.File {
	return file.File{
		Name: entFile.Name,
		Path: e.env.AssetStorage() + entFile.Path,
	}
}
