package ion18n

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
	"studi-guide/pkg/env"
	"studi-guide/pkg/locales"
)

type Ion18nRouter struct {
	router *gin.RouterGroup
	env    *env.Env
}

func NewIon18nRouter(router *gin.RouterGroup, env *env.Env) (*Ion18nRouter, error) {
	i := Ion18nRouter{router: router, env: env}

	i.router.GET("", i.HandleRequest)

	// somehow this does not work with gin.RouterGroup
	// is there any advantage of gin-contrib/static over the default router.Static?
	//i.router.Use(static.Serve("/en-US/", static.LocalFile(i.ionPath + "/en-US/", true)))
	//i.router.Use(static.Serve("/de/", static.LocalFile(i.ionPath + "/de/", true)))

	for _, tag := range locales.GetSupportedLocales() {
		base, _ := tag.Base()
		i.router.Static(base.String()+"/", i.env.FrontendPath()+"/"+base.String()+"/")
	}

	return &i, nil
}

func (i *Ion18nRouter) HandleRequest(c *gin.Context) {

	bestMatch := locales.GetBestSupportedLocale(c.GetHeader("Accept-Language"))

	c.Redirect(http.StatusTemporaryRedirect, "/"+bestMatch+"/")

}

func (i *Ion18nRouter) HandleNoRoute(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
	for _, tag := range locales.GetSupportedLocales() {
		base, _ := tag.Base()
		if strings.HasPrefix(c.Request.URL.Path, "/"+base.String()) {
			c.File(i.env.FrontendPath() + "/" + base.String() + "/index.html")
			return
		}
	}

	i.HandleNotFound(c)
}

func (i *Ion18nRouter) HandleNotFound(c *gin.Context) {
	if i.env.Develop() {
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
}
