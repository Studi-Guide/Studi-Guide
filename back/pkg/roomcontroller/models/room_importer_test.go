package models

import (
	"os"
	"testing"

)

func TestNewRoomImporter(t *testing.T) {
	fileStringXml := "some_file.xml"
	fileStringJson := "some_file.json"
	fileOther := "some_file.other"

	var dbService RoomServiceProvider

	importer, _ := NewRoomImporter(fileStringXml, dbService)
	if _, ok := importer.(*RoomXmlImporter); !ok {
		t.Error("expected RoomXmlImporter; got: ", importer)
	}

	importer, _ = NewRoomImporter(fileStringJson, dbService)
	if _, ok := importer.(*RoomJsonImporter); !ok {
		t.Error("expected RoomJsonImporter; got: ", importer)
	}

	importer, err := NewRoomImporter(fileOther, dbService)
	if err == nil {
		t.Error("expected error; got: ", nil)
	}
	if importer != nil {
		t.Error("expected: ", nil, "; got:", importer)
	}

}

func TestRoomJsonImporter_RunImport(t *testing.T) {
	var dbService RoomServiceProvider

	jsonImporter := RoomJsonImporter{dbService:dbService, file: "some_file"}
	err := jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	jsonImporter = RoomJsonImporter{dbService:dbService, file: "/random/file/which/does/not/exist.json"}
	err = jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}


	tmpFile := "/tmp/test.json"
	_, err = os.Create(tmpFile)
	if err != nil {
		t.Error(err)
	}

	jsonImporter = RoomJsonImporter{dbService:dbService, file:tmpFile}
	err = jsonImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	os.Remove(tmpFile)

}

func TestRoomXmlImporter_RunImport(t *testing.T) {

	var dbService RoomServiceProvider


	xmlImporter := RoomXmlImporter{dbService:dbService, file: "some_file"}
	err := xmlImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	xmlImporter = RoomXmlImporter{dbService:dbService, file: "/random/file/which/does/not/exist.xml"}
	err = xmlImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	tmpFile := "/tmp/test.xml"
	_, err = os.Create(tmpFile)
	if err != nil {
		t.Error(err)
	}

	xmlImporter = RoomXmlImporter{dbService:dbService, file:tmpFile}
	err = xmlImporter.RunImport()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	os.Remove(tmpFile)

}