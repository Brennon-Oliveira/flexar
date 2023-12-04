package compilation

import (
	"fmt"
	"github.com/Brennon-Oliveira/flexar/first_semantic_pass"
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"path/filepath"
)

func Compile(filePath string) {
	fmt.Printf("Compiling %s\n", filePath)
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewFlexarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parserInstace := parser.NewFlexarParser(tokens)
	parserInstace.BuildParseTrees = true
	tree := parserInstace.Program()
	first_semantic_pass.Visit(tree)
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
