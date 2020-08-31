package ion18n

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Ion18nRouter struct {
	router   *gin.RouterGroup
	ionPath  string
}

func NewIon18nRouter(router *gin.RouterGroup, ionPath string) (*Ion18nRouter, error) {
	i := Ion18nRouter{router: router, ionPath: ionPath}

	i.router.GET("", i.HandleRequest)

	// somehow this does not work with gin.RouterGroup
	// is there any advantage of gin-contrib/static over the default router.Static?
	//i.router.Use(static.Serve("/en-US/", static.LocalFile(i.ionPath + "/en-US/", true)))
	//i.router.Use(static.Serve("/de/", static.LocalFile(i.ionPath + "/de/", true)))
	i.router.Static("en-US/", i.ionPath + "/en-US/")
	i.router.Static("de/", i.ionPath + "/de/")

	return &i, nil
}

func (i *Ion18nRouter) HandleRequest(c *gin.Context) {

	if strings.Contains(c.GetHeader("Accept-Language"), "de") {
		c.Redirect(http.StatusTemporaryRedirect, "/de/")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/en-US/")
	}

}