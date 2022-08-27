package controller

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type PromoControllerImpl struct {
	PromoService service.PromoService
}

func NewPromoController(promoService service.PromoService) PromoController {
	return &PromoControllerImpl{
		PromoService: promoService,
	}
}

// GetAllPromo implements PromoController
func (controller *PromoControllerImpl) GetAllPromo(c *gin.Context) {
	categoryResponse := controller.PromoService.GetAllPromo()
	responses := helper.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Data promo has been listed",
		Data:    categoryResponse,
		Status:  true,
	}
	c.JSON(http.StatusOK, responses)
}

// InsertPromo implements PromoController
func (controller *PromoControllerImpl) InsertPromo(c *gin.Context) {
	promoCreateRequest := dto.PromoCreateDTO{}
	err := c.ShouldBind(&promoCreateRequest)

	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "please check your image file",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		errBind := c.ShouldBindUri(&promoCreateRequest)
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
			if strings.Contains(promoCreateRequest.Image.Filename, "jpg") {
				formatFile = ".jpg"
			} else if strings.Contains(promoCreateRequest.Image.Filename, "jpeg") {
				formatFile = ".jpeg"
			} else if strings.Contains(promoCreateRequest.Image.Filename, "png") {
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
				checkPath := "assets/images/promos"
				if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
					err := os.Mkdir(checkPath, os.ModePerm)
					if err != nil {
						log.Println(err)
					}
				}

				path := "assets/images/promos/" + strings.ToLower(strings.ReplaceAll(promoCreateRequest.Name, " ", "_")) + formatFile
				errUpload := c.SaveUploadedFile(promoCreateRequest.Image, path)
				if errUpload != nil {
					responses := helper.DefaultResponse{
						Code:    http.StatusBadRequest,
						Message: errUpload.Error(),
						Data:    helper.ObjectKosongResponse{},
						Status:  false,
					}
					c.JSON(http.StatusBadRequest, responses)
				} else {
					promoCreateRequest.Path = "/" + path
					promoResponse := controller.PromoService.InsertPromo(&promoCreateRequest)
					if promoResponse.Error {
						responses := helper.DefaultResponse{
							Code:    http.StatusBadRequest,
							Message: promoResponse.Message,
							Data:    helper.ObjectKosongResponse{},
							Status:  false,
						}
						c.JSON(http.StatusBadRequest, responses)
					} else {
						responses := helper.DefaultResponse{
							Code:    http.StatusOK,
							Message: "New promo has been added",
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
