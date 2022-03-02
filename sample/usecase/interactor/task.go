package interactor

import (
	"errors"

	"github.com/teppei22/fji-codegen/sample/domain/entity"
)

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) CreateTask(t entity.Task) (task entity.Task, err error) {
	task, err = interactor.TaskRepository.Store(t)
	if err != nil {
		return type struct{}, errors.New("failed to create task")
	}
	return task, nil
}

// func (interactor *TaskInteractor) getTaskById(id int) (user domain.User, err error) {
// 	user, err = interactor.UserRepository.FindById(id)
// 	return
// }

// func (interactor *TaskInteractor) s() (users domain.Users, err error) {
// 	users, err = interactor.UserRepository.FindAll()
// 	return
// }

// func (interactor *TaskInteractor) Update(u domain.User) (user domain.User, err error) {
// 	user, err = interactor.UserRepository.Update(u)
// 	return
// }

// func (interactor *TaskInteractor) DeleteById(u domain.User) (err error) {
// 	err = interactor.UserRepository.DeleteById(u)
// 	return
// }
