package utils

import (
	"bufio"
	"fmt"
	"go/ast"
	"os"
	"strings"
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

func InitialToUpper(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func InitialToLower(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func SnakeToCamel(str string) string {
	strSplit := strings.Split(str, "_")

	resStr := ""

	for _, s := range strSplit {
		resStr = resStr + strings.Title(s)
	}
	return resStr
}

func CamelToSnake(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	sLen := len(s)
	var snake string
	for i, current := range s {
		if i > 0 && i+1 < sLen {
			if current >= 'A' && current <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(current)
	}

	snake = strings.ToLower(snake)
	return snake
}

type moduleInspector struct {
	moduleName string
	moduleSpec *ast.TypeSpec
}

func findModule(moduleName string) error {

	fp, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	line, err := reader.ReadString('\n')
	fmt.Println(line)

	return nil

}
