package first_semantic_pass

import (
	parser "github.com/Brennon-Oliveira/flexar/grammar"
	"github.com/Brennon-Oliveira/flexar/program_data"
	"github.com/Brennon-Oliveira/flexar/utils"
)

func (v *FirstPassVisitor) VisitFunc(ctx *parser.FuncContext, exported bool, namespace *program_data.Namespace) interface{} {
	name := ctx.NAME().GetText()
	var params []program_data.FunctionParam

	if namespace.ExistsFunction(name) {
		utils.SemanticError(ctx.NAME().GetSymbol().GetLine(), "Function "+name+" already exists")
	}

	if ctx.Func_param() != nil {
		params = v.VisitFuncParam(ctx.Func_param().(*parser.Func_paramContext)).([]program_data.FunctionParam)
	}

	var returns []string

	if ctx.Func_return() != nil {
		returns = v.VisitFuncReturn(ctx.Func_return().(*parser.Func_returnContext))
	}

	function := program_data.Function{
		Name:    name,
		Returns: returns,
		Params:  params,
	}

	program_data.AddFunctionToNamespace(utils.GetCurrentNamespace(), function)

	return nil
}

func (v *FirstPassVisitor) VisitFuncParam(ctx *parser.Func_paramContext) interface{} {
	var params []program_data.FunctionParam
	for index, param := range ctx.AllFunc_param_rule() {
		newParam := v.VisitParamRule(param.(*parser.Func_param_ruleContext), index)
		// already in params
		for _, p := range params {
			if p.Name == newParam.Name {
				utils.SemanticError(param.(*parser.Func_param_ruleContext).NAME().GetSymbol().GetLine(), "Param "+p.Name+" already exists")
			}
		}

		params = append(params, newParam)
	}
	return params
}

func (v *FirstPassVisitor) VisitParamRule(ctx *parser.Func_param_ruleContext, order int) program_data.FunctionParam {

	funcParam := program_data.FunctionParam{
		Name:     ctx.NAME().GetText(),
		Type:     ctx.Type_().GetText(),
		Order:    order,
		Optional: ctx.QUESTION() != nil,
	}

	return funcParam
}

func (v *FirstPassVisitor) VisitFuncReturn(ctx *parser.Func_returnContext) []string {
	var returns []string
	for _, returnType := range ctx.AllType_() {
		returns = append(returns, returnType.GetText())
	}
	return returns
}
