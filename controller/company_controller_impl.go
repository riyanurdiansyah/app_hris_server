package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type CompanyControllerImpl struct {
	CompanyService service.CompanyService
	JWTService     service.JWTService
}

func NewCompanyController(service service.CompanyService, jwtService service.JWTService) CompanyController {
	return &CompanyControllerImpl{
		CompanyService: service,
		JWTService:     jwtService,
	}
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// InsertCompany implements CompanyController
func (controller *CompanyControllerImpl) InsertCompany(c *gin.Context) {
	request := dto.CompanyCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &request)
	secretKey := StringWithCharset(64, charset)

	request.SecretKey = secretKey

	createResponse := controller.CompanyService.InsertCompany(&request)
	if createResponse.Error {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: createResponse.Message,
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		responses := helper.DefaultResponse{
			Code:    http.StatusOK,
			Status:  true,
			Message: "New company has been added",
			Data:    createResponse,
		}
		c.JSON(http.StatusOK, responses)
	}
}
