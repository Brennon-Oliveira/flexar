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
	if ctx.Class() != nil {
		v.VisitClass(ctx.Class().(*parser.ClassContext))
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

	utils.SetCurrentClass(className)

	for _, classRule := range ctx.Class_body().AllClass_body_rule() {
		if classRule.(*parser.Class_body_ruleContext).Class_attribute() != nil {
			v.VisitAttribute(classRule.(*parser.Class_body_ruleContext).Class_attribute().(*parser.Class_attributeContext))
		}
	}

	return nil
}

func (v *FirstPassVisitor) VisitAttribute(ctx *parser.Class_attributeContext) interface{} {
	attributeName, attributeType := v.VisitVariableDeclaration(ctx.Variable_declaration().(*parser.Variable_declarationContext))

	program_data.AddAttributeToClass(utils.GetCurrentNamespace(), utils.GetCurrentClass(), program_data.Attribute{Name: attributeName, Type: attributeType})
	return nil
}

func (v *FirstPassVisitor) VisitVariableDeclaration(ctx *parser.Variable_declarationContext) (variableName string, variableType string) {
	variableName = ctx.NAME().GetText()
	variableType = ctx.Type_().GetText()

	return
}
