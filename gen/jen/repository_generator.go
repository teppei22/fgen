package gen

import (
	"fmt"

	jen "github.com/dave/jennifer/jen"
)

type RepositoryGenerator interface {
	Generate()
}

type Repository struct {
	Generator ModelGenerator
}

func NewRepository() *Repository {
	return &Repository{}
}

func (g *Repository) Generate() {

	f := jen.NewFile("repository")

	f.Type().Id("TaskRepository").Interface(

		jen.Id("Create").Params(jen.Id("task").Id("*model.Task")).Params(jen.Id("*model.Task"), jen.Id("error")),
		jen.Id("FindByID").Params(jen.Id("id").Id("int")).Params(jen.Id("*model.Task"), jen.Id("error")),
		jen.Id("Update").Params(jen.Id("task").Id("*model.Task")).Params(jen.Id("*model.Task"), jen.Id("error")),
		jen.Id("Delete").Params(jen.Id("task").Id("*model.Task")).Params(jen.Id("error")),
	)

	if err := f.Save(`./output/sample_repository.go`); err != nil {
		fmt.Printf("Error: %#v\n", fmt.Errorf("failed to save file"))
	}

	// fmt.Printf("%#v", f)
}
