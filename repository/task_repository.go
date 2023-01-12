package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	InsertTask(db *gorm.DB, ent *entity.Task) *entity.Task
	UpdateTask(db *gorm.DB, ent *entity.Task) *entity.Task
	GetTaskByUserId(db *gorm.DB, id int) []*entity.Task
	CheckTask(db *gorm.DB, id int) bool
}
