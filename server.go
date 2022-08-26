package main

import (
	"app-ecommerce-server/config"
	"app-ecommerce-server/controller"
	"app-ecommerce-server/repository"
	"app-ecommerce-server/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func main() {
	validate := validator.New()
	db := config.SetupGetConnection()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	jwtService := service.NewJWTService()
	authController := controller.NewAuthController(authService, jwtService)

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

		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
		}
	}
	r.Run()
}
