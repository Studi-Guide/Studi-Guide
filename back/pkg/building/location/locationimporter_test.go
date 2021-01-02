package location

import (
"github.com/golang/mock/gomock"
"os"
"testing"
)

func TestNewLocationImporter(t *testing.T) {
	fileStringJson := "some_file.json"
	fileOther := "some_file.other"

	var dbService LocationProvider

	importer, _ := NewLocationImporter(fileStringJson, dbService)
	if _, ok := importer.(*JsonImporter); !ok {
		t.Error("expected JsonImporter; got: ", importer)
	}

	importer, err := NewLocationImporter(fileOther, dbService)
	if err == nil {
		t.Error("expected error; got: ", nil)
	}
	if importer != nil {
		t.Error("expected: ", nil, "; got:", importer)
	}

}

func TestLocationJsonImporter_RunImport(t *testing.T) {
	var dbService LocationProvider

	jsonImporter := JsonImporter{dbService: dbService, file: "some_file"}
	err := jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	jsonImporter = JsonImporter{dbService: dbService, file: "/random/file/which/does/not/exist.json"}
	err = jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	tmpFile := "/tmp/test.json"
	ensureDir("/tmp")
	_, err = os.Create(tmpFile)
	if err != nil {
		t.Error(err)
	}

	jsonImporter = JsonImporter{dbService: dbService, file: tmpFile}
	err = jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	os.Remove(tmpFile)
}

func TestLocationImporter_ImportRealFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	campusprovider := NewMockLocationProvider(ctrl)
	campusprovider.EXPECT().AddLocation(gomock.Any()).Return(nil).MinTimes(8)
	jsonImporter := JsonImporter{dbService: campusprovider, file: "../../../internal/location.json"}
	err := jsonImporter.RunImport()
	if err != nil {
		t.Error("expected error; got: ", err)
	}
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

