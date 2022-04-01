package usecase

import (
	"github.com/teppei22/fgen/sample_layered/domain/model"
	"github.com/teppei22/fgen/sample_layered/domain/repository"
)

type TaskUseCase interface {
	FindByID(id int64) (task *model.Task, err error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id int64) error
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: r,
	}
}

func (u *taskUseCase) FindByID(id int64) (task *model.Task, err error) {
	task, err = u.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *taskUsecase) Create(task *model.Task) (*model.Task, error) {
	createdTask, err := u.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (u *taskUsecase) Update(task *model.Task) (*model.Task, error) {

	updatedTask, err := u.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskUsecase) Delete(id int64) error {

	err := u.taskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
