package route

import (
	"github.com/gin-gonic/gin"
	"github.com/piTch-time/pitch-backend/application/controller"
)

// RoomRoutes is room api handler
func RoomRoutes(router *gin.RouterGroup, controller controller.RoomController) {
	rooms := router.Group("/rooms")
	{
		rooms.GET("/", controller.GetAll())
		rooms.GET("/:room_id", controller.Get())
		rooms.POST("/", controller.Post())
		rooms.DELETE("/:room_id", controller.Delete())
	}
}
