package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuControllerImpl struct {
	MenuService service.MenuService
	JWTService  service.JWTService
}

func NewMenuController(menuService service.MenuService, jwtService service.JWTService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService,
		JWTService:  jwtService,
	}
}

// CheckMenu implements MenuController
func (controller *MenuControllerImpl) CheckMenu(id int) bool {
	return controller.MenuService.CheckMenu(id)
}

// InsertMenu implements MenuController
func (controller *MenuControllerImpl) InsertMenu(c *gin.Context) {
	menuCreateRequest := dto.MenuCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &menuCreateRequest)
	userResponse := controller.MenuService.InsertMenu(&menuCreateRequest)
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
}

// UpdateMenu implements MenuController
func (controller *MenuControllerImpl) UpdateMenu(c *gin.Context) {
	menuCreateRequest := dto.MenuUpdateDTO{}
	helper.ReadFromRequestBody(c.Request, &menuCreateRequest)

	checkUser := controller.CheckMenu(menuCreateRequest.Id)
	if !checkUser {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "id is not found",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		userResponse := controller.MenuService.UpdateMenu(&menuCreateRequest)
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
