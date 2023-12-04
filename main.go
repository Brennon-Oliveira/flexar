package main

import (
	"fmt"
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// receive file path as argument

	if len(os.Args) < 2 {
		fmt.Println("You must provide a path as argument")
		return
	}

	projectPath := os.Args[1]

	files := []string{}

	// check if path is a flexar file (.fl)

	if filepath.Ext(projectPath) != ".fl" {
		// check if path is a directory
		if info, err := os.Stat(projectPath); err == nil && info.IsDir() {
			findFiles(projectPath, &files)
		} else {
			fmt.Println("Invalid path")
			return
		}
	} else {
		// check if file exists
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		files = append(files, projectPath)
	}

	for _, file := range files {
		compile(file)
	}

}

func compile(filePath string) {
	fmt.Printf("Compiling %s\n", filePath)
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewFlexarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parserInstace := parser.NewFlexarParser(tokens)
	parserInstace.BuildParseTrees = true
	parserInstace.Program()
	fmt.Printf("Compiled %s\n", filePath)
	fmt.Println("=====================================")
}

func findFiles(path string, files *[]string) {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		itemPath := filepath.Join(path, item.Name())
		if item.IsDir() {
			findFiles(itemPath, files)
		} else {
			*files = append(*files, itemPath)
		}
	}

}
