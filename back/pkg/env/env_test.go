package env

import (
	"os"
	"testing"
)

func Test_NewEnv_1(t *testing.T) {

	dBDriverNameVar := "A"
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
