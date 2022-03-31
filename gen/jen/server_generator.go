package gen

import (
	"fmt"
	"log"
	"os"
	"strings"

	jen "github.com/dave/jennifer/jen"
)

type ServerGenerator interface {
	Execute()
}

type Server struct {
	// Generator ServerGenerator
}

func NewServer() ServerGenerator {
	return &Server{}
}

const (
	MOD_DIR = "github.com/teppei22/fgen/sample_layered"
)

func (s *Server) Execute() {

	// make router
	if _, err := os.Stat("./output/router"); err != nil {
		fmt.Println("router dir doesn't exist")
		if err := os.Mkdir("./output/router", 0777); err != nil {
			fmt.Println(err)
		}
	}

	r := jen.NewFile("router")

	r.ImportName("github.com/labstack/echo/v4", "echo")
	// module path を取得
	r.ImportName(strings.Join([]string{MOD_DIR, "handler"}, "/"), "handler")
	r.ImportName(strings.Join([]string{MOD_DIR, "usecase"}, "/"), "usecase")
	r.ImportName(strings.Join([]string{MOD_DIR, "infra", "persistence"}, "/"), "persistence")
	r.ImportName(strings.Join([]string{MOD_DIR, "infra"}, "/"), "infra")

	r.Func().Id("Init").Params().Op("*").Qual("github.com/labstack/echo/v4", "Echo").BlockFunc(func(g *jen.Group) {

		g.Id("e").Op(":=").Qual("github.com/labstack/echo/v4", "New").Call()

		g.Line().Commentf("DB connect")

		g.Id("conn").Op(":=").Qual("github.com/teppei22/fgen/sample_layered/infra", "DBConnect").Call()

		g.Line().Commentf("task handler")
		g.Id("taskHandler").Op(":=").Qual("github.com/teppei22/fgen/sample_layered/handler", "NewTaskHandler").Parens(
			jen.Qual("github.com/teppei22/fgen/sample_layered/usecase", "NewTaskUseCase").Parens(
				jen.Qual("github.com/teppei22/fgen/sample_layered/infra/persistence", "NewTaskPersistence").Parens(jen.Id("conn")),
			),
		)

		g.Line().Commentf("task router")
		g.Id("e").Dot("GET").Call(jen.Lit("task/:id"), jen.Id("taskHandler").Op(".").Id("FindByID"))
		g.Id("e").Dot("POST").Call(jen.Lit("task"), jen.Id("taskHandler").Op(".").Id("Create"))
		g.Id("e").Dot("PUT").Call(jen.Lit("task/:id"), jen.Id("taskHandler").Op(".").Id("Update"))
		g.Id("e").Dot("DELETE").Call(jen.Lit("task/:id"), jen.Id("taskHandler").Op(".").Id("Delete"))

		g.Line().Return(jen.Id("e"))
	})

	if err := r.Save(`./output/router/sample_server.go`); err != nil {
		log.Fatal(err)
	}

	// f := jen.NewFile("main")
	// f.List(jen.Id("router")).Op(":=").Qual("route", "Init").Call()

	// if err := f.Save(`./sample_server.go`); err != nil {
	// 	// fmt.Printf(fmt.Errorf("error failed to save file: %s", err))
	// 	log.Fatal(err)
	// 	// fmt.Printf("Error: %#v\n", err)
	// }

	// fmt.Printf("%#v", f)
}
