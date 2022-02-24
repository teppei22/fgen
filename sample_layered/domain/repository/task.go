package repository

import "github.com/teppei22/fji-codegen/sample_layered/domain/model"

type TaskRepository interface {
	FindTask(id string) ([]*model.Task, error)
}
