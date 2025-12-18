package first_semantic_pass

import (
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/Brennon-Oliveira/flexar/program_data"
)

func (v *FirstPassVisitor) VisitProgram(ctx *parser.ProgramContext, app *program_data.Application) interface{} {
	namespace := v.VisitNamespace(ctx.Namespace().(*parser.NamespaceContext), app)
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
