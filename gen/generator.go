package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Init(config *Config) error {

	outputPath := config.OutputPath

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

	if err := MakeDirInit(config); err != nil {
		return fmt.Errorf("make dir init error: %w", err)
	}

	for _, iI := range initInfo {
		if err := OutputFile(iI.TemplatePath, iI.OutputDir, iI.Name, nil); err != nil {
			return fmt.Errorf("init output file error: %w", err)
		}
	}
	return nil

}

func FileGenerate(config *Config) error {
	outputPath := config.OutputPath

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
			StructName:    "handler",
			InterfaceName: "Handler",
			ReceiverChar:  "h",
		},
		UseCase: ImplFileData{
			PkgName:       "usecase",
			StructName:    "usecase",
			InterfaceName: "Usecase",
			ReceiverChar:  "u",
		},
		Persistence: ImplFileData{
			PkgName:       "persistence",
			StructName:    "persistence",
			InterfaceName: "Persistence",
			ReceiverChar:  "p",
		},
		Repository: ImplFileData{
			PkgName:       "repository",
			StructName:    "repository",
			InterfaceName: "Repository",
			ReceiverChar:  "r",
		},
		Model: ModelInfo{
			Name:   config.Model,
			Fields: []FieldInfo{},
		},
	}

	for _, tI := range tempInfo {
		if err := OutputFile(tI.TemplatePath, tI.OutputDir, tI.Name, data); err != nil {
			return fmt.Errorf("output file error: %w", err)
			// panic(err)
		}
	}
	return nil

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

func MakeDirInit(config *Config) error {

	outputPath := config.OutputPath

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
