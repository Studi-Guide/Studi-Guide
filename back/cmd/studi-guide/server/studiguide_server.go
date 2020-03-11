package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"studi-guide/pkg/env"
	"studi-guide/pkg/roomcontroller"
	"studi-guide/pkg/roomcontroller/models"
)

type StudiGuideServer struct {
	router *gin.Engine
}

func NewStudiGuideServer(env *env.Env, roomprovider models.RoomServiceProvider) *StudiGuideServer {
	log.Print("Starting initializing main controllers ...")
	router := gin.Default()

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
		log.Print("Creating room controller")
		roomController := roomcontroller.RoomControllerApp{}

		err := roomController.Initialize(roomprovider, roomRouter)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("Successfully initialized room controller")
		}
		//a.Run(":8080")
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/")
	})

	server := StudiGuideServer{router: router}
	return &server
}

func (server *StudiGuideServer) Start(port string) error {
	http.ListenAndServe(port, server.router)
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

// RedirectRootToAPI redirects all calls from root endpoint to current API documentation endpoint
func RedirectRootToAPI(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/shoppinglist/index.html" // <- this line
		r.HandleContext(c)
	}
}
