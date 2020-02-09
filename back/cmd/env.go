package cmd

import (
	"log"
	"os"
)
type Env struct {
	dbDriverName, dbDataSource string
}

var dBDriverNameKey string = "DB_DRIVER_NAME"
var dbDataSourceKey string = "DB_DATA_SOURCE"

var env *Env

func GetEnv() (*Env) {

	if env == nil {
		env = &Env{os.Getenv(dBDriverNameKey), os.Getenv(dbDataSourceKey)}
	}


	if len(env.dbDriverName) == 0 && len(env.dbDataSource) == 0 {
		log.Println("Using sqlite3 DB driver as no environment variables were provided.")
		env.dbDriverName = "sqlite3"
		env.dbDataSource = "db.sqlite3"
	}

	return env
}

func (e Env) String() string {
	return dBDriverNameKey + "=" + e.dbDriverName + ";" +
		dbDataSourceKey + "=" + e.dbDataSource
}

func (e *Env) DbDriverName() string {
	return e.dbDriverName
}

func (e *Env) DbDataSource() string {
	return e.dbDataSource
}