package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
	JWTService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &UserControllerImpl{
		UserService: userService,
		JWTService:  jwtService,
	}
}

// AddUserInfoPersonal implements UserController
func (controller *UserControllerImpl) AddUserInfoPersonal(c *gin.Context) {
	userCreateRequest := dto.UserInfoCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &userCreateRequest)
	checkUser := controller.CheckUser(userCreateRequest.IdEmployee)
	if !checkUser {
		userResponse := controller.UserService.AddUserInfoPersonal(&userCreateRequest)
		if userResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: userResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Status:  true,
				Message: "Data has been created",
				Data:    userResponse,
			}
			c.JSON(http.StatusOK, responses)
		}
	} else {
		userResponse := controller.UserService.UpdateUserInfoPersonal(&userCreateRequest)
		if userResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: userResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Status:  true,
				Message: "Data has been updated",
				Data:    userResponse,
			}
			c.JSON(http.StatusOK, responses)
		}
	}

}

// UpdateUserInfoPersonal implements UserController
func (controller *UserControllerImpl) UpdateUserInfoPersonal(c *gin.Context) {
	userCreateRequest := dto.UserInfoCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &userCreateRequest)

	checkUser := controller.CheckUser(userCreateRequest.IdEmployee)
	if !checkUser {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Employee is not registered",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		userResponse := controller.UserService.UpdateUserInfoPersonal(&userCreateRequest)
		if userResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: userResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Status:  true,
				Message: "Data has been updated",
				Data:    userResponse,
			}
			c.JSON(http.StatusOK, responses)
		}
	}
}

// CheckUser implements UserController
func (controller *UserControllerImpl) CheckUser(employeeId string) bool {
	return controller.UserService.CheckUser(employeeId)
}
