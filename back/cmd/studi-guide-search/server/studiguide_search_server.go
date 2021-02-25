package searchnservice

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"studi-guide/pkg/building/campus"
	buildingInfo "studi-guide/pkg/building/info"
	buildingLocation "studi-guide/pkg/building/location"
	buildingMap "studi-guide/pkg/building/map"
	"studi-guide/pkg/building/room/controllers"
	buildingRoom "studi-guide/pkg/building/room/models"
	"studi-guide/pkg/env"
)

type SearchMicroService struct {
	router *gin.Engine
}

func NewSearchService(env *env.Env,
	buildingProvider buildingInfo.BuildingProvider, roomProvider buildingRoom.RoomServiceProvider,
	locationProvider buildingLocation.LocationProvider, mapProvider buildingMap.MapServiceProvider,
	campusProvider campus.CampusProvider) *SearchMicroService {
	log.Print("Starting initializing main controllers ...")
	router := gin.Default()

	if env.Develop() {
		log.Println("allowing all origins in develop mode")
		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
		}))
	}

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.GET("/", RedirectRootToAPI(router))
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	roomRouter := router.Group("/rooms")
	{
		log.Print("Creating room controllers")
		err := controllers.NewRoomController(roomRouter, roomProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized room controllers")
		}
		//a.Run(":8080")
	}

	locationRouter := router.Group("/locations")
	{
		log.Println("Creating location controller")
		err := buildingLocation.NewLocationController(locationRouter, locationProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized location controller")
		}
	}

	buildingRouter := router.Group("/buildings")
	{
		log.Println("Creating building controller")
		err := buildingInfo.NewBuildingController(buildingRouter, buildingProvider, roomProvider, locationProvider, mapProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized building controller")
		}
	}

	mapRouter := router.Group("/maps")
	{
		log.Print("Creating map controllers")
		err := buildingMap.NewMapController(mapRouter, mapProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized map controllers")
		}
	}

	campusRouter := router.Group("/campus")
	{
		log.Println("Creating campus controller")
		err := campus.NewCampusController(campusRouter, campusProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized campus controller")
		}
	}

	server := SearchMicroService{router: router}
	return &server
}

func (server *SearchMicroService) Start(port string) error {
	err := http.ListenAndServe(port, server.router)
	if err != nil {
		return err
	}
	return nil
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			// Authorization example
			// httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			//c.Abort()
		}

		c.Next()
	}
}
