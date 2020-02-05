package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"httpExample/pkg/roomcontroller/models"
	"io/ioutil"
	"log"
	"net/http"
)

type RoomController struct {
	router   *gin.RouterGroup
	roomList []models.Room
}

func NewRoomController(router *gin.RouterGroup, roomList []models.Room) (*RoomController) {
	r := RoomController{router: router, roomList: roomList}
	r.router.GET("/", r.GetRoomList)
	return &r
}

// GetRoomList godoc
// @Summary Get Room List
// @Description Gets all available rooms
// @ID get-room-list
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Success 200 {array} Room
// @Router /roomlist/ [get]
func (l *RoomController) GetRoomList(c *gin.Context) {
	log.Print("[RoomController] Request RoomList received")
	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message": l.roomList,// cast it to string before showing
	})
}

func (l *RoomController) GetItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(r)["name"]

	for _, item := range l.roomList {
		if item.Name == name {
			c.JSON(http.StatusOK ,item)
		}
	}
}

func (l *RoomController) AddItem(c *gin.Context) {
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	var item models.Room
	json.Unmarshal(reqBody, &item)

	l.roomList = append(l.roomList, item)
}

func (l *RoomController) RemoveItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(r)["name"]
	for index, item := range l.roomList {
		if item.Name == name {
			l.roomList = append(l.roomList[:index], l.roomList[index+1:]...)
		}
	}
}