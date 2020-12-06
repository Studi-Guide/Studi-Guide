package osmnavigation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/utils"
)

type Controller struct {
	router        *gin.RouterGroup
	routeProvider OpenStreetMapNavigationProvider
	httpClient    utils.HttpClient
}

func NewOpenStreetMapController(router *gin.RouterGroup, provider OpenStreetMapNavigationProvider, client utils.HttpClient) error {
	b := Controller{
		router:        router,
		routeProvider: provider,
		httpClient:    client,
	}

	b.router.GET("/route", b.GetRoute)
	return nil
}

// Get Route for Open Street Map godoc
// @Summary Get Route for Open Street Map
// @Description Route for Open Street Map only possible for configured bounds
// @ID get-osmroute
// @Accept  json
// @Produce  plain
// @Tags OsmRouteController
// @Param start query string true "start point of route"
// @Param end query string true "end point of route"
// @Success 200
// @Router /osm/route [get]
func (c *Controller) GetRoute(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": "not implemented",
	})
}
