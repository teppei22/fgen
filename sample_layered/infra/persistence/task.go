package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/teppei22/fji-codegen/sample_layered/repository"
)

type taskPersistence struct {
	Conn *gorm.DB
}

func NewTaskPersistence(conn *gorm.DB) repository.TaskRepository {
	return &taskPersistence{Conn: conn}
}

func (t *TaskPersistence) FindTask(id string) (*model.Task, error) {
	var task model.User

	// DB接続確認
	// if err := t.Conn.Take(&user).Error; err != nil {
	// 	return nil, err
	// }

	db := t.Conn.Find(&task)
	if id != "" {
		db = db.Where("id = ?", id).Find(&task)
	}

	return task, nil
}
