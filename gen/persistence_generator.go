package fgen

import (
	"log"

	jen "github.com/dave/jennifer/jen"
)

type PersistenceGenerator interface {
	Generate()
}

type Persistence struct {
	Generator PersistenceGenerator
}

func NewPersistence() *Persistence {
	return &Persistenceπ{}
}

func (g *Persistence) Generate() {

	f := jen.NewFile("persistence")

	f.Type().Id("taskPersistence").Struct(
		jen.Id("Conn").Id("*gorm.DB"),
	)

	// if err := f.Save(`./sample_model.go`); err != nil {
	// 	fmt.Printf("Error: %#v\n", fmt.Errorf("failed to save file"))
	// }

	f.Func().Id("NewTask").Params(jen.Id("conn").Id("*gorm.DB")).Params(jen.Id("repository").Op(".").Id("TaskRepository")).Block(
		jen.Return(jen.Id("&taskPersistence").Block(jen.Id("Conn:").Id("conn").Op(","))),
	)

	f.Func().Params(jen.Id("p").Id("*taskPersistence")).Id("FindByID").Params(jen.Id("id").Id("int")).Params(jen.Id("*model.Task"), jen.Id("error")).Block(

		jen.Id("task").Op(":=").Op("&").Id("model").Id("Task").Block(jen.Id("ID").Op(":").Id("id"))
	)

	f.Line()
	f.Func().Params(jen.Id("p").Id("*taskPersistence")).Id("Create").Params(jen.Id("*model.Task")).Params(jen.Id("*model.Task"), jen.Id("error")).Block(
	// jen.Return(jen.Nil),
	)

	f.Line()
	f.Func().Params(jen.Id("p").Id("*taskPersistence")).Id("Update").Params(jen.Id("*model.Task")).Params(jen.Id("*model.Task"), jen.Id("error")).Block(
	// jen.Return(jen.Nil),
	)

	f.Line()
	f.Func().Params(jen.Id("p").Id("*taskPersistence")).Id("Delete").Params(jen.Id("*model.Task")).Params(jen.Id("*model.Task"), jen.Id("error")).Block(
	// jen.Return(jen.Nil),
	)

	if err := f.Save(`./output/sample_persistence.go`); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%#v", f)
}
