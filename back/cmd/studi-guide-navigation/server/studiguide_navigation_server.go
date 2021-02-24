package navigationservice

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"studi-guide/pkg/env"
	navigation "studi-guide/pkg/navigation/controllers"
	"studi-guide/pkg/navigation/services"
)

type NavigationMicroService struct {
	router *gin.Engine
}

func NewNavigationService(env *env.Env,
	navigationProvider services.NavigationServiceProvider) *NavigationMicroService {
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

	server := NavigationMicroService{router: router}
	return &server
}

func (server *NavigationMicroService) Start(port string) error {
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
