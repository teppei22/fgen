package usecase

import (
	"github.com/teppei22/fji-codegen/sample_layered/domain/model"
	"github.com/teppei22/fji-codegen/sample_layered/domain/repository"
)

type TaskUseCase interface {
	FindTask(id string) ([]*model.Task, error)
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: r,
	}
}

func (u taskUseCase) FindTask(id string) (task []*model.Task, err error) {
	task, err = u.taskRepository.FindTask(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
