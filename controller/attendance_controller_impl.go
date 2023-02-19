package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"

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

// CheckIfDoneClockin implements AttendanceController
func (*AttendanceControllerImpl) CheckIfDoneClockin(absent *dto.ClockinCreateDTO) bool {
	panic("unimplemented")
}

// CheckIfDoneClockout implements AttendanceController
func (*AttendanceControllerImpl) CheckIfDoneClockout(absent *dto.ClockoutCreateDTO) bool {
	panic("unimplemented")
}

// Clockin implements AttendanceController
func (controller *AttendanceControllerImpl) Clockin(c *gin.Context) {

	request := dto.ClockinCreateDTO{}
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
				Message: "please check your file image",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			promoResponse := controller.AttendanceService.Clockin(&request)
			if promoResponse.Error {
				responses := helper.DefaultResponse{
					Code:    http.StatusBadRequest,
					Message: promoResponse.Message,
					Data:    helper.ObjectKosongResponse{},
					Status:  false,
				}
				c.JSON(http.StatusBadRequest, responses)
			} else {

				c.SaveUploadedFile(request.Image, promoResponse.ImageClockin)
				responses := helper.DefaultResponse{
					Code:    http.StatusCreated,
					Message: "Attendance has been added",
					Data:    promoResponse,
					Status:  true,
				}
				c.JSON(http.StatusOK, responses)
			}
		}
	}
}

// Clockout implements AttendanceController
func (*AttendanceControllerImpl) Clockout(c *gin.Context) {
	panic("unimplemented")
}
