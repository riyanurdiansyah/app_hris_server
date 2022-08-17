package controller

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// SignUp implements AuthController
func (controller *AuthControllerImpl) SignUp(c *gin.Context) {
	userCreateRequest := dto.UserCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &userCreateRequest)

	userCreateResponse := controller.AuthService.SignUp(c, &userCreateRequest)
	if userCreateResponse.Error {
		responses := helper.DefaultResponse{
			Code:   http.StatusBadRequest,
			Status: userCreateResponse.Message,
			Data:   helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		responses := helper.DefaultResponse{
			Code:   http.StatusOK,
			Status: "New user has been added",
			Data:   userCreateResponse,
		}
		c.JSON(http.StatusOK, responses)
	}
}
