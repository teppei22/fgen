package persistence

import (
	"sample_layered/domain/model"
	"sample_layered/domain/repository"

	"github.com/jinzhu/gorm"
)

type taskPersistence struct {
	Conn *gorm.DB
}

func NewTask(conn *gorm.DB) repository.TaskRepository {
	return &taskPersistence{Conn: conn}
}

func (p *taskPersistence) FindByID(id int) (*model.Task, error) {

	task := &model.Task{ID: id}

	if err := p.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (p *taskPersistence) Create(task *model.Task) (*model.Task, error) {

	if err := p.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (p *taskPersistence) Update(task *model.Task) (*model.Task, error) {

	if err := p.Conn.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskの削除
func (p *taskPersistence) Delete(task *model.Task) error {

	if err := p.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
