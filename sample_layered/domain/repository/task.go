package repository

import "github.com/teppei22/fji-codegen/sample_layered/domain/model"

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindByID(id string) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}
