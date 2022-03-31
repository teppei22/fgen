package {{ .UseCase.PkgName }}

import (
	"sample_layered/domain/model"
	"sample_layered/domain/repository"
)

type {{ .UseCase.InterfaceName }} interface {
	FindByID(id int) ({{ .Model.Name }} *model.{{ .Model.Name }}, err error)
	Create({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error)
	Update({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error)
	Delete(id int) error
}

type {{ .UseCase.StructName }} struct {
	{{ .Repository.StructName }} repository.{{ .Repository.InterfaceName }}
}

func New{{ .Model.Name }}(r repository.{{ .Repository.InterfaceName }}) {{ .UseCase.InterfaceName }} {ß
	return &{{ .UseCase.StructName }}{
		{{ .Repository.StructName }}: r,
	}
}

func (u *{{ .UseCase.StructName }}) FindByID(id int) (*model.{{ .Model.Name }}, error) {
	{{ .Model.Name }}, err := u.{{ .Repository.StructName }}.FindByID(id)
	if err != nil {
		return nil, err
	}

	return {{ .Model.Name }}, nil
}

func (u *{{ .UseCase.StructName }}) Create({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error) {
	created{{ .Model.Name }}, err := u.{{ .Repository.StructName }}.Create({{ .Model.Name }})
	if err != nil {
		return nil, err
	}

	return created{{ .Model.Name }}, nil
}

func (u *{{ .UseCase.StructName }}) Update({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error) {

	updated{{ .Model.Name }}, err := u.{{ .Repository.StructName }}.Update({{ .Model.Name }})
	if err != nil {
		return nil, err
	}

	return updated{{ .Model.Name }}, nil
}

func (u *{{ .UseCase.StructName }}) Delete(id int) error {

	{{ .Model.Name }}, err := u.{{ .Repository.StructName }}.FindByID(id)
	if err != nil {
		return err
	}

	err = u.{{ .Repository.StructName }}.Delete({{ .Model.Name }})
	if err != nil {
		return err
	}

	return nil
}
