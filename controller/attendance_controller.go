package controller

import (
	"app-hris-server/data/dto"

	"github.com/gin-gonic/gin"
)

type AttendanceController interface {
	Clockin(c *gin.Context)
	Clockout(c *gin.Context)
	CheckIfDoneClockin(absent *dto.ClockinCreateDTO) bool
	CheckIfDoneClockout(absent *dto.ClockoutCreateDTO) bool
}
