package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piTch-time/pitch-backend/application"
	"github.com/piTch-time/pitch-backend/domain/entity"
	"github.com/piTch-time/pitch-backend/domain/service"
	"github.com/piTch-time/pitch-backend/infrastructure/logger"
)

// TaskController ...
type TaskController interface {
	Post() gin.HandlerFunc
	Patch() gin.HandlerFunc
	// VerifyPassword() gin.HandlerFunc
}

type taskController struct {
	taskService service.TaskService
}

// NewTaskController ...
func NewTaskController(ts service.TaskService) TaskController {
	return &taskController{
		taskService: ts,
	}
}

type postRequestTask struct {
	CreatedBy   string `json:"createdBy"`
	Description string `json:"description"`
}

type postResponseTask struct {
	TaskID uint `json:"taskId"`
}

// @Summary      create a task
// @Description  태스크 생성
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        room_id  path 	int  true "room ID"
// @Param        room  body     postRequestTask  true  "태스크 생성요청 body"
// @Success      200  {object}   postResponseTask
// @Failure      400
// @Router       /rooms/{room_id}/tasks [post]
func (tc *taskController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postRequestTask
		if err := c.BindJSON(&req); err != nil {
			fmt.Println("1")
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		roomID, err := application.ParseUint(c.Param("room_id"))
		if err != nil {
			fmt.Println("2")
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		taskID, err := tc.taskService.Create(roomID, req.CreatedBy, req.Description)
		if err != nil {
			fmt.Println("3")
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := postResponseTask{TaskID: taskID}
		c.JSON(http.StatusOK, res)
	}
}

type patchRequestTask struct {
	IsDone      bool   `json:"isDone,omitempty"`
	Description string `json:"description,omitempty"`
}

func (t *patchRequestTask) ToEntity(task *entity.Task) *entity.Task {
	if t.IsDone {
		task.IsDone = t.IsDone
	}
	if t.Description != "" {
		task.Description = t.Description
	}
	fmt.Println(task)
	return task
}

type patchResponseTask struct {
	RoomID      uint   `json:"roomID"`
	CreatedBy   string `json:"createdBy"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

// @Summary      update a task
// @Description  태스크 업데이트
// @Description  1. 목표를 클릭 했을 때 task의 isDone을 변경해달라고 요청 시 사용
// @Description  2. 목표의 내용을 수정하고 싶을 때 description 수정
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        room_id  path 	int  true "room ID"
// @Param        task_id  path 	int  true "task ID"
// @Param        task  body 	patchRequestTask  true "태스크 수정 요청 body"
// @Success      200  {object}   patchResponseTask
// @Failure      400
// @Router       /rooms/{room_id}/tasks/{task_id} [patch]
func (tc *taskController) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req patchRequestTask
		var taskID uint
		var err error
		var task *entity.Task
		var updatedTask *entity.Task

		if err := c.BindJSON(&req); err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := application.ParseUint(c.Param("room_id")); err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		taskID, err = application.ParseUint(c.Param("task_id"))
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task, err = tc.taskService.Get(taskID)

		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedTask, err = tc.taskService.Update(req.ToEntity(task))
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := patchResponseTask{RoomID: updatedTask.RoomID,
			CreatedBy:   updatedTask.CreatedBy,
			Description: updatedTask.Description,
			IsDone:      updatedTask.IsDone}
		c.JSON(http.StatusOK, res)
	}
}
