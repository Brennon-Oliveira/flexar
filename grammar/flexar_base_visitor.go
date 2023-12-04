// Code generated from .//grammar//Flexar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Flexar

import "github.com/antlr4-go/antlr/v4"

type BaseFlexarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFlexarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitProgram_rule(ctx *Program_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitImport_group(ctx *Import_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitImport_rule(ctx *Import_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitImport_namespace(ctx *Import_namespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNamespace_name(ctx *Namespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNamespace_call(ctx *Namespace_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass(ctx *ClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_extends(ctx *Class_extendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_implements(ctx *Class_implementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_body(ctx *Class_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_body_rule(ctx *Class_body_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_attribute(ctx *Class_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitPrivacy_modifier(ctx *Privacy_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_method(ctx *Class_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAbstract_method(ctx *Abstract_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_modifier(ctx *Class_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitConstructor(ctx *ConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitClass_new_instance(ctx *Class_new_instanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitMethod_call(ctx *Method_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAttribute_call(ctx *Attribute_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface(ctx *InterfaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface_extends(ctx *Interface_extendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface_body(ctx *Interface_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface_body_rule(ctx *Interface_body_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface_method(ctx *Interface_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitInterface_attribute(ctx *Interface_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitStruct_body(ctx *Struct_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitStruct_attribute(ctx *Struct_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitEnum(ctx *EnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitEnum_body(ctx *Enum_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitEnum_attribute(ctx *Enum_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitComposed_value(ctx *Composed_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitList_value(ctx *List_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitMap_value(ctx *Map_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitTuple_value(ctx *Tuple_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNamed_tuple_value(ctx *Named_tuple_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAnd_expression(ctx *And_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitOr_expression(ctx *Or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitXor_expression(ctx *Xor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNot_expression(ctx *Not_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitExpression_math(ctx *Expression_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitTerm_math(ctx *Term_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFactor_math(ctx *Factor_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitBitwise_math(ctx *Bitwise_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitShift_math(ctx *Shift_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitUnary_math(ctx *Unary_mathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitBefore_unary(ctx *Before_unaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAfter_unary(ctx *After_unaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitMath_value(ctx *Math_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitParenthesis_expression(ctx *Parenthesis_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitBitwise_operator(ctx *Bitwise_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitShift_operator(ctx *Shift_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitTerm_operator(ctx *Term_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFactor_operator(ctx *Factor_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitComparision_operator(ctx *Comparision_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitNew_scope(ctx *New_scopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFor_rule(ctx *For_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFor_in(ctx *For_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFull_for(ctx *Full_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitDo_while_statement(ctx *Do_while_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitSwitch_statement(ctx *Switch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitSwitch_case(ctx *Switch_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitTry_statement(ctx *Try_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitCatch_statement(ctx *Catch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFinally_statement(ctx *Finally_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitElif_statement(ctx *Elif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitElse_statement(ctx *Else_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitStruct_func(ctx *Struct_funcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAnonymous_func(ctx *Anonymous_funcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_param(ctx *Func_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_callback(ctx *Func_callbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_param_callback(ctx *Func_param_callbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_param_rule(ctx *Func_param_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_call(ctx *Func_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_call_params(ctx *Func_call_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_return(ctx *Func_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFunc_body(ctx *Func_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitVariable_assign(ctx *Variable_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitAssing(ctx *AssingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlexarVisitor) VisitFinal_type(ctx *Final_typeContext) interface{} {
	return v.VisitChildren(ctx)
}
