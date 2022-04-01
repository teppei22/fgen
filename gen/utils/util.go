package utils

import (
	"fmt"
	"os"
)

func MakeDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func LogCreated(outputPath string) {
	fmt.Println("Created:", outputPath)
}
