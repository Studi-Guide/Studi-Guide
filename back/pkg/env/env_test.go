package env

import (
	"os"
	"testing"
)

func Test_NewEnv_1(t *testing.T) {

	dBDriverNameVar := "SouthWest"
	dbDataSourceVar := "b"
	frontendPathVar := "3"
	graphHopperApiKeyVar := "D"
	developVar := "-"

	_ = os.Setenv(dBDriverNameKey, dBDriverNameVar)
	_ = os.Setenv(dbDataSourceKey, dbDataSourceVar)
	_ = os.Setenv(frontendPath, frontendPathVar)
	_ = os.Setenv(graphHopperApiKey, graphHopperApiKeyVar)
	_ = os.Setenv(develop, developVar)

	env := NewEnv()

	if env.DbDriverName() != dBDriverNameVar {
		t.Error("env.dbDriverName", env.dbDriverName)
	}

	if env.DbDataSource() != dbDataSourceVar {
		t.Error("env.dbDataSource", env.dbDataSource)
	}

	if env.FrontendPath() != frontendPathVar {
		t.Error("env.frontendPath", env.frontendPath)
	}

	if env.GraphHopperApiKey() != graphHopperApiKeyVar {
		t.Error("env.graphHopperApiKey", env.graphHopperApiKey)
	}

	if env.Develop() != false {
		t.Error("env.develop", env.develop)
	}

}

func Test_NewEnv_2(t *testing.T) {

	graphHopperApiKeyVar := ""
	developVar := "TRUE"

	// reset env variables
	_ = os.Setenv(dBDriverNameKey, "")
	_ = os.Setenv(dbDataSourceKey, "")
	_ = os.Setenv(frontendPath, "")
	_ = os.Setenv(graphHopperApiKey, graphHopperApiKeyVar)
	_ = os.Setenv(develop, developVar)

	env := NewEnv()

	if env.DbDriverName() != defaultDbDriverName {
		t.Error("env.dbDriverName", env.dbDriverName)
	}

	if env.DbDataSource() != defaultDbDataSource {
		t.Error("env.dbDataSource", env.dbDataSource)
	}

	if env.FrontendPath() != defaultFrontendPath {
		t.Error("env.frontendPath", env.frontendPath)
	}

	if env.GraphHopperApiKey() != graphHopperApiKeyVar {
		t.Error("env.graphHopperApiKey", env.graphHopperApiKey)
	}

	if env.Develop() != true {
		t.Error("env.develop", env.develop)
	}

}

func Test_NewEnv_3(t *testing.T) {
	latLngBounds1 := "49.4126,11.0111;49.5118,11.2167"
	latLngBounds2 := "49.4126,11.0111 49.5118,11.2167"

	os.Setenv(openStreetMapBounds, latLngBounds1)

	env := NewEnv()

	if env.OpenStreetMapBounds() != latLngBounds1 {
		t.Error("expected:", latLngBounds1)
	}

	os.Setenv(openStreetMapBounds, latLngBounds2)

	env = NewEnv()

	if len(env.OpenStreetMapBounds()) != 0 {
		t.Error("expected no bounds")
	}
}

func TestNewEnv_4(t *testing.T) {
	storage := "https://some.url.de/path/"

	os.Setenv(assetStorage, storage)

	env := NewEnv()

	if env.AssetStorage() != storage[:len(storage)-1] {
		t.Error("expected asset storage to be set")
	}

	storage = "/xyz/"

	os.Setenv(assetStorage, storage)

	env = NewEnv()

	if len(env.assetStorage) != 0 {
		t.Error("expected asset storage string to be empty")
	}

	storage = "jklasdfja fasdklfjdghäöüüüüüüüüüü &&&&&& §§§§§"

	os.Setenv(assetStorage, storage)

	env = NewEnv()

	if len(env.assetStorage) != 0 {
		t.Error("expected asset storage string to be empty")
	}
}
