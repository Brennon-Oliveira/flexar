// Code generated from .//grammar//Flexar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Flexar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FlexarParser.
type FlexarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FlexarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FlexarParser#program_rule.
	VisitProgram_rule(ctx *Program_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#import_group.
	VisitImport_group(ctx *Import_groupContext) interface{}

	// Visit a parse tree produced by FlexarParser#import_rule.
	VisitImport_rule(ctx *Import_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#import_namespace.
	VisitImport_namespace(ctx *Import_namespaceContext) interface{}

	// Visit a parse tree produced by FlexarParser#namespace_name.
	VisitNamespace_name(ctx *Namespace_nameContext) interface{}

	// Visit a parse tree produced by FlexarParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by FlexarParser#namespace_call.
	VisitNamespace_call(ctx *Namespace_callContext) interface{}

	// Visit a parse tree produced by FlexarParser#class.
	VisitClass(ctx *ClassContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_extends.
	VisitClass_extends(ctx *Class_extendsContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_implements.
	VisitClass_implements(ctx *Class_implementsContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_body.
	VisitClass_body(ctx *Class_bodyContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_body_rule.
	VisitClass_body_rule(ctx *Class_body_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_attribute.
	VisitClass_attribute(ctx *Class_attributeContext) interface{}

	// Visit a parse tree produced by FlexarParser#privacy_modifier.
	VisitPrivacy_modifier(ctx *Privacy_modifierContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_method.
	VisitClass_method(ctx *Class_methodContext) interface{}

	// Visit a parse tree produced by FlexarParser#abstract_method.
	VisitAbstract_method(ctx *Abstract_methodContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_modifier.
	VisitClass_modifier(ctx *Class_modifierContext) interface{}

	// Visit a parse tree produced by FlexarParser#constructor.
	VisitConstructor(ctx *ConstructorContext) interface{}

	// Visit a parse tree produced by FlexarParser#class_new_instance.
	VisitClass_new_instance(ctx *Class_new_instanceContext) interface{}

	// Visit a parse tree produced by FlexarParser#method_call.
	VisitMethod_call(ctx *Method_callContext) interface{}

	// Visit a parse tree produced by FlexarParser#attribute_call.
	VisitAttribute_call(ctx *Attribute_callContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface.
	VisitInterface(ctx *InterfaceContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface_extends.
	VisitInterface_extends(ctx *Interface_extendsContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface_body.
	VisitInterface_body(ctx *Interface_bodyContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface_body_rule.
	VisitInterface_body_rule(ctx *Interface_body_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface_method.
	VisitInterface_method(ctx *Interface_methodContext) interface{}

	// Visit a parse tree produced by FlexarParser#interface_attribute.
	VisitInterface_attribute(ctx *Interface_attributeContext) interface{}

	// Visit a parse tree produced by FlexarParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by FlexarParser#struct_body.
	VisitStruct_body(ctx *Struct_bodyContext) interface{}

	// Visit a parse tree produced by FlexarParser#struct_attribute.
	VisitStruct_attribute(ctx *Struct_attributeContext) interface{}

	// Visit a parse tree produced by FlexarParser#enum.
	VisitEnum(ctx *EnumContext) interface{}

	// Visit a parse tree produced by FlexarParser#enum_body.
	VisitEnum_body(ctx *Enum_bodyContext) interface{}

	// Visit a parse tree produced by FlexarParser#enum_attribute.
	VisitEnum_attribute(ctx *Enum_attributeContext) interface{}

	// Visit a parse tree produced by FlexarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#composed_value.
	VisitComposed_value(ctx *Composed_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#list_value.
	VisitList_value(ctx *List_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#map_value.
	VisitMap_value(ctx *Map_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#tuple_value.
	VisitTuple_value(ctx *Tuple_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#named_tuple_value.
	VisitNamed_tuple_value(ctx *Named_tuple_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#and_expression.
	VisitAnd_expression(ctx *And_expressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#or_expression.
	VisitOr_expression(ctx *Or_expressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#xor_expression.
	VisitXor_expression(ctx *Xor_expressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#not_expression.
	VisitNot_expression(ctx *Not_expressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#expression_math.
	VisitExpression_math(ctx *Expression_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#term_math.
	VisitTerm_math(ctx *Term_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#factor_math.
	VisitFactor_math(ctx *Factor_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#bitwise_math.
	VisitBitwise_math(ctx *Bitwise_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#shift_math.
	VisitShift_math(ctx *Shift_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#unary_math.
	VisitUnary_math(ctx *Unary_mathContext) interface{}

	// Visit a parse tree produced by FlexarParser#before_unary.
	VisitBefore_unary(ctx *Before_unaryContext) interface{}

	// Visit a parse tree produced by FlexarParser#after_unary.
	VisitAfter_unary(ctx *After_unaryContext) interface{}

	// Visit a parse tree produced by FlexarParser#math_value.
	VisitMath_value(ctx *Math_valueContext) interface{}

	// Visit a parse tree produced by FlexarParser#parenthesis_expression.
	VisitParenthesis_expression(ctx *Parenthesis_expressionContext) interface{}

	// Visit a parse tree produced by FlexarParser#bitwise_operator.
	VisitBitwise_operator(ctx *Bitwise_operatorContext) interface{}

	// Visit a parse tree produced by FlexarParser#shift_operator.
	VisitShift_operator(ctx *Shift_operatorContext) interface{}

	// Visit a parse tree produced by FlexarParser#term_operator.
	VisitTerm_operator(ctx *Term_operatorContext) interface{}

	// Visit a parse tree produced by FlexarParser#factor_operator.
	VisitFactor_operator(ctx *Factor_operatorContext) interface{}

	// Visit a parse tree produced by FlexarParser#comparision_operator.
	VisitComparision_operator(ctx *Comparision_operatorContext) interface{}

	// Visit a parse tree produced by FlexarParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FlexarParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by FlexarParser#new_scope.
	VisitNew_scope(ctx *New_scopeContext) interface{}

	// Visit a parse tree produced by FlexarParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#for_rule.
	VisitFor_rule(ctx *For_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#for_in.
	VisitFor_in(ctx *For_inContext) interface{}

	// Visit a parse tree produced by FlexarParser#full_for.
	VisitFull_for(ctx *Full_forContext) interface{}

	// Visit a parse tree produced by FlexarParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#do_while_statement.
	VisitDo_while_statement(ctx *Do_while_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#switch_statement.
	VisitSwitch_statement(ctx *Switch_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#switch_case.
	VisitSwitch_case(ctx *Switch_caseContext) interface{}

	// Visit a parse tree produced by FlexarParser#try_statement.
	VisitTry_statement(ctx *Try_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#catch_statement.
	VisitCatch_statement(ctx *Catch_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#finally_statement.
	VisitFinally_statement(ctx *Finally_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#elif_statement.
	VisitElif_statement(ctx *Elif_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#else_statement.
	VisitElse_statement(ctx *Else_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#struct_func.
	VisitStruct_func(ctx *Struct_funcContext) interface{}

	// Visit a parse tree produced by FlexarParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by FlexarParser#anonymous_func.
	VisitAnonymous_func(ctx *Anonymous_funcContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_param.
	VisitFunc_param(ctx *Func_paramContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_callback.
	VisitFunc_callback(ctx *Func_callbackContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_param_callback.
	VisitFunc_param_callback(ctx *Func_param_callbackContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_param_rule.
	VisitFunc_param_rule(ctx *Func_param_ruleContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_call.
	VisitFunc_call(ctx *Func_callContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_call_params.
	VisitFunc_call_params(ctx *Func_call_paramsContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_return.
	VisitFunc_return(ctx *Func_returnContext) interface{}

	// Visit a parse tree produced by FlexarParser#func_body.
	VisitFunc_body(ctx *Func_bodyContext) interface{}

	// Visit a parse tree produced by FlexarParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by FlexarParser#variable_declaration.
	VisitVariable_declaration(ctx *Variable_declarationContext) interface{}

	// Visit a parse tree produced by FlexarParser#variable_assign.
	VisitVariable_assign(ctx *Variable_assignContext) interface{}

	// Visit a parse tree produced by FlexarParser#variable_name.
	VisitVariable_name(ctx *Variable_nameContext) interface{}

	// Visit a parse tree produced by FlexarParser#assing.
	VisitAssing(ctx *AssingContext) interface{}

	// Visit a parse tree produced by FlexarParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by FlexarParser#final_type.
	VisitFinal_type(ctx *Final_typeContext) interface{}
}
