package gen

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func Init(config Config) {

	outputPath := config.outputPath

	initInfo := []InitInfo{{
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

func OutputFile(tmpPath string, outputDir string, outputName string, data interface{}) error {
	t := template.Must(template.ParseFiles(tmpPath))

	of, err := os.Create(filepath.Join(outputDir, outputName+".go"))
	if err != nil {
		return fmt.Errorf("creating file error: %w", err)
	}
	defer of.Close()

	if err := t.Execute(of, data); err != nil {
		return fmt.Errorf("execute file error: %w", err)
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
