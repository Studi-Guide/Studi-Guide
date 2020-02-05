package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"httpExample/pkg/roomcontroller"
	"httpExample/pkg/shoppinglist"
	"log"
	"net/http"
)

func StudiGuideServer() error {
	log.Print("Starting initializing main controllers ...")
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", RedirectRootToAPI(router))
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	shoppingRouter := router.Group("/shoppinglist")
	{
		shoppingRouter.Use(auth())
		log.Print("Creating shopping list")
		a := shoppinglist.ShoppingListApp{}
		a.Initialize(shoppingRouter,
			"/shoppinglist",
			"pkg/shoppinglist")

		//v1.GET("/users/:id", apis.GetUser)
	}

	roomRouter := router.Group("/roomlist")
	{
		log.Print("Creating room controller")
		roomController := roomcontroller.RoomControllerApp{}
		roomController.Initialize(roomRouter)
		//a.Run(":8080")
	}

	port := ":8080"
	log.Printf("Starting http listener on %s", port)
	log.Fatal(http.ListenAndServe(port, router))

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
		c.Request.URL.Path = "/shoppinglist/index" // <- this line
		r.HandleContext(c)
	}
}
