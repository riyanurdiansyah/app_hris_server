package controller

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"net/http"
	"strconv"

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
	helper.ReadFromRequestBody(c.Request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.InsertCategory(c, &categoryCreateRequest)
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

func (controller *CategoryControllerImpl) FindAllCategory(c *gin.Context) {
	categoryResponse := controller.CategoryService.FindAllCategory(c)
	responses := helper.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Data category has been listed",
		Data:    categoryResponse,
		Status:  true,
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
		categoryResponse := controller.CategoryService.FindByIdCategory(c, id)
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
			Message: "terjadi kesalahan... silahkan coba lagi",
			Data:    helper.ObjectKosongResponse{},
			Status:  false,
		}
		c.JSON(http.StatusBadRequest, responses)
	} else {
		categoryResponse := controller.CategoryService.DeleteCategory(c, id)
		if categoryResponse > 0 {
			responses := helper.DefaultResponse{
				Code:    http.StatusOK,
				Message: "data category has been deleted",
				Data:    helper.ObjectKosongResponse{},
				Status:  true,
			}
			c.JSON(http.StatusOK, responses)
		} else if categoryResponse == 0 {
			responses := helper.DefaultResponse{
				Code:    http.StatusNotFound,
				Message: "category id is not found",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusNotFound, responses)
		} else {
			responses := helper.DefaultResponse{
				Code:    http.StatusInternalServerError,
				Message: "terjadi kesalahan... silahkan coba beberapa saat lagi",
				Data:    helper.ObjectKosongResponse{},
				Status:  false,
			}
			c.JSON(http.StatusInternalServerError, responses)
		}
	}
}

func (controller *CategoryControllerImpl) UpdateCategory(c *gin.Context) {
	categoryUpdateRequest := dto.CategoryUpdateDTO{}
	helper.ReadFromRequestBody(c.Request, &categoryUpdateRequest)

	categoryResponse := controller.CategoryService.UpdateCategory(c, &categoryUpdateRequest)
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
