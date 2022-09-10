package main

import (
	"app-ecommerce-server/config"
	"app-ecommerce-server/controller"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/middleware"
	"app-ecommerce-server/repository"
	"app-ecommerce-server/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	validate := validator.New()
	db := config.SetupGetConnection()
	jwtService := service.NewJWTService()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService, jwtService)

	promoRepository := repository.NewPromoRepository()
	promoService := service.NewPromoService(promoRepository, db, validate)
	promoController := controller.NewPromoController(promoService)

	r := gin.Default()
	r.Static("assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.Request.Header.Add("Access-Control-Allow-Origin", "*")
		c.Request.Header.Add("Access-Control-Allow-Headers", "*")
		// c.File("index.html")
		c.JSON(200, helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "WKWKWKW",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		})
	})
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
			auth.POST("/signin", authController.SigninWithUsername)
		}
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryController.InsertCategory)
			categories.GET("", categoryController.FindAllCategory)
			categories.GET("/:id", categoryController.FindByIdCategory)
			categories.PUT("", categoryController.UpdateCategory)
			categories.DELETE("/:id", categoryController.DeleteCategory)
		}
		promos := v1.Group("/promos", middleware.AuthorizeJWT(jwtService))
		{
			promos.POST("", promoController.InsertPromo)
			promos.GET("", promoController.GetAllPromo)
			promos.PUT("", promoController.UpdatePromo)
			promos.DELETE("/:id", promoController.DeletePromo)
		}
	}
	log.Printf("connect to http://localhost:%s/", port)
	r.Run(":" + port)
}
