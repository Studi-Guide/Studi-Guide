package env

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"regexp"
	"studi-guide/pkg/osm/latlng"
)

type Env struct {
	dbDriverName, dbDataSource, frontendPath string
	graphHopperApiKey, openStreetMapBounds   string
	develop                                  bool
}

var dBDriverNameKey string = "DB_DRIVER_NAME"
var dbDataSourceKey string = "DB_DATA_SOURCE"
var frontendPath string = "FRONTEND_PATH"
var graphHopperApiKey = "GRAPHHOPPER_API_KEY"
var openStreetMapBounds = "OPENSTREETMAP_BOUNDS"
var develop string = "DEVELOP"

var defaultDbDriverName = "sqlite3"
var defaultDbDataSource = "db.sqlite3"
var defaultFrontendPath = "./ionic"

func NewEnv() *Env {

	env := &Env{
		dbDriverName:      os.Getenv(dBDriverNameKey),
		dbDataSource:      os.Getenv(dbDataSourceKey),
		frontendPath:      os.Getenv(frontendPath),
		graphHopperApiKey: os.Getenv(graphHopperApiKey),
		openStreetMapBounds: os.Getenv(openStreetMapBounds),
		develop:           false,
	}

	if len(env.dbDriverName) == 0 && len(env.dbDataSource) == 0 {
		log.Println("Using sqlite3 DB driver as no environment variables were provided.")
		env.dbDriverName = defaultDbDriverName
		env.dbDataSource = defaultDbDataSource
	}

	if len(env.frontendPath) == 0 {
		log.Println("Using default frontend path ...")
		env.frontendPath = defaultFrontendPath
	}

	if len(env.graphHopperApiKey) == 0 {
		log.Println("No Graphhopper API key provided. Openstreetmap route calculation will not be possible.")
	}

	if len(env.openStreetMapBounds) == 0 {
		log.Println("No OpenStreetMap bounds were given! Make sure to provide bounds via environment variables in production.")
	} else {
		// check bounds format
		regexStr := latlng.LatLngLiteralRegex +";"+ latlng.LatLngLiteralRegex
		if match, err := regexp.MatchString(regexStr, env.openStreetMapBounds); err != nil || !match {
			env.openStreetMapBounds = ""
			log.Println("Regex match of OpenStreetMap bounds failed. Check your configuration. Bounds for OpenStreetMap are now disabled.")
		}
	}

	if os.Getenv(develop) == "TRUE" {
		log.Println("Running in development mode now. Make sure to disable this in production.")
		env.develop = true
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return env
}

func (e Env) String() string {
	return dBDriverNameKey + "=" + e.dbDriverName + ";" +
		dbDataSourceKey + "=" + e.dbDataSource + ";" +
		frontendPath + "=" + e.frontendPath + ";" +
		graphHopperApiKey + "=" + e.graphHopperApiKey
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

func (e *Env) GraphHopperApiKey() string {
	return e.graphHopperApiKey
}

func (e *Env) OpenStreetMapBounds() string {
	return e.openStreetMapBounds
}

func (e *Env) Develop() bool {
	return e.develop
}
