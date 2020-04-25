package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"studi-guide/pkg/building"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/env"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
	navigation "studi-guide/pkg/navigation/controllers"
	"studi-guide/pkg/navigation/services"
	"studi-guide/pkg/roomcontroller/controllers"
)

type StudiGuideServer struct {
	router *gin.Engine
}

func NewStudiGuideServer(env *env.Env, entityService *entityservice.EntityService, navigationprovider services.NavigationServiceProvider) *StudiGuideServer {
	log.Print("Starting initializing main controllers ...")
	router := gin.Default()

	if env.Develop() {
		log.Println("allowing all origins in develop mode")
		router.Use(cors.New(cors.Config{
			AllowAllOrigins:        true,
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
	if _, err := os.Stat(env.FrontendPath()); err == nil {
		log.Print("IONIC found! Serving files ....")
		router.Use(static.Serve("/", static.LocalFile(env.FrontendPath(), true)))
	}

	roomRouter := router.Group("/roomlist")
	{
		log.Print("Creating room controllers")
		err := controllers.NewRoomController(roomRouter, entityService)
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
		err := navigation.NewNavigationController(navigationRouter, navigationprovider)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized navigation controllers")
		}
	}

	mapRouter := router.Group("/map")
	{
		log.Print("Creating map controllers")
		err := maps.NewMapController(mapRouter, entityService)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized map controllers")
		}
	}

	locationRouter := router.Group("/location")
	{
		log.Println("Creating location controller")
		err := location.NewLocationController(locationRouter, entityService)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized location controller")
		}
	}

	buildingRouter := router.Group("/buildings")
	{
		log.Println("Creating building controller")
		err := building.NewBuildingController(buildingRouter, entityService)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully initialized building controller")
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/")
	})

	server := StudiGuideServer{router: router}
	return &server
}

func (server *StudiGuideServer) Start(port string) error {
	error := http.ListenAndServe(port, server.router)
	if error != nil{
		return error
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
