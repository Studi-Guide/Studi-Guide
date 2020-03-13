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
// @Param startroom query string false "the start room name"
// @Param endroom query string false "the end room name"
// @Success 200 {array} navigation.Coordinate
// @Router /navigation/dir [get]
func (l *NavigationController) GetNavigationRoute(c *gin.Context) {
	log.Print("[NavigationController] Request navigation received")
	startroom := c.Query("startroom")
	endroom := c.Query("endroom")

	coordinates, err := l.service.CalculateFromString(startroom, endroom)

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
