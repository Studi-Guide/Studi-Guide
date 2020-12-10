package osm

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"studi-guide/pkg/env"
	"studi-guide/pkg/locales"
	"studi-guide/pkg/osm/latlng"
)

type Controller struct {
	router        *gin.RouterGroup
	bounds        latlng.LatLngBounds
	routeProvider OpenStreetMapNavigationProvider
}

func NewOpenStreetMapController(router *gin.RouterGroup, provider OpenStreetMapNavigationProvider, env *env.Env) error {

	southWest, _ := latlng.NewLatLngLiteral(0, 0)
	northEast, _ := latlng.NewLatLngLiteral(0, 0)
	boundLiteral, _ := latlng.NewLatLngBounds(southWest, northEast)

	if len(env.OpenStreetMapBounds()) != 0 {

		bounds := strings.Split(env.OpenStreetMapBounds(), ";")
		a := strings.Split(bounds[0], ",")
		b := strings.Split(bounds[1], ",")

		southWest, err := latlng.ParseLatLngLiteral(a[0], a[1])
		if err != nil {
			return err
		}

		northEast, err := latlng.ParseLatLngLiteral(b[0], b[1])
		if err != nil {
			return err
		}

		boundLiteral, err = latlng.NewLatLngBounds(southWest, northEast)
		if err != nil {
			return err
		}
	}

	b := Controller{
		router:        router,
		bounds:        boundLiteral,
		routeProvider: provider,
	}

	b.router.GET("/route", b.GetRoute)
	b.router.GET("/bounds", b.GetBounds)

	return nil
}

// Get Bounds for Open Street Map godoc
// @Summary Get Bounds for Open Street Map
// @Description Bounds for Open Street Map
// @ID get-osmbounds
// @Accept  json
// @Produce  plain
// @Tags OsmRouteController
// @Success 200
// @Router /osm/bounds [get]
func (c *Controller) GetBounds(context *gin.Context) {
	context.JSON(http.StatusOK, c.bounds)
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
// @Param locale query string false "locale for route instructions"
// @Success @Success 200 {array} osm.Route
// @Router /osm/route [get]
func (c *Controller) GetRoute(context *gin.Context) {

	startStr := context.Query("start")
	endStr := context.Query("end")

	if match, err := regexp.MatchString(latlng.LatLngLiteralRegex, startStr); err != nil || !match {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}
	if match, err := regexp.MatchString(latlng.LatLngLiteralRegex, endStr); err != nil || !match {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "end does not match required format",
		})
		return
	}

	start := strings.Split(startStr, ",")
	end := strings.Split(endStr, ",")

	startLiteral, err := latlng.ParseLatLngLiteral(start[0], start[1])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}

	endLiteral, err := latlng.ParseLatLngLiteral(end[0], end[1])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "start does not match required format",
		})
		return
	}

	if !c.bounds.IncludeLiteral(startLiteral) || !c.bounds.IncludeLiteral(endLiteral) {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "start or end not included in navigation bounds",
		})
		return
	}

	locale := locales.GetBestSupportedLocale(context.Query("locale"))

	routes, err := c.routeProvider.GetRoute(startLiteral, endLiteral, locale)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	var data []byte
	data, err = json.Marshal(routes);
	if  err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}


	context.Data(http.StatusOK, "application/json;charset=utf-8", data)
}
