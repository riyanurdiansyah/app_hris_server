package main

import (
	"app-hris-server/config"
	"app-hris-server/controller"
	"app-hris-server/middleware"
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

	companyRepository := repository.NewCompanyRepository()
	companyService := service.NewCompanyService(companyRepository, db, validate)
	companyController := controller.NewCompanyController(companyService, jwtService)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService, jwtService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService, jwtService)

	// menuRepository := repository.NewMenuRepository()
	// menuService := service.NewMenuService(menuRepository, db, validate)
	// menuController := controller.NewMenuController(menuService, jwtService)

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
			auth.POST("/signin", authController.SigninWithEmail)
		}
		company := v1.Group("/company")
		{
			company.POST("", companyController.InsertCompany)
		}

		userInfo := v1.Group("/user-info", middleware.AuthorizeJWT(jwtService))
		{
			userInfo.POST("", userController.AddUserInfoPersonal)
			userInfo.PUT("", userController.UpdateUserInfoPersonal)
		}

		// menu := v1.Group("/menu", middleware.AuthorizeJWT(jwtService))
		// {
		// 	userInfo.POST("", menuController.InsertMenu)
		// 	userInfo.PUT("", menuController.UpdateMenu)
		// }
	}
	log.Printf("connect to http://localhost:%s/", port)
	r.Run(":" + port)
}
