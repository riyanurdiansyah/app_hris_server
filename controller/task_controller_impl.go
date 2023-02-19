package controller

import (
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskControllerImpl struct {
	TaskService service.TaskService
	JWTService  service.JWTService
}

func NewTaskController(taskService service.TaskService, jwtService service.JWTService) TaskController {
	return &TaskControllerImpl{
		TaskService: taskService,
		JWTService:  jwtService,
	}
}

// CheckTask implements TaskController
func (*TaskControllerImpl) CheckTask(idTask int) bool {
	panic("unimplemented")
}

// GetTaskByUserId implements TaskController
func (controller *TaskControllerImpl) GetTaskByUserId(c *gin.Context) {
	userId := c.Param("id")
	// id, err := strconv.Atoi(userId)
	if userId == "" {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "User id is required",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {

		userResponse := controller.TaskService.GetTaskByUserId(userId)
		responses := helper.DefaultResponse{
			Code:    http.StatusOK,
			Status:  true,
			Message: "Data has been listed",
			Data:    userResponse,
		}
		c.JSON(http.StatusOK, responses)
	}

}

// InsertTask implements TaskController
func (*TaskControllerImpl) InsertTask(c *gin.Context) {
	panic("unimplemented")
}

// UpdateTask implements TaskController
func (*TaskControllerImpl) UpdateTask(c *gin.Context) {
	panic("unimplemented")
}
