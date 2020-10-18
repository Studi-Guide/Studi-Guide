package campus

import (
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

func TestNewRoomImporter(t *testing.T) {
	fileStringJson := "some_file.json"
	fileOther := "some_file.other"

	var dbService CampusProvider

	importer, _ := NewCampusImporter(fileStringJson, dbService)
	if _, ok := importer.(*CampusJsonImporter); !ok {
		t.Error("expected RoomJsonImporter; got: ", importer)
	}

	importer, err := NewCampusImporter(fileOther, dbService)
	if err == nil {
		t.Error("expected error; got: ", nil)
	}
	if importer != nil {
		t.Error("expected: ", nil, "; got:", importer)
	}

}

func TestRoomJsonImporter_RunImport(t *testing.T) {
	var dbService CampusProvider

	jsonImporter := CampusJsonImporter{dbService: dbService, file: "some_file"}
	err := jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	jsonImporter = CampusJsonImporter{dbService: dbService, file: "/random/file/which/does/not/exist.json"}
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

	jsonImporter = CampusJsonImporter{dbService: dbService, file: tmpFile}
	err = jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	os.Remove(tmpFile)
}

func TestRoomJsonImporter_ImportRealFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	campusprovider := NewMockCampusProvider(ctrl)
	campusprovider.EXPECT().AddCampus(gomock.Any()).Return(nil).MinTimes(2)
	jsonImporter := CampusJsonImporter{dbService: campusprovider, file: "../../../internal/campus.json"}
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
