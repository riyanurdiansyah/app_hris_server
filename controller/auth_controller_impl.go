package controller

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		JWTService:  jwtService,
	}
}

// SignUp implements AuthController
func (controller *AuthControllerImpl) SignUp(c *gin.Context) {
	userCreateRequest := dto.UserCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &userCreateRequest)

	checkEmail := controller.CheckEmail(c, userCreateRequest.Email)
	checkUsername := controller.CheckUsername(c, userCreateRequest.Username)
	if checkEmail {
		responses := helper.DefaultResponse{
			Code:   http.StatusBadRequest,
			Status: "Email already registered",
			Data:   helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else if checkUsername {
		responses := helper.DefaultResponse{
			Code:   http.StatusBadRequest,
			Status: "Username already registered",
			Data:   helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		newPassword, err := controller.HashPassword(userCreateRequest.Password)
		helper.PanicIfError(err)

		userCreateRequest.Password = newPassword

		userCreateResponse := controller.AuthService.SignUp(c, &userCreateRequest)
		if userCreateResponse.Error {
			responses := helper.DefaultResponse{
				Code:   http.StatusBadRequest,
				Status: userCreateResponse.Message,
				Data:   helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(userCreateResponse.Id), 10), userCreateResponse.Email)
			responses := helper.DefaultLoginResponse{
				Code:   http.StatusOK,
				Status: "New user has been added",
				Data:   userCreateResponse,
				Token:  token,
			}
			c.JSON(http.StatusOK, responses)
		}
	}
}

// CheckPasswordHash implements AuthController
func (*AuthControllerImpl) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword implements AuthController
func (*AuthControllerImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckEmail implements AuthController
func (controller *AuthControllerImpl) CheckEmail(ctx *gin.Context, email string) bool {
	check := controller.AuthService.CheckEmail(ctx, email)
	return check
}

// CheckUsername implements AuthController
func (controller *AuthControllerImpl) CheckUsername(ctx *gin.Context, username string) bool {
	check := controller.AuthService.CheckUsername(ctx, username)
	return check
}
