package repository

import (
	"github.com/teppei22/fji-codegen/sample/domain/entity"
	"github.com/teppei22/fji-codegen/sample/interface/database"
)

type TaskRepository struct {
	SqlHandler database.SqlHandler
}

func (repo *TaskRepository) Store(u entity.Task) (task entity.Task, err error) {
	if err = repo.Create(&task).Error; err != nil {
		return
	}
	user = u
	return
}
