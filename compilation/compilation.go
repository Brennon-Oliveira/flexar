package compilation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Brennon-Oliveira/flexar/first_semantic_pass"
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/antlr4-go/antlr/v4"
)

func Tokenize(input *antlr.FileStream) *antlr.CommonTokenStream {
	lexer := parser.NewFlexarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	return tokens
}

func Parse(tokens *antlr.CommonTokenStream) *parser.IProgramContext {
	parserInstace := parser.NewFlexarParser(tokens)
	parserInstace.BuildParseTrees = true
	tree := parserInstace.Program()
	if parserInstace.HasError() {
		// TODO: Don't exit the process deep in the compilation logic. Return an error instead.
		os.Exit(0)
	}

	return &tree
}

func Compile(projectPath string) {
	fmt.Printf("Compiling %s\n", filePath)

	input, _ := antlr.NewFileStream(filePath)

	tokens := Tokenize(input)

	tree := Parse(tokens)

	first_semantic_pass.Visit(*tree)
	fmt.Printf("Compiled %s\n", filePath)
	fmt.Println("=====================================")
}

func FindFiles(path string, files *[]string) {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		itemPath := filepath.Join(path, item.Name())
		if item.IsDir() {
			FindFiles(itemPath, files)
		} else if filepath.Ext(itemPath) == ".fl" {
			*files = append(*files, itemPath)
		}
	}
}
