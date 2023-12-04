package main

import (
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filepath := "./grammar/full_grammar.fl"

	input, _ := antlr.NewFileStream(filepath)
	lexer := parser.NewFlexarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parserInstace := parser.NewFlexarParser(tokens)
	parserInstace.BuildParseTrees = true
	parserInstace.Program()

}
