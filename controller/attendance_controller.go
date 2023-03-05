package controller

import (
	"github.com/gin-gonic/gin"
)

type AttendanceController interface {
	Attendance(c *gin.Context)
}
