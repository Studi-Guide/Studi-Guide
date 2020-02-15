package roomcontroller

import (
	"github.com/gin-gonic/gin"
	"studi-guide/pkg/env"
	"studi-guide/pkg/roomcontroller/controllers"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomControllerApp struct {
	roomcontroller *controllers.RoomController
}

func (r *RoomControllerApp) Initialize(env *env.Env, router *gin.RouterGroup) (error) {

	provider, err := models.NewRoomDbService(env.DbDriverName(), env.DbDataSource(), "rooms")
	if err != nil {
		return err
	}

	r.roomcontroller = controllers.NewRoomController(router, provider)
	return nil
}
