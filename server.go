package main

import (
	"belajar/config"
	"belajar/controller"
	"belajar/repository"
	"belajar/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	validate := validator.New()
	db := config.SetupGetConnection()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryController.InsertCategory)
			categories.GET("", categoryController.FindAllCategory)
			categories.GET("/:id", categoryController.FindByIdCategory)
			categories.PUT("", categoryController.UpdateCategory)
			categories.DELETE("/:id", categoryController.DeleteCategory)
		}
	}
	r.Run()
}
