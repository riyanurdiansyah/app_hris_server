package service

import (
	"app-hris-server/data/dto"
)

type TaskService interface {
	InsertTask(request *dto.TaskCreateDTO) *dto.TaskResponseDTO
	UpdateTask(request *dto.TaskUpdateDTO) *dto.TaskResponseDTO
	GetTaskByUserId(userId int) []*dto.TaskResponseDTO
	CheckTask(id int) bool
}
