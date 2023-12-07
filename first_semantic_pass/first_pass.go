package first_semantic_pass

import (
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/Brennon-Oliveira/flexar/program_data"
	"github.com/Brennon-Oliveira/flexar/utils"
	"github.com/antlr4-go/antlr/v4"
)

type FirstPassVisitor struct {
	antlr.ParseTreeVisitor
}

func Visit(ctx antlr.ParseTree) {
	visitor := &FirstPassVisitor{}
	visitor.VisitProgram(ctx.(*parser.ProgramContext))
}

func (v *FirstPassVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.VisitNamespace(ctx.Namespace().(*parser.NamespaceContext))
	for _, programRule := range ctx.AllProgram_rule() {
		v.VisitProgramRule(programRule.(*parser.Program_ruleContext))
	}
	return nil
}

func (v *FirstPassVisitor) VisitProgramRule(ctx *parser.Program_ruleContext) interface{} {
	if ctx.Func_() != nil {
		v.VisitFunc(ctx.Func_().(*parser.FuncContext), ctx.EXPORT() != nil)
	}
	return nil
}

func (v *FirstPassVisitor) VisitNamespace(ctx *parser.NamespaceContext) interface{} {
	namespace := ctx.Namespace_name().GetText()
	program_data.AddFileToNamespace(namespace, utils.GetCurrentFile())
	utils.SetCurrentNamespace(namespace)

	return nil
}

func (v *FirstPassVisitor) VisitClass(ctx *parser.ClassContext) interface{} {

	className := ctx.NAME().GetText()
	program_data.AddClass(className, utils.GetCurrentNamespace(), utils.GetCurrentFile())

	return nil
}
