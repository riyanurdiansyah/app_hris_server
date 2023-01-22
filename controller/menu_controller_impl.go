package controller

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/service"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type MenuControllerImpl struct {
	MenuService service.MenuService
	JWTService  service.JWTService
}

func NewMenuController(menuService service.MenuService, jwtService service.JWTService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService,
		JWTService:  jwtService,
	}
}

// CheckMenu implements MenuController
func (controller *MenuControllerImpl) CheckMenu(id int) bool {
	return controller.MenuService.CheckMenu(id)
}

// InsertMenu implements MenuController
func (controller *MenuControllerImpl) InsertMenu(c *gin.Context) {
	menuCreateRequest := dto.MenuCreateDTO{}
	err := c.ShouldBind(&menuCreateRequest)

	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "please check your image file",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		errBind := c.ShouldBindUri(&menuCreateRequest)
		if errBind != nil {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: "please check your file image",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			var formatFile string
			if strings.Contains(menuCreateRequest.Image.Filename, "jpg") {
				formatFile = ".jpg"
			} else if strings.Contains(menuCreateRequest.Image.Filename, "jpeg") {
				formatFile = ".jpeg"
			} else if strings.Contains(menuCreateRequest.Image.Filename, "png") {
				formatFile = ".png"
			} else {
				formatFile = ""
			}

			if formatFile == "" {
				responses := helper.DefaultResponse{
					Code:    http.StatusBadRequest,
					Message: "format file must .jpg/.jpeg/.png",
					Data:    helper.ObjectKosongResponse{},
					Status:  false,
				}
				c.JSON(http.StatusBadRequest, responses)
			} else {

				checkPath := "assets/images/menu"
				if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
					err := os.Mkdir(checkPath, os.ModePerm)
					if err != nil {
						log.Println(err)
					}
				}

				path := checkPath + "/" + strings.ToLower(strings.ReplaceAll(menuCreateRequest.Title, " ", "_")) + formatFile
				menuCreateRequest.Path = "/" + path
				promoResponse := controller.MenuService.InsertMenu(&menuCreateRequest)
				if promoResponse.Error {
					responses := helper.DefaultResponse{
						Code:    http.StatusBadRequest,
						Message: promoResponse.Message,
						Data:    helper.ObjectKosongResponse{},
						Status:  false,
					}
					c.JSON(http.StatusBadRequest, responses)
				} else {
					c.SaveUploadedFile(menuCreateRequest.Image, path)

					// if errUpload != nil {
					// 	responses := helper.DefaultResponse{
					// 		Code:    http.StatusBadRequest,
					// 		Message: errUpload.Error(),
					// 		Data:    helper.ObjectKosongResponse{},
					// 		Status:  false,
					// 	}
					// 	c.JSON(http.StatusBadRequest, responses)
					// } else {

					// }
					responses := helper.DefaultResponse{
						Code:    http.StatusCreated,
						Message: "New menu has been added",
						Data:    promoResponse,
						Status:  true,
					}
					c.JSON(http.StatusOK, responses)
				}
			}
		}
	}
}

// UpdateMenu implements MenuController
func (controller *MenuControllerImpl) UpdateMenu(c *gin.Context) {
	menuUpdateRequest := dto.MenuUpdateDTO{}
	err := c.ShouldBind(&menuUpdateRequest)

	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "please check your image file",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		errBind := c.ShouldBindUri(&menuUpdateRequest)
		if errBind != nil {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: "please check your file image",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			var formatFile string
			if strings.Contains(menuUpdateRequest.Image.Filename, "jpg") {
				formatFile = ".jpg"
			} else if strings.Contains(menuUpdateRequest.Image.Filename, "jpeg") {
				formatFile = ".jpeg"
			} else if strings.Contains(menuUpdateRequest.Image.Filename, "png") {
				formatFile = ".png"
			} else {
				formatFile = ""
			}

			if formatFile == "" {
				responses := helper.DefaultResponse{
					Code:    http.StatusBadRequest,
					Message: "format file must .jpg/.jpeg/.png",
					Data:    helper.ObjectKosongResponse{},
					Status:  false,
				}
				c.JSON(http.StatusBadRequest, responses)
			} else {
				checkId := controller.MenuService.CheckMenu(menuUpdateRequest.Id)
				if checkId {
					checkPath := "assets/images/menu"
					oldMenu := controller.MenuService.GetMenuById(menuUpdateRequest.Id)

					if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
						err := os.Mkdir(checkPath, os.ModePerm)
						if err != nil {
							log.Println(err)
						}
					}

					path := "assets/images/menu/" + strings.ToLower(strings.ReplaceAll(menuUpdateRequest.Title, " ", "_")) + formatFile
					errUpload := c.SaveUploadedFile(menuUpdateRequest.Image, path)
					if errUpload != nil {
						responses := helper.DefaultResponse{
							Code:    http.StatusBadRequest,
							Message: errUpload.Error(),
							Data:    helper.ObjectKosongResponse{},
							Status:  false,
						}
						c.JSON(http.StatusBadRequest, responses)
					} else {
						menuUpdateRequest.Path = "/" + path
						promoResponse := controller.MenuService.UpdateMenu(&menuUpdateRequest)
						if promoResponse.Error {
							responses := helper.DefaultResponse{
								Code:    http.StatusBadRequest,
								Message: promoResponse.Message,
								Data:    helper.ObjectKosongResponse{},
								Status:  false,
							}
							c.JSON(http.StatusBadRequest, responses)
						} else {
							os.Remove(string([]rune(oldMenu.Image)[:0]) + "" + string([]rune(oldMenu.Image)[1:]))
							responses := helper.DefaultResponse{
								Code:    http.StatusOK,
								Message: "New menu has been updated",
								Data:    promoResponse,
								Status:  true,
							}
							c.JSON(http.StatusOK, responses)
						}
					}
				}
			}
		}
	}
}

// GetMenu implements MenuController
func (controller *MenuControllerImpl) GetMenu(c *gin.Context) {
	userResponse := controller.MenuService.GetMenu()
	responses := helper.DefaultResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Data has been listed",
		Data:    userResponse,
	}
	c.JSON(http.StatusOK, responses)
}

// GetMenuById implements MenuController
func (controller *MenuControllerImpl) GetMenuById(c *gin.Context) {
}
