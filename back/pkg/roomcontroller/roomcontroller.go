package roomcontroller

import (
	"github.com/gin-gonic/gin"
	"studi-guide/pkg/roomcontroller/controllers"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomControllerApp struct {
	roomcontroller *controllers.RoomController
}

func (r *RoomControllerApp) Initialize(provider models.RoomServiceProvider, router *gin.RouterGroup) error {
	r.roomcontroller = controllers.NewRoomController(router, provider)
	return nil
}
