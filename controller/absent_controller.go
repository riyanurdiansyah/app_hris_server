package controller

import (
	"app-hris-server/data/dto"

	"github.com/gin-gonic/gin"
)

type AbsentController interface {
	InsertAbsent(c *gin.Context)
	CheckIfDoneAbsent(absent *dto.AbsentCreateDTO) bool
}
