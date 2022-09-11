package controller

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"errors"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) InsertCategory(c *gin.Context) {
	categoryCreateRequest := dto.CategoryCreateDTO{}
	err := c.ShouldBind(&categoryCreateRequest)

	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "please check your image file",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		errBind := c.ShouldBindUri(&categoryCreateRequest)
		if errBind != nil {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: errBind.Error(),
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			var formatFile string
			if strings.Contains(categoryCreateRequest.Image.Filename, "jpg") {
				formatFile = ".jpg"
			} else if strings.Contains(categoryCreateRequest.Image.Filename, "jpeg") {
				formatFile = ".jpeg"
			} else if strings.Contains(categoryCreateRequest.Image.Filename, "png") {
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
				checkPath := "assets/images/categories"
				if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
					err := os.Mkdir(checkPath, os.ModePerm)
					if err != nil {
						log.Println(err)
					}
				}

				path := "assets/images/categories/" + strings.ToLower(strings.ReplaceAll(categoryCreateRequest.Name, " ", "_")) + formatFile
				errUpload := c.SaveUploadedFile(categoryCreateRequest.Image, path)
				if errUpload != nil {
					responses := helper.DefaultResponse{
						Code:    http.StatusBadRequest,
						Message: errUpload.Error(),
						Data:    helper.ObjectKosongResponse{},
						Status:  false,
					}
					c.JSON(http.StatusBadRequest, responses)
				} else {
					categoryCreateRequest.Path = "/" + path
					categoryResponse := controller.CategoryService.InsertCategory(&categoryCreateRequest)
					if categoryResponse.Error {
						responses := helper.DefaultResponse{
							Code:    http.StatusBadRequest,
							Message: categoryResponse.Message,
							Data:    helper.ObjectKosongResponse{},
							Status:  false,
						}
						c.JSON(http.StatusBadRequest, responses)
					} else {
						responses := helper.DefaultResponse{
							Code:    http.StatusOK,
							Message: "New category has been added",
							Data:    categoryResponse,
							Status:  true,
						}
						c.JSON(http.StatusOK, responses)
					}
				}
			}
		}
	}
}

func (controller *CategoryControllerImpl) FindAllCategory(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	categoryResponse, total := controller.CategoryService.FindAllCategory(c)
	var lastpage = int(math.Ceil(float64(total) / float64(4)))
	responses := helper.DefaultPaginationResponse{
		Code:     http.StatusOK,
		Message:  "Data category has been listed",
		Data:     categoryResponse,
		Status:   true,
		Page:     page,
		Total:    total,
		LastPage: lastpage,
	}
	c.JSON(http.StatusOK, responses)
}

func (controller *CategoryControllerImpl) FindByIdCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "terjadi kesalahan... silahkan coba lagi",
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		categoryResponse := controller.CategoryService.FindByIdCategory(id)
		if categoryResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: categoryResponse.Message,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else if categoryResponse.Name == "" && categoryResponse.Created == "" && categoryResponse.Updated == "" {
			responses := helper.DefaultResponse{
				Code:    http.StatusNotFound,
				Status:  false,
				Message: "category id is not found",
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusNotFound, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Status:  true,
				Message: "category id is found",
				Data:    categoryResponse,
			}
			c.JSON(http.StatusOK, responses)
		}
	}
}

func (controller *CategoryControllerImpl) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Parameter id is not found",
			Status:  false,
			Data:    helper.ObjectKosongResponse{},
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		categoryResponse := controller.CategoryService.FindByIdCategory(id)
		if categoryResponse.Error {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: categoryResponse.Message,
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else if categoryResponse.Name == "" {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: "Category id is not found",
				Status:  false,
				Data:    helper.ObjectKosongResponse{},
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			deleteResponse := controller.CategoryService.DeleteCategory(categoryResponse)

			if deleteResponse.Error {
				responses := helper.DefaultResponse{
					Code:    http.StatusBadRequest,
					Message: deleteResponse.Message,
					Data:    helper.ObjectKosongResponse{},
					Status:  false,
				}
				c.JSON(http.StatusBadRequest, responses)
			} else {
				os.Remove("." + deleteResponse.Image)
				responses := helper.DefaultResponse{
					Code:    http.StatusOK,
					Message: "Category has been deleted",
					Data:    helper.ObjectKosongResponse{},
					Status:  true,
				}
				c.JSON(http.StatusOK, responses)
			}
		}
	}
}

func (controller *CategoryControllerImpl) UpdateCategory(c *gin.Context) {
	categoryUpdateRequest := dto.CategoryUpdateDTO{}
	helper.ReadFromRequestBody(c.Request, &categoryUpdateRequest)

	categoryResponse := controller.CategoryService.UpdateCategory(&categoryUpdateRequest)
	if categoryResponse.Error {
		responses := helper.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: categoryResponse.Message,
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		/// -1 hanya penanda error
		if categoryResponse.Id == -1 {
			responses := helper.DefaultResponse{
				Code:    http.StatusBadRequest,
				Message: categoryResponse.Message,
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusBadRequest, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Message: "Category has been updated",
				Data:    categoryResponse,
				Status:  true,
			}
			c.JSON(http.StatusOK, responses)
		}
	}
}
