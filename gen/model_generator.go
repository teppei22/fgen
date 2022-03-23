package fgen

import (
	"fmt"
	"strings"

	jen "github.com/dave/jennifer/jen"
)

type ModelGenerator interface {
	generate()
}

type Model struct {
	Generator ModelGenerator
}

func NewModel() *Model {
	return &Model{}
}

func (g *Model) generate() {

	f := jen.NewFile("model")

	f.Type().Id("Task").Struct(
		jen.Id("ID").Id("int").Tag(map[string]string{"json": CamelToSnake("ID")}),
		jen.Id("Title").Id("string").Tag(map[string]string{"json": CamelToSnake("Title")}),
		jen.Id("Description").Id("string").Tag(map[string]string{"json": CamelToSnake("Description")}),
		jen.Id("CreatedAt").Id("time.Time").Tag(map[string]string{"json": CamelToSnake("CreatedAt")}),
		jen.Id("UpdatedAt").Id("time.Time").Tag(map[string]string{"json": CamelToSnake("UpdatedAt")}),
	)

	if err := f.Save(`./sample_model.go`); err != nil {
		fmt.Printf("Error: %#v\n", fmt.Errorf("failed to save file"))
	}

	// fmt.Printf("%#v", f)
}

func CamelToSnake(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	len := len(s)
	var snake string
	for i, curr := range s {
		if i > 0 && i+1 < len {
			if curr >= 'A' && curr <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(curr)
	}

	snake = strings.ToLower(snake)
	return snake
}
