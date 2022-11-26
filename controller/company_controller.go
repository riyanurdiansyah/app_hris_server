package controller

import "github.com/gin-gonic/gin"

type CompanyController interface {
	InsertCompany(c *gin.Context)
}
