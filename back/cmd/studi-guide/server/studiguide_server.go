package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"net/url"
	"os"
	"studi-guide/pkg/building/campus"
	buildingInfo "studi-guide/pkg/building/info"
	buildingLocation "studi-guide/pkg/building/location"
	buildingMap "studi-guide/pkg/building/map"
	"studi-guide/pkg/building/room/controllers"
	buildingRoom "studi-guide/pkg/building/room/models"
	"studi-guide/pkg/env"
	"studi-guide/pkg/ion18n"
	navigation "studi-guide/pkg/navigation/controllers"
	"studi-guide/pkg/navigation/services"
	"studi-guide/pkg/osm"
	"studi-guide/pkg/rssFeed"
	"studi-guide/pkg/utils"
)

type StudiGuideServer struct {
	router *gin.Engine
}

func NewStudiGuideServer(env *env.Env,
	buildingProvider buildingInfo.BuildingProvider, roomProvider buildingRoom.RoomServiceProvider,
	locationProvider buildingLocation.LocationProvider, mapProvider buildingMap.MapServiceProvider,
	navigationProvider services.NavigationServiceProvider,
	campusProvider campus.CampusProvider,
	rssFeedProvider rssFeed.Provider,
	httpClient utils.HttpClient,
	osmNav osm.OpenStreetMapNavigationProvider) *StudiGuideServer {
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

	// TODO verify IONIC input
	_, err := os.Stat(env.FrontendPath())
	if err == nil {
		log.Print("IONIC found! Serving files using ion18n router....")

		ionRouter := router.Group("/")

		if _, err := ion18n.NewIon18nRouter(ionRouter, env.FrontendPath()); err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized ion18n router")
		}

	}

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

	navigationRouter := router.Group("/navigation")
	{
		log.Print("Creating navigation controllers")
		err := navigation.NewNavigationController(navigationRouter, navigationProvider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized navigation controllers")
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

	rssfeedRouter := router.Group("/rssfeed")
	{
		log.Println("Creating rss feed controller")
		err := rssFeed.NewRssFeedController(rssfeedRouter, rssFeedProvider, httpClient)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized rss feed controller")
		}
	}

	osmRouter := router.Group("/osm")
	{
		log.Println("Creating open street map controller")
		err := osm.NewOpenStreetMapController(osmRouter, osmNav, env)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized open street map controller")
		}
	}

	// redirect refresh to startpage
	router.GET("/tabs/navigation", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/")
	})

	router.NoRoute(func(c *gin.Context) {
		if env.Develop() {
			type ErrInfo struct {
				Status int
				Url    url.URL
				Header http.Header
				Proto  string
				Host   string
				Err    error
			}
			c.JSON(http.StatusNotFound, ErrInfo{
				Status: http.StatusNotFound,
				Url:    *c.Request.URL,
				Header: c.Request.Header,
				Proto:  c.Request.Proto,
				Host:   c.Request.Host,
				Err:    c.Err(),
			})
		} else {
			_, _ = c.Writer.WriteString(c.Request.URL.Path + " not found")
			c.Status(http.StatusNotFound)
		}
	})

	server := StudiGuideServer{router: router}
	return &server
}

func (server *StudiGuideServer) Start(port string) error {
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
