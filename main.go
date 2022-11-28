package main

import (
	"app-hris-server/config"
	"app-hris-server/controller"
	"app-hris-server/repository"
	"app-hris-server/service"
	"log"
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

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService, jwtService)

	r := gin.Default()
	r.Static("assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
			auth.POST("/signin", authController.SigninWithUsername)
		}
	}
	log.Printf("connect to http://localhost:%s/", port)
	r.Run(":" + port)
}
