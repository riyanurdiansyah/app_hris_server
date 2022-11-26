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
		c.File("index.html")
	})
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
			auth.POST("/signin", authController.SigninWithUsername)
		}
		company := v1.Group("/company", middleware.AuthorizeJWT(jwtService))
		{
			company.POST("", companyController.InsertCompany)
		}
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryController.InsertCategory)
			categories.GET("", categoryController.FindAllCategory)
			categories.GET("/:id", categoryController.FindByIdCategory)
			categories.PUT("", categoryController.UpdateCategory)
			categories.DELETE("/:id", categoryController.DeleteCategory)
		}
		//, middleware.AuthorizeJWT(jwtService)
		promos := v1.Group("/promos")
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
