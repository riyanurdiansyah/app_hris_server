package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigControllerImpl struct {
	ConfigService service.ConfigService
	JWTService    service.JWTService
}

func NewConfigController(configService service.ConfigService, jwtService service.JWTService) ConfigController {
	return &ConfigControllerImpl{
		ConfigService: configService,
		JWTService:    jwtService,
	}
}

// CheckConfig implements ConfigController
func (controller *ConfigControllerImpl) CheckConfig(name string) bool {
	return controller.ConfigService.CheckConfig(name)
}

// GetConfig implements ConfigController
func (controller *ConfigControllerImpl) GetConfig(c *gin.Context) {
	userResponse := controller.ConfigService.GetConfig()
	responses := helper.DefaultResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Data has been listed",
		Data:    userResponse,
	}
	c.JSON(http.StatusOK, responses)
}

// GetConfigByName implements ConfigController
func (controller *ConfigControllerImpl) GetConfigByName(c *gin.Context) {
	configName := c.Param("name")
	cekConfig := controller.ConfigService.CheckConfig(configName)
	if cekConfig {

		userResponse := controller.ConfigService.GetConfigByName(configName)
		responses := helper.DefaultResponse{
			Code:    http.StatusOK,
			Status:  true,
			Message: "Data has been listed",
			Data:    userResponse,
		}
		c.JSON(http.StatusOK, responses)
	} else {

		responses := helper.DefaultResponse{
			Code:    http.StatusForbidden,
			Message: "Config is not found",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusForbidden, responses)
	}
}

// InsertConfig implements ConfigController
func (controller *ConfigControllerImpl) InsertConfig(c *gin.Context) {
	configRequest := dto.ConfigCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &configRequest)
	checkUser := controller.CheckConfig(configRequest.Name)
	if !checkUser {
		userResponse := controller.ConfigService.InsertConfig(&configRequest)
		if userResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: userResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusCreated,
				Status:  true,
				Message: "Data has been created",
				Data:    userResponse,
			}
			c.JSON(http.StatusCreated, responses)
		}
	} else {

		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Config already registered",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	}
}

// UpdateConfig implements ConfigController
func (controller *ConfigControllerImpl) UpdateConfig(c *gin.Context) {
	configRequest := dto.ConfigUpdateDTO{}
	helper.ReadFromRequestBody(c.Request, &configRequest)
	checkUser := controller.CheckConfig(configRequest.Name)
	if checkUser {
		userResponse := controller.ConfigService.UpdateConfig(&configRequest)
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
	} else {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Config is not registered",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	}
}
