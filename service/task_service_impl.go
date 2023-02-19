package service

import (
	"app-hris-server/data/dto"
	"app-hris-server/helper"
	"app-hris-server/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewTaskService(repo repository.TaskRepository, DB *gorm.DB, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

// CheckTask implements TaskService
func (*TaskServiceImpl) CheckTask(id int) bool {
	panic("unimplemented")
}

// GetTaskByUserId implements TaskService
func (service *TaskServiceImpl) GetTaskByUserId(userId string) []*dto.TaskResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.TaskResponseDTO{}
	} else {
		tasks := service.TaskRepository.GetTaskByUserId(tx, userId)
		return dto.ToListTaskResponseDTO(tasks)
	}
}

// InsertTask implements TaskService
func (*TaskServiceImpl) InsertTask(request *dto.TaskCreateDTO) *dto.TaskResponseDTO {
	panic("unimplemented")
}

// UpdateTask implements TaskService
func (*TaskServiceImpl) UpdateTask(request *dto.TaskUpdateDTO) *dto.TaskResponseDTO {
	panic("unimplemented")
}
