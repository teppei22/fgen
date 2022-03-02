package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/teppei22/fji-codegen/sample_layered/domain/model"
	"github.com/teppei22/fji-codegen/sample_layered/domain/repository"
)

type taskPersistence struct {
	Conn *gorm.DB
}

func NewTaskPersistence(conn *gorm.DB) repository.TaskRepository {
	return &taskPersistence{Conn: conn}
}

func (p *taskPersistence) FindByID(id string) (*model.Task, error) {

	var task *model.Task

	if err := p.Conn.Where("id = ?", id).Find(&task).Error; err != nil {
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
