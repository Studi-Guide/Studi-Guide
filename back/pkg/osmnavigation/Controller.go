package osmnavigation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"studi-guide/pkg/utils"
)

var regex = "([0-9]+\\.?[0-9]+,[0-9]+\\.?[0-9]+)"

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

	startStr := context.Query("start")
	endStr := context.Query("end")

	fmt.Println(startStr, endStr)

	if match, err := regexp.MatchString(regex, startStr); err != nil || !match {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}
	if match, err := regexp.MatchString(regex, endStr); err != nil || !match {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "end does not match required format",
		})
		return
	}

	start := strings.Split(startStr, ",")
	end := strings.Split(endStr, ",")

	startLiteral, err := ParseLatLngLiteral(start[0], start[1])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}

	endLiteral, err := ParseLatLngLiteral(end[0], end[1])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}

	data, err := c.routeProvider.GetRoute(startLiteral, endLiteral, "en-US")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	context.Data(http.StatusOK, "application/json;charset=utf-8", data)
}
