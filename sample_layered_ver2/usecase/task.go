package usecase

import (
	"sample_layered/domain/model"
	"sample_layered/domain/repository"
)

type TaskUseCase interface {
	FindByID(id int) (task *model.Task, err error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id int) error
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTask(r repository.TaskRepository) TaskUseCase {
	// return &taskUseCase{
	// 	taskRepository: r,
	// }
	return &taskUseCase{
		taskRepository: r,
	}
	// return nil
}

func (u *taskUseCase) FindByID(id int) (*model.Task, error) {
	task, err := u.taskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskUseCase) Create(task *model.Task) (*model.Task, error) {
	createdTask, err := u.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (u *taskUseCase) Update(task *model.Task) (*model.Task, error) {

	updatedTask, err := u.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskUseCase) Delete(id int) error {

	task, err := u.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	err = u.taskRepository.Delete(task)
	if err != nil {
		return err
	}

	return nil
}
