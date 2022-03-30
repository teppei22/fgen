package {{ .Repository.pkgName }}

import "sample_layered/domain/model"

type {{ .Repository.InterfaceName }} interface {
	Create(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}
