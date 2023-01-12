package repository

import (
	"app-hris-server/data/entity"
	"app-hris-server/helper"
	"fmt"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct{}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

// CheckTask implements TaskRepository
func (*TaskRepositoryImpl) CheckTask(db *gorm.DB, id int) bool {
	var task = entity.Task{}
	result := db.Table("user_info_personal").Select("*").Where("id = ?", task.ID).Scan(&task)
	if result.Error != nil {
		return false
	}

	if task.IdUser == "" {
		return false
	} else {
		return true
	}
}

// GetTaskByUserId implements TaskRepository
func (*TaskRepositoryImpl) GetTaskByUserId(db *gorm.DB, id int) []*entity.Task {
	var tasks = []*entity.Task{}
	result := db.Table("tasks").Joins("INNER JOIN user_info_personal user on user.id_user = tasks.task_by").Where("tasks.id_user= ?", id).Select("*").Find(&tasks)
	helper.PanicIfError(result.Error)
	fmt.Println(tasks[0].NamaBelakang)
	return tasks
}

// InsertTask implements TaskRepository
func (*TaskRepositoryImpl) InsertTask(db *gorm.DB, ent *entity.Task) *entity.Task {
	panic("unimplemented")
}

// UpdateTask implements TaskRepository
func (*TaskRepositoryImpl) UpdateTask(db *gorm.DB, ent *entity.Task) *entity.Task {
	panic("unimplemented")
}
