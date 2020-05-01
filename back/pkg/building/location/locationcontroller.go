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
	r.router.GET("", r.GetLocations)
	r.router.GET("/:location", r.GetLocationByName)
	return nil
}

// GetLocations godoc
// @Summary Query locations
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
// @Router /locations [get]
func (l LocationController) GetLocations(c *gin.Context) {

	building := c.Param("building")

	name := c.Query("name")
	tag := c.Query("tag")
	floor := c.Query("floor")
	campus := c.Query("campus")
	if len(building) == 0 {
		building = c.Query("building")
	}

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
// @Router /locations/{location} [get]
func (l LocationController) GetLocationByName(c *gin.Context) {
	buildingName := c.Param("building")
	locationName := c.Param("location") //mux.Vars(r)["name"]
	campusName := c.Param("campus")

	location, err := l.provider.GetLocation(locationName, buildingName, campusName)
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
