package route

import (
	"github.com/gin-gonic/gin"
	"github.com/piTch-time/pitch-backend/application/controller"
)

// TaskRoutes is room api handler
func TaskRoutes(router *gin.RouterGroup, controller controller.TaskController) {
	tasks := router.Group("/rooms/:room_id/tasks")
	{
		tasks.POST("/", controller.Post())
		tasks.PATCH("/:task_id", controller.Patch())
	}
}
