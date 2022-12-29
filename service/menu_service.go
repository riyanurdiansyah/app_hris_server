package service

import (
	"app-hris-server/data/dto"
)

type MenuService interface {
	InsertMenu(request *dto.MenuCreateDTO) *dto.MenuResponseDTO
	UpdateMenu(request *dto.MenuUpdateDTO) *dto.MenuResponseDTO
	GetMenu() []*dto.MenuResponseDTO
	CheckMenu(id int) bool
}
