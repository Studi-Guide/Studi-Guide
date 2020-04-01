package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomController struct {
	router   *gin.RouterGroup
	provider models.RoomServiceProvider
}

func NewRoomController(router *gin.RouterGroup, provider models.RoomServiceProvider) error {
	r := RoomController{router: router, provider: provider}
	r.router.GET("/", r.GetRoomList)
	r.router.GET("/name", r.GetRoom)
	r.router.GET("/floor", r.GetRoomListFromFloor)
	return nil
}

// GetRoomList godoc
// @Summary Get Room List
// @Description Gets all available rooms
// @ID get-room-list
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Success 200 {array} models.Room
// @Router /roomlist/ [get]
func (l *RoomController) GetRoomList(c *gin.Context) {
	rooms, err := l.provider.GetAllRooms()
	if err != nil {
		fmt.Println("GetAllRomms() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		fmt.Println(rooms)
		c.JSON(http.StatusOK, rooms)
	}
}

func (l *RoomController) GetRoom(c *gin.Context) {
	name := c.Query("name") //mux.Vars(r)["name"]

	room, err := l.provider.GetRoom(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, room)
	}
}

// GetRoomListFromFloor godoc
// @Summary Get Room List From Floor
// @Description Gets all available rooms for a certain floor
// @ID get-room-list-floor
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Param floor query int false "filter rooms by floor"
// @Success 200 {array} models.Room
// @Router /roomlist/floor [get]
func (l *RoomController) GetRoomListFromFloor(c *gin.Context) {
	floor := c.Query("floor")

	floorInt, err := strconv.Atoi(floor)
	if err != nil {
		// handle error
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	rooms, err := l.provider.GetRoomsFromFloor(floorInt)
	if err != nil {
		fmt.Println("GetRoomListFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		fmt.Println(rooms)
		c.JSON(http.StatusOK, rooms)
	}
}

//func (l *RoomController) AddItem(c *gin.Context) {
//	reqBody, _ := ioutil.ReadAll(c.Request.Body)
//	var item models.Room
//	json.Unmarshal(reqBody, &item)
//
//	l.roomList = append(l.roomList, item)
//}
//
//func (l *RoomController) RemoveItem(c *gin.Context) {
//	name := c.Param("name") //mux.Vars(r)["name"]
//	for index, item := range l.roomList {
//		if item.Name == name {
//			l.roomList = append(l.roomList[:index], l.roomList[index+1:]...)
//		}
//	}
//}
