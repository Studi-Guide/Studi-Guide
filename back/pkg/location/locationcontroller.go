package location

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/entityservice"
)

type LocationController struct {
	router   *gin.RouterGroup
	provider LocationProvider
}

func NewLocationController(router *gin.RouterGroup, provider LocationProvider) error {
	r := LocationController{router: router, provider: provider}
	r.router.GET("/", r.GetLocations)
	r.router.GET("/name/:name", r.GetLocationByName)
	return nil
}

// GetLocations godoc
// @Summary Get All Locations
// @Description Gets locations by possible filters
// @Accept  json
// @Produce  json
// @Tags LocationController
// @Param name query string false "name of the location"
// @Param tag query string false "a tag of the location"
// @Param floor query string false "floor of the location"
// @Param campus query string false "campus of the location"
// @Param building query string false "building of the location"
// @Success 200 {array} entityservice.Location
// @Router /location [get]
func (l LocationController) GetLocations(c *gin.Context) {
	name := c.Query("name")
	tag := c.Query("tag")
	floor := c.Query("floor")
	campus := c.Query("campus")
	building := c.Query("building")

	var locations []entityservice.Location
	var err error

	var useFilterApi bool

	if len(name) == 0 && len(tag) == 0 && len(floor) == 0 && len(building) == 0 && len(campus) == 0 {
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		locations, err = l.provider.FilterLocations(name, tag, floor, building, campus)
	} else {
		locations, err = l.provider.GetAllLocations()
	}

	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, locations)
}

// GetLocationByName godoc
// @Summary Get Location by a certain name
// @Description Get one location by name
// @ID get-location-name
// @Accept  json
// @Produce  json
// @Tags LocationController
// @Param name path int true "get location by name"
// @Success 200 {array} entityservice.Location
// @Router /location/name/{name} [get]
func (l LocationController) GetLocationByName(c *gin.Context) {
	name := c.Param("name") //mux.Vars(r)["name"]

	location, err := l.provider.GetLocation(name)
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, location)
}
