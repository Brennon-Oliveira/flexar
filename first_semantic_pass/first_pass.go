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
	namespace := v.VisitNamespace(ctx.Namespace().(*parser.NamespaceContext))
	for _, programRule := range ctx.AllProgram_rule() {
		v.VisitProgramRule(programRule.(*parser.Program_ruleContext), namespace)
	}
	return nil
}

func (v *FirstPassVisitor) VisitProgramRule(ctx *parser.Program_ruleContext, namespace *program_data.Namespace) interface{} {
	if ctx.Func_() != nil {
		v.VisitFunc(ctx.Func_().(*parser.FuncContext), ctx.EXPORT() != nil, namespace)
	}
	if ctx.Class() != nil {
		v.VisitClass(ctx.Class().(*parser.ClassContext), namespace)
	}
	return nil
}

func (v *FirstPassVisitor) VisitNamespace(ctx *parser.NamespaceContext) *program_data.Namespace {
	namespace := program_data.GetNamespace(ctx.Namespace_name().GetText())

	namespace.AddFile(utils.GetCurrentFile())

	return namespace
}

func (v *FirstPassVisitor) VisitClass(ctx *parser.ClassContext, namespace *program_data.Namespace) interface{} {
	className := ctx.NAME().GetText()

	if namespace.ExistsClass(className) {
		utils.SemanticError(ctx.NAME().GetSymbol().GetLine(), "Class "+className+" already exists in namespace "+namespace.Name)
	}

	class := namespace.AddClass(className)

	for _, classAttribute := range ctx.Class_body().AllClass_attribute() {
			v.VisitAttribute(classAttribute.(*parser.Class_attributeContext), class)
	}

	return nil
}

func (v *FirstPassVisitor) VisitAttribute(ctx *parser.Class_attributeContext, class *program_data.Class) interface{} {
	attributeName, attributeType := v.VisitVariableDeclaration(ctx.Variable_declaration().(*parser.Variable_declarationContext))

	if class.ExistsAttribute(attributeName) {
		utils.SemanticError(ctx.Variable_declaration().(*parser.Variable_declarationContext).NAME().GetSymbol().GetLine(), "Attribute "+attributeName+" already exists in class "+class.Name)
	}

	attribute := program_data.Attribute{
		Name:      attributeName,
		Type:      attributeType,
		Readonly:  ctx.Class_modifier().READONLY() != nil,
		Overriden: ctx.Class_modifier().OVERRIDE() != nil,
		Abstract:  ctx.Class_modifier().ABSTRACT() != nil,
		Static:    ctx.Class_modifier().STATIC() != nil,
		Final:     ctx.Class_modifier().FINAL() != nil,
	}

	if ctx.Class_modifier().Privacy_modifier() != nil {
		attribute.PrivacyModifier = utils.PrivacyModifiers[v.VisitPrivacyModifier(ctx.Class_modifier().Privacy_modifier().(*parser.Privacy_modifierContext))]
	} else {
		attribute.PrivacyModifier = utils.PrivacyModifiers["public"]
	}

	class.AddAttribute(&attribute)

	return nil
}

func (v *FirstPassVisitor) VisitVariableDeclaration(ctx *parser.Variable_declarationContext) (variableName string, variableType string) {
	variableName = ctx.NAME().GetText()
	variableType = ctx.Type_().GetText()

	return
}

func (v *FirstPassVisitor) VisitPrivacyModifier(ctx *parser.Privacy_modifierContext) string {
	if ctx.PRIVATE() != nil {
		return "private"
	}
	if ctx.PROTECTED() != nil {
		return "protected"
	}
	return "public"
}
