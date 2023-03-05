package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AttendanceControllerImpl struct {
	AttendanceService service.AttendanceService
	JWTService        service.JWTService
}

func NewAttendanceController(attendanceService service.AttendanceService, jwtService service.JWTService) AttendanceController {
	return &AttendanceControllerImpl{
		AttendanceService: attendanceService,
		JWTService:        jwtService,
	}
}

// Attendance implements AttendanceController
func (controller *AttendanceControllerImpl) Attendance(c *gin.Context) {
	request := dto.AttendanceCreateDTO{}
	err := c.ShouldBind(&request)

	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "please check your image file",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		errBind := c.ShouldBindUri(&request)
		if errBind != nil {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: "please check your file images",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			checkAttendance := controller.AttendanceService.CheckAttendance(&request)
			if checkAttendance {
				responses := helper.DefaultResponse{
					Code:    http.StatusBadRequest,
					Message: "Kamu sudah melakukan absen",
					Data:    helper.ObjectKosongResponse{},
					Status:  false,
				}
				c.JSON(http.StatusBadRequest, responses)
			} else {
				promoResponse := controller.AttendanceService.Attendance(&request)
				if strings.ToLower(request.Kode) == "clockin" {
					c.SaveUploadedFile(request.Image, promoResponse.ImageClockin)
				} else {
					c.SaveUploadedFile(request.Image, promoResponse.ImageClockout)
				}
				responses := helper.DefaultResponse{
					Code:    http.StatusCreated,
					Message: "Data has been recorded",
					Data:    promoResponse,
					Status:  true,
				}
				c.JSON(http.StatusOK, responses)
			}
		}
	}
}
