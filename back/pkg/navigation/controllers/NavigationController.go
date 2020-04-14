package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"studi-guide/pkg/navigation/services"
)

type NavigationController struct {
	router  *gin.RouterGroup
	service services.NavigationServiceProvider
}

func NewNavigationController(router *gin.RouterGroup, service services.NavigationServiceProvider) error {
	r := NavigationController{router: router, service: service}

	// TODO decide whether to use url routing or parameter query (google maps uses url routing)
	router.GET("/dir", r.GetNavigationRoute)
	return nil
}

// GetNavigationRoute godoc
// @Summary Get Navigation Route
// @Description Gets the navigation route for a start and end room
// @Accept  json
// @Produce  json
// @Tags NavigationController
// @Param startlocation query string false "the start location name"
// @Param endlocation query string false "the end location name"
// @Success 200 {array} navigation.Coordinate
// @Router /navigation/dir [get]
func (l *NavigationController) GetNavigationRoute(c *gin.Context) {
	startlocation := c.Query("startlocation")
	endlocation := c.Query("endlocation")

	log.Printf("[NavigationController] Request navigation received for start '%v' and end '%v'", startlocation, endlocation)
	coordinates, err := l.service.CalculateFromString(startlocation, endlocation)

	if err != nil {
		fmt.Println("GetAllRomms() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		fmt.Println(coordinates)
		c.JSON(http.StatusOK, coordinates)
	}
}
