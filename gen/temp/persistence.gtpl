package persistence

import (
	"sample_layered/domain/model"
	"sample_layered/domain/repository"

	"github.com/jinzhu/gorm"
)

type {{ .Persistence.StructName }} struct {
	Conn *gorm.DB
}

func New{{ .Model.Name }} (conn *gorm.DB) repository.{{ .Repository.InterfaceName }} {
	return &{{ .Persistence.StructName }}{Conn: conn}
}

func (p *{{ .Persistence.StructName }}) FindByID(id int) (*model.{{ .Model.Name }}, error) {

	{{ .Model.Name }} := &model.{{ .Model.Name }}{ID: id}

	if err := p.Conn.First(&{{ .Model.Name }}).Error; err != nil {
		return nil, err
	}

	return {{ .Model.Name }}, nil
}

func (p *{{ .Persistence.StructName }}) Create({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error) {

	if err := p.Conn.Create(&{{ .Model.Name }}).Error; err != nil {
		return nil, err
	}

	return {{ .Model.Name }}, nil
}

func (p *{{ .Persistence.StructName }}) Update({{ .Model.Name }} *model.{{ .Model.Name }}) (*model.{{ .Model.Name }}, error) {

	if err := p.Conn.Model(&{{ .Model.Name }}).Update(&{{ .Model.Name }}).Error; err != nil {
		return nil, err
	}

	return {{ .Model.Name }}, nil
}

func (p *{{ .Persistence.StructName }}) Delete({{ .Model.Name }} *model.{{ .Model.Name }}) error {

	if err := p.Conn.Delete(&{{ .Model.Name }}).Error; err != nil {
		return err
	}

	return nil
}
