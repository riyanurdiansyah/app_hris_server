package controller

import (
	"app-travel-server/data/dto"
	"app-travel-server/helper"
	"app-travel-server/service"
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

// SigninWithUsername implements AuthController
func (controller *AuthControllerImpl) SigninWithUsername(c *gin.Context) {
	userLoginRequest := dto.UserLoginUsernameDTO{}
	helper.ReadFromRequestBody(c.Request, &userLoginRequest)
	userResponse := controller.AuthService.FindUserByUsername(&userLoginRequest)
	if userResponse.Username == "" {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Username is not register",
			Status:  false,
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		checkPassword := controller.CheckPasswordHash(userLoginRequest.Password, userResponse.Password)
		if checkPassword {
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(userResponse.Id), 10), userResponse.Email)
			responses := helper.DefaultLoginResponse{
				Code:    http.StatusOK,
				Message: "Login is successfull",
				Status:  true,
				Data:    userResponse,
				Token:   token,
			}
			c.JSON(http.StatusOK, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: "Password is wrong",
				Status:  false,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		}
	}
}

// SignUp implements AuthController
func (controller *AuthControllerImpl) SignUp(c *gin.Context) {
	userCreateRequest := dto.UserCreateDTO{}
	helper.ReadFromRequestBody(c.Request, &userCreateRequest)

	checkEmail := controller.CheckEmail(userCreateRequest.Email)
	checkUsername := controller.CheckUsername(userCreateRequest.Username)
	if checkEmail {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already registered",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else if checkUsername {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Username already registered",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		newPassword, err := controller.HashPassword(userCreateRequest.Password)
		helper.PanicIfError(err)

		userCreateRequest.Password = newPassword

		userCreateResponse := controller.AuthService.SignUp(&userCreateRequest)
		if userCreateResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: userCreateResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(userCreateResponse.Id), 10), userCreateResponse.Email)
			responses := helper.DefaultLoginResponse{
				Code:    http.StatusOK,
				Message: "New user has been added",
				Data:    userCreateResponse,
				Token:   token,
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
func (controller *AuthControllerImpl) CheckEmail(email string) bool {
	check := controller.AuthService.CheckEmail(email)
	return check
}

// CheckUsername implements AuthController
func (controller *AuthControllerImpl) CheckUsername(username string) bool {
	check := controller.AuthService.CheckUsername(username)
	return check
}
