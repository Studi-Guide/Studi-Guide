package env

import (
	"log"
	"os"
)

type Env struct {
	dbDriverName, dbDataSource, frontendPath string
	develop bool
}

var dBDriverNameKey string = "DB_DRIVER_NAME"
var dbDataSourceKey string = "DB_DATA_SOURCE"
var frontendPath string = "FRONTEND_PATH"
var develop string = "DEVELOP"

func NewEnv() *Env {

	env := &Env{os.Getenv(dBDriverNameKey), os.Getenv(dbDataSourceKey), os.Getenv(frontendPath), false}

	if len(env.dbDriverName) == 0 && len(env.dbDataSource) == 0 {
		log.Println("Using sqlite3 DB driver as no environment variables were provided.")
		env.dbDriverName = "sqlite3"
		env.dbDataSource = "db.sqlite3"
	}

	if len(env.frontendPath) == 0 {
		log.Print("Using default frontend path ...")
		env.frontendPath = "./ionic"
	}

	if (os.Getenv(develop) == "TRUE") {
		log.Println("Running in development mode now. Make sure to disable this in production.")
		env.develop = true
	}

	return env
}

func (e Env) String() string {
	return dBDriverNameKey + "=" + e.dbDriverName + ";" +
		dbDataSourceKey + "=" + e.dbDataSource + ";" +
		frontendPath + "=" + e.frontendPath
}

func (e *Env) DbDriverName() string {
	return e.dbDriverName
}

func (e *Env) DbDataSource() string {
	return e.dbDataSource
}

func (e *Env) FrontendPath() string {
	return e.frontendPath
}

func (e *Env) Develop() bool {
	return e.develop
}