package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
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

// SigninWithEmail implements AuthController
func (controller *AuthControllerImpl) SigninWithEmail(c *gin.Context) {
	userLoginRequest := dto.UserLoginEmailDTO{}
	helper.ReadFromRequestBody(c.Request, &userLoginRequest)
	userResponse := controller.AuthService.FindUserByEmail(&userLoginRequest)
	if userResponse.Username == "" {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Email is not register",
			Status:  false,
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		checkPassword := controller.CheckPasswordHash(userLoginRequest.Password, userResponse.Password)
		cvtUuid, _ := strconv.Atoi(userResponse.Uuid)
		if checkPassword {
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(cvtUuid), 10), userResponse.Email)
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
		cvtUuid, _ := strconv.Atoi(userResponse.Uuid)
		if checkPassword {
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(cvtUuid), 10), userResponse.Email)
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
	checkCompany := controller.CheckCompany(userCreateRequest.CompanySecretKey)
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
	} else if !checkCompany {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Company key is wrong!",
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
			cvtUuid, _ := strconv.Atoi(userCreateResponse.Uuid)
			token := controller.JWTService.GenerateToken(strconv.FormatUint(uint64(cvtUuid), 10), userCreateResponse.Email)
			responses := helper.DefaultLoginResponse{
				Code:    http.StatusCreated,
				Status:  true,
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

// CheckCompany implements AuthController
func (controller *AuthControllerImpl) CheckCompany(key string) bool {
	check := controller.AuthService.CheckCompany(key)
	return check
}
