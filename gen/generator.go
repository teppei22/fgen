package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/teppei22/fji-codegen/gen/utils"
)

type Generator interface {
	Init() error
	GenerateFileAll() error
	OutputFile(tmpPath string, dir string, name string, data interface{}) error
	MakeDirInit() error
}

type AutoGen struct {
	Config Config
}

func NewAutoGen(conf Config) Generator {
	return &AutoGen{Config: conf}

}

func (g *AutoGen) Init() error {

	outputPath := g.Config.OutputPath

	initInfo := []TempInfo{{
		Name:         "main",
		TemplatePath: "./templates/init/main.gtpl",
		OutputDir:    outputPath,
	}, {
		Name:         "database",
		TemplatePath: "./templates/init/database.gtpl",
		OutputDir:    filepath.Join(outputPath, "infra"),
	}, {
		Name:         "router",
		TemplatePath: "./templates/init/router.gtpl",
		OutputDir:    filepath.Join(outputPath, "router"),
	}}

	if err := g.MakeDirInit(); err != nil {
		return fmt.Errorf("make dir init error: %w", err)
	}

	for _, iI := range initInfo {
		if err := g.OutputFile(iI.TemplatePath, iI.OutputDir, iI.Name, nil); err != nil {
			return fmt.Errorf("init output file error: %w", err)
		}
		utils.LogCreated(filepath.Join(iI.OutputDir, iI.Name+".go"))
	}
	return nil

}

func (g *AutoGen) GenerateFileAll() error {
	outputPath := g.Config.OutputPath
	model := g.Config.Model

	// NOTE: dir exists judge

	tempInfo := []TempInfo{
		{
			Name:         "handler",
			TemplatePath: "./templates/layer/handler.gtpl",
			OutputDir:    filepath.Join(outputPath, "handler"),
		},
		{
			Name:         "usecase",
			TemplatePath: "./templates/layer/usecase.gtpl",
			OutputDir:    filepath.Join(outputPath, "usecase"),
		},
		{
			Name:         "persistence",
			TemplatePath: "./templates/layer/persistence.gtpl",
			OutputDir:    filepath.Join(outputPath, "infra", "persistence"),
		},
		{
			Name:         "repository",
			TemplatePath: "./templates/layer/repository.gtpl",
			OutputDir:    filepath.Join(outputPath, "domain", "repository"),
		},
	}

	data := &OutputImplData{
		Handler: ImplFileData{
			PkgName:       "handler",
			StructName:    utils.InitialToLower(model) + "Handler",
			InterfaceName: model + "Handler",
			ReceiverChar:  "h",
		},
		UseCase: ImplFileData{
			PkgName:       "usecase",
			StructName:    utils.InitialToLower(model) + "UseCase",
			InterfaceName: model + "UseCase",
			ReceiverChar:  "u",
		},
		Persistence: ImplFileData{
			PkgName:       "persistence",
			StructName:    utils.InitialToLower(model) + "Persistence",
			InterfaceName: model + "Persistence",
			ReceiverChar:  "p",
		},
		Repository: ImplFileData{
			PkgName:       "repository",
			StructName:    utils.InitialToLower(model) + "Repository",
			InterfaceName: model + "Repository",
			ReceiverChar:  "r",
		},
		Model: ModelInfo{
			Name:   g.Config.Model,
			Fields: []FieldInfo{},
		},
	}

	for _, tI := range tempInfo {
		if err := g.OutputFile(tI.TemplatePath, tI.OutputDir, tI.Name, data); err != nil {
			return fmt.Errorf("output file error: %w", err)
			// panic(err)
		}
		utils.LogCreated(filepath.Join(tI.OutputDir, g.Config.Model+".go"))
	}
	return nil

}

func (g *AutoGen) OutputFile(tmpPath string, dir string, name string, data interface{}) error {

	t := template.Must(template.ParseFiles(tmpPath))

	of, err := os.Create(filepath.Join(dir, name+".go"))
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer of.Close()

	if err := t.Execute(of, data); err != nil {
		return fmt.Errorf("error executing file: %w", err)
	}
	return nil
}

func ListFiles(root string) ([]string, error) {

	fileList := []string{}

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {

				return filepath.SkipDir
			}

			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			fileList = append(fileList, rel)

			return nil
		})
	if err != nil {
		return nil, err
	}

	return fileList, nil

}

func (g *AutoGen) MakeDirInit() error {

	outputPath := g.Config.OutputPath

	dirPath := []string{
		filepath.Join(outputPath, "router"),
		filepath.Join(outputPath, "handler"),
		filepath.Join(outputPath, "usecase"),
		filepath.Join(outputPath, "domain/model"),
		filepath.Join(outputPath, "domain/repository"),
		filepath.Join(outputPath, "infra/persistence"),
	}

	for _, path := range dirPath {

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		utils.LogCreated(path)

	}

	return nil
}
