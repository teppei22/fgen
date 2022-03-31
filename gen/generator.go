package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Init(config *Config) {

	outputPath := config.outputPath

	initInfo := []TempInfo{{
		Name:         "main",
		TemplatePath: "github.com/teppei22/fgen/gen/temp/init/main.gtpl",
		OutputDir:    outputPath,
	}, {
		Name:         "database",
		TemplatePath: "github.com/teppei22/fgen/gen/temp/init/database.gtpl",
		OutputDir:    filepath.Join(outputPath, "infra"),
	}, {
		Name:         "router",
		TemplatePath: "github.com/teppei22/fgen/gen/temp/init/router.gtpl",
		OutputDir:    filepath.Join(outputPath, "router"),
	}}

	if err := MakeDirInit(); err != nil {
		panic(err)
	}

	for _, iI := range initInfo {
		if err := OutputFile(iI.TemplatePath, iI.OutputDir, iI.Name, nil); err != nil {
			panic(err)
		}
	}

}

func FileGenerate(config *Config) {
	outputPath := config.outputPath

	tempInfo := []TempInfo{
		{
			Name:         "handler",
			TemplatePath: "github.com/teppei22/fgen/gen/temp/layer/handler.gtpl",
			OutputDir:    outputPath,
		},
		{
			Name:         "usecase",
			TemplatePath: "github.com/teppei22/fgen/gen/temp/layer/usecase.gtpl",
			OutputDir:    outputPath,
		},
		{
			Name:         "persistence",
			TemplatePath: "github.com/teppei22/fgen/gen/temp/layer/persistence.gtpl",
			OutputDir:    filepath.Join(outputPath, "infra", "persistence"),
		},
		{
			Name:         "repository",
			TemplatePath: "github.com/teppei22/fgen/gen/temp/layer/repository.gtpl",
			OutputDir:    filepath.Join(outputPath, "domain", "repository"),
		},
	}

	data := &OutputImplData{
		Handler: ImplFileData{
			PkgName: "handler",

			Name:         "handler",
			TypeName:     "Handler",
			ReceiverChar: "h",
		},
		Usecase: ImplFileData{
			PkgName:      "usecase",
			Name:         "usecase",
			TypeName:     "Usecase",
			ReceiverChar: "u",
		},
		Persistence: ImplFileData{
			PkgName:      "persistence",
			Name:         "persistence",
			TypeName:     "Persistence",
			ReceiverChar: "p",
		},
		Repository: ImplFileData{
			PkgName:      "repository",
			Name:         "repository",
			TypeName:     "Repository",
			ReceiverChar: "r",
		},
	}

	for _, tI := range tempInfo {
		if err := OutputFile(tI.TemplatePath, tI.OutputDir, tI.Name, data); err != nil {
			panic(err)
		}
	}

}

func OutputFile(tmpPath string, dir string, name string, data interface{}) error {
	t := template.Must(template.ParseFiles(tmpPath))

	of, err := os.Create(filepath.Join(dir, name+".go"))
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer of.Close()

	if err := t.Execute(of, data); err != nil {
		return fmt.Errorf("error executing file: %w", err)
	}
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

func MakeDirInit() error {

	dirPath := []string{
		"router",
		"handler",
		"usecase",
		"domain/model",
		"domain/repository",
		"infra/persistence",
	}

	for _, path := range dirPath {

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}

	}

	return nil
}

func MakeDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func NewInitInfo() {

}
