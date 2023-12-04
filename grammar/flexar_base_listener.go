// Code generated from .//grammar//Flexar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Flexar

import "github.com/antlr4-go/antlr/v4"

// BaseFlexarListener is a complete listener for a parse tree produced by FlexarParser.
type BaseFlexarListener struct{}

var _ FlexarListener = &BaseFlexarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlexarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlexarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlexarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlexarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseFlexarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseFlexarListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgram_rule is called when production program_rule is entered.
func (s *BaseFlexarListener) EnterProgram_rule(ctx *Program_ruleContext) {}

// ExitProgram_rule is called when production program_rule is exited.
func (s *BaseFlexarListener) ExitProgram_rule(ctx *Program_ruleContext) {}

// EnterImport_group is called when production import_group is entered.
func (s *BaseFlexarListener) EnterImport_group(ctx *Import_groupContext) {}

// ExitImport_group is called when production import_group is exited.
func (s *BaseFlexarListener) ExitImport_group(ctx *Import_groupContext) {}

// EnterImport_rule is called when production import_rule is entered.
func (s *BaseFlexarListener) EnterImport_rule(ctx *Import_ruleContext) {}

// ExitImport_rule is called when production import_rule is exited.
func (s *BaseFlexarListener) ExitImport_rule(ctx *Import_ruleContext) {}

// EnterImport_namespace is called when production import_namespace is entered.
func (s *BaseFlexarListener) EnterImport_namespace(ctx *Import_namespaceContext) {}

// ExitImport_namespace is called when production import_namespace is exited.
func (s *BaseFlexarListener) ExitImport_namespace(ctx *Import_namespaceContext) {}

// EnterNamespace_name is called when production namespace_name is entered.
func (s *BaseFlexarListener) EnterNamespace_name(ctx *Namespace_nameContext) {}

// ExitNamespace_name is called when production namespace_name is exited.
func (s *BaseFlexarListener) ExitNamespace_name(ctx *Namespace_nameContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseFlexarListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseFlexarListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterNamespace_call is called when production namespace_call is entered.
func (s *BaseFlexarListener) EnterNamespace_call(ctx *Namespace_callContext) {}

// ExitNamespace_call is called when production namespace_call is exited.
func (s *BaseFlexarListener) ExitNamespace_call(ctx *Namespace_callContext) {}

// EnterClass is called when production class is entered.
func (s *BaseFlexarListener) EnterClass(ctx *ClassContext) {}

// ExitClass is called when production class is exited.
func (s *BaseFlexarListener) ExitClass(ctx *ClassContext) {}

// EnterClass_extends is called when production class_extends is entered.
func (s *BaseFlexarListener) EnterClass_extends(ctx *Class_extendsContext) {}

// ExitClass_extends is called when production class_extends is exited.
func (s *BaseFlexarListener) ExitClass_extends(ctx *Class_extendsContext) {}

// EnterClass_implements is called when production class_implements is entered.
func (s *BaseFlexarListener) EnterClass_implements(ctx *Class_implementsContext) {}

// ExitClass_implements is called when production class_implements is exited.
func (s *BaseFlexarListener) ExitClass_implements(ctx *Class_implementsContext) {}

// EnterClass_body is called when production class_body is entered.
func (s *BaseFlexarListener) EnterClass_body(ctx *Class_bodyContext) {}

// ExitClass_body is called when production class_body is exited.
func (s *BaseFlexarListener) ExitClass_body(ctx *Class_bodyContext) {}

// EnterClass_body_rule is called when production class_body_rule is entered.
func (s *BaseFlexarListener) EnterClass_body_rule(ctx *Class_body_ruleContext) {}

// ExitClass_body_rule is called when production class_body_rule is exited.
func (s *BaseFlexarListener) ExitClass_body_rule(ctx *Class_body_ruleContext) {}

// EnterClass_attribute is called when production class_attribute is entered.
func (s *BaseFlexarListener) EnterClass_attribute(ctx *Class_attributeContext) {}

// ExitClass_attribute is called when production class_attribute is exited.
func (s *BaseFlexarListener) ExitClass_attribute(ctx *Class_attributeContext) {}

// EnterPrivacy_modifier is called when production privacy_modifier is entered.
func (s *BaseFlexarListener) EnterPrivacy_modifier(ctx *Privacy_modifierContext) {}

// ExitPrivacy_modifier is called when production privacy_modifier is exited.
func (s *BaseFlexarListener) ExitPrivacy_modifier(ctx *Privacy_modifierContext) {}

// EnterClass_method is called when production class_method is entered.
func (s *BaseFlexarListener) EnterClass_method(ctx *Class_methodContext) {}

// ExitClass_method is called when production class_method is exited.
func (s *BaseFlexarListener) ExitClass_method(ctx *Class_methodContext) {}

// EnterAbstract_method is called when production abstract_method is entered.
func (s *BaseFlexarListener) EnterAbstract_method(ctx *Abstract_methodContext) {}

// ExitAbstract_method is called when production abstract_method is exited.
func (s *BaseFlexarListener) ExitAbstract_method(ctx *Abstract_methodContext) {}

// EnterClass_modifier is called when production class_modifier is entered.
func (s *BaseFlexarListener) EnterClass_modifier(ctx *Class_modifierContext) {}

// ExitClass_modifier is called when production class_modifier is exited.
func (s *BaseFlexarListener) ExitClass_modifier(ctx *Class_modifierContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BaseFlexarListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BaseFlexarListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterClass_new_instance is called when production class_new_instance is entered.
func (s *BaseFlexarListener) EnterClass_new_instance(ctx *Class_new_instanceContext) {}

// ExitClass_new_instance is called when production class_new_instance is exited.
func (s *BaseFlexarListener) ExitClass_new_instance(ctx *Class_new_instanceContext) {}

// EnterMethod_call is called when production method_call is entered.
func (s *BaseFlexarListener) EnterMethod_call(ctx *Method_callContext) {}

// ExitMethod_call is called when production method_call is exited.
func (s *BaseFlexarListener) ExitMethod_call(ctx *Method_callContext) {}

// EnterAttribute_call is called when production attribute_call is entered.
func (s *BaseFlexarListener) EnterAttribute_call(ctx *Attribute_callContext) {}

// ExitAttribute_call is called when production attribute_call is exited.
func (s *BaseFlexarListener) ExitAttribute_call(ctx *Attribute_callContext) {}

// EnterInterface is called when production interface is entered.
func (s *BaseFlexarListener) EnterInterface(ctx *InterfaceContext) {}

// ExitInterface is called when production interface is exited.
func (s *BaseFlexarListener) ExitInterface(ctx *InterfaceContext) {}

// EnterInterface_extends is called when production interface_extends is entered.
func (s *BaseFlexarListener) EnterInterface_extends(ctx *Interface_extendsContext) {}

// ExitInterface_extends is called when production interface_extends is exited.
func (s *BaseFlexarListener) ExitInterface_extends(ctx *Interface_extendsContext) {}

// EnterInterface_body is called when production interface_body is entered.
func (s *BaseFlexarListener) EnterInterface_body(ctx *Interface_bodyContext) {}

// ExitInterface_body is called when production interface_body is exited.
func (s *BaseFlexarListener) ExitInterface_body(ctx *Interface_bodyContext) {}

// EnterInterface_body_rule is called when production interface_body_rule is entered.
func (s *BaseFlexarListener) EnterInterface_body_rule(ctx *Interface_body_ruleContext) {}

// ExitInterface_body_rule is called when production interface_body_rule is exited.
func (s *BaseFlexarListener) ExitInterface_body_rule(ctx *Interface_body_ruleContext) {}

// EnterInterface_method is called when production interface_method is entered.
func (s *BaseFlexarListener) EnterInterface_method(ctx *Interface_methodContext) {}

// ExitInterface_method is called when production interface_method is exited.
func (s *BaseFlexarListener) ExitInterface_method(ctx *Interface_methodContext) {}

// EnterInterface_attribute is called when production interface_attribute is entered.
func (s *BaseFlexarListener) EnterInterface_attribute(ctx *Interface_attributeContext) {}

// ExitInterface_attribute is called when production interface_attribute is exited.
func (s *BaseFlexarListener) ExitInterface_attribute(ctx *Interface_attributeContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseFlexarListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseFlexarListener) ExitStruct(ctx *StructContext) {}

// EnterStruct_body is called when production struct_body is entered.
func (s *BaseFlexarListener) EnterStruct_body(ctx *Struct_bodyContext) {}

// ExitStruct_body is called when production struct_body is exited.
func (s *BaseFlexarListener) ExitStruct_body(ctx *Struct_bodyContext) {}

// EnterStruct_attribute is called when production struct_attribute is entered.
func (s *BaseFlexarListener) EnterStruct_attribute(ctx *Struct_attributeContext) {}

// ExitStruct_attribute is called when production struct_attribute is exited.
func (s *BaseFlexarListener) ExitStruct_attribute(ctx *Struct_attributeContext) {}

// EnterEnum is called when production enum is entered.
func (s *BaseFlexarListener) EnterEnum(ctx *EnumContext) {}

// ExitEnum is called when production enum is exited.
func (s *BaseFlexarListener) ExitEnum(ctx *EnumContext) {}

// EnterEnum_body is called when production enum_body is entered.
func (s *BaseFlexarListener) EnterEnum_body(ctx *Enum_bodyContext) {}

// ExitEnum_body is called when production enum_body is exited.
func (s *BaseFlexarListener) ExitEnum_body(ctx *Enum_bodyContext) {}

// EnterEnum_attribute is called when production enum_attribute is entered.
func (s *BaseFlexarListener) EnterEnum_attribute(ctx *Enum_attributeContext) {}

// ExitEnum_attribute is called when production enum_attribute is exited.
func (s *BaseFlexarListener) ExitEnum_attribute(ctx *Enum_attributeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFlexarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFlexarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComposed_value is called when production composed_value is entered.
func (s *BaseFlexarListener) EnterComposed_value(ctx *Composed_valueContext) {}

// ExitComposed_value is called when production composed_value is exited.
func (s *BaseFlexarListener) ExitComposed_value(ctx *Composed_valueContext) {}

// EnterList_value is called when production list_value is entered.
func (s *BaseFlexarListener) EnterList_value(ctx *List_valueContext) {}

// ExitList_value is called when production list_value is exited.
func (s *BaseFlexarListener) ExitList_value(ctx *List_valueContext) {}

// EnterMap_value is called when production map_value is entered.
func (s *BaseFlexarListener) EnterMap_value(ctx *Map_valueContext) {}

// ExitMap_value is called when production map_value is exited.
func (s *BaseFlexarListener) ExitMap_value(ctx *Map_valueContext) {}

// EnterTuple_value is called when production tuple_value is entered.
func (s *BaseFlexarListener) EnterTuple_value(ctx *Tuple_valueContext) {}

// ExitTuple_value is called when production tuple_value is exited.
func (s *BaseFlexarListener) ExitTuple_value(ctx *Tuple_valueContext) {}

// EnterNamed_tuple_value is called when production named_tuple_value is entered.
func (s *BaseFlexarListener) EnterNamed_tuple_value(ctx *Named_tuple_valueContext) {}

// ExitNamed_tuple_value is called when production named_tuple_value is exited.
func (s *BaseFlexarListener) ExitNamed_tuple_value(ctx *Named_tuple_valueContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseFlexarListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseFlexarListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterOr_expression is called when production or_expression is entered.
func (s *BaseFlexarListener) EnterOr_expression(ctx *Or_expressionContext) {}

// ExitOr_expression is called when production or_expression is exited.
func (s *BaseFlexarListener) ExitOr_expression(ctx *Or_expressionContext) {}

// EnterXor_expression is called when production xor_expression is entered.
func (s *BaseFlexarListener) EnterXor_expression(ctx *Xor_expressionContext) {}

// ExitXor_expression is called when production xor_expression is exited.
func (s *BaseFlexarListener) ExitXor_expression(ctx *Xor_expressionContext) {}

// EnterNot_expression is called when production not_expression is entered.
func (s *BaseFlexarListener) EnterNot_expression(ctx *Not_expressionContext) {}

// ExitNot_expression is called when production not_expression is exited.
func (s *BaseFlexarListener) ExitNot_expression(ctx *Not_expressionContext) {}

// EnterExpression_math is called when production expression_math is entered.
func (s *BaseFlexarListener) EnterExpression_math(ctx *Expression_mathContext) {}

// ExitExpression_math is called when production expression_math is exited.
func (s *BaseFlexarListener) ExitExpression_math(ctx *Expression_mathContext) {}

// EnterTerm_math is called when production term_math is entered.
func (s *BaseFlexarListener) EnterTerm_math(ctx *Term_mathContext) {}

// ExitTerm_math is called when production term_math is exited.
func (s *BaseFlexarListener) ExitTerm_math(ctx *Term_mathContext) {}

// EnterFactor_math is called when production factor_math is entered.
func (s *BaseFlexarListener) EnterFactor_math(ctx *Factor_mathContext) {}

// ExitFactor_math is called when production factor_math is exited.
func (s *BaseFlexarListener) ExitFactor_math(ctx *Factor_mathContext) {}

// EnterBitwise_math is called when production bitwise_math is entered.
func (s *BaseFlexarListener) EnterBitwise_math(ctx *Bitwise_mathContext) {}

// ExitBitwise_math is called when production bitwise_math is exited.
func (s *BaseFlexarListener) ExitBitwise_math(ctx *Bitwise_mathContext) {}

// EnterShift_math is called when production shift_math is entered.
func (s *BaseFlexarListener) EnterShift_math(ctx *Shift_mathContext) {}

// ExitShift_math is called when production shift_math is exited.
func (s *BaseFlexarListener) ExitShift_math(ctx *Shift_mathContext) {}

// EnterUnary_math is called when production unary_math is entered.
func (s *BaseFlexarListener) EnterUnary_math(ctx *Unary_mathContext) {}

// ExitUnary_math is called when production unary_math is exited.
func (s *BaseFlexarListener) ExitUnary_math(ctx *Unary_mathContext) {}

// EnterBefore_unary is called when production before_unary is entered.
func (s *BaseFlexarListener) EnterBefore_unary(ctx *Before_unaryContext) {}

// ExitBefore_unary is called when production before_unary is exited.
func (s *BaseFlexarListener) ExitBefore_unary(ctx *Before_unaryContext) {}

// EnterAfter_unary is called when production after_unary is entered.
func (s *BaseFlexarListener) EnterAfter_unary(ctx *After_unaryContext) {}

// ExitAfter_unary is called when production after_unary is exited.
func (s *BaseFlexarListener) ExitAfter_unary(ctx *After_unaryContext) {}

// EnterMath_value is called when production math_value is entered.
func (s *BaseFlexarListener) EnterMath_value(ctx *Math_valueContext) {}

// ExitMath_value is called when production math_value is exited.
func (s *BaseFlexarListener) ExitMath_value(ctx *Math_valueContext) {}

// EnterParenthesis_expression is called when production parenthesis_expression is entered.
func (s *BaseFlexarListener) EnterParenthesis_expression(ctx *Parenthesis_expressionContext) {}

// ExitParenthesis_expression is called when production parenthesis_expression is exited.
func (s *BaseFlexarListener) ExitParenthesis_expression(ctx *Parenthesis_expressionContext) {}

// EnterBitwise_operator is called when production bitwise_operator is entered.
func (s *BaseFlexarListener) EnterBitwise_operator(ctx *Bitwise_operatorContext) {}

// ExitBitwise_operator is called when production bitwise_operator is exited.
func (s *BaseFlexarListener) ExitBitwise_operator(ctx *Bitwise_operatorContext) {}

// EnterShift_operator is called when production shift_operator is entered.
func (s *BaseFlexarListener) EnterShift_operator(ctx *Shift_operatorContext) {}

// ExitShift_operator is called when production shift_operator is exited.
func (s *BaseFlexarListener) ExitShift_operator(ctx *Shift_operatorContext) {}

// EnterTerm_operator is called when production term_operator is entered.
func (s *BaseFlexarListener) EnterTerm_operator(ctx *Term_operatorContext) {}

// ExitTerm_operator is called when production term_operator is exited.
func (s *BaseFlexarListener) ExitTerm_operator(ctx *Term_operatorContext) {}

// EnterFactor_operator is called when production factor_operator is entered.
func (s *BaseFlexarListener) EnterFactor_operator(ctx *Factor_operatorContext) {}

// ExitFactor_operator is called when production factor_operator is exited.
func (s *BaseFlexarListener) ExitFactor_operator(ctx *Factor_operatorContext) {}

// EnterComparision_operator is called when production comparision_operator is entered.
func (s *BaseFlexarListener) EnterComparision_operator(ctx *Comparision_operatorContext) {}

// ExitComparision_operator is called when production comparision_operator is exited.
func (s *BaseFlexarListener) ExitComparision_operator(ctx *Comparision_operatorContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFlexarListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFlexarListener) ExitValue(ctx *ValueContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFlexarListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFlexarListener) ExitStatement(ctx *StatementContext) {}

// EnterNew_scope is called when production new_scope is entered.
func (s *BaseFlexarListener) EnterNew_scope(ctx *New_scopeContext) {}

// ExitNew_scope is called when production new_scope is exited.
func (s *BaseFlexarListener) ExitNew_scope(ctx *New_scopeContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseFlexarListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseFlexarListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterFor_rule is called when production for_rule is entered.
func (s *BaseFlexarListener) EnterFor_rule(ctx *For_ruleContext) {}

// ExitFor_rule is called when production for_rule is exited.
func (s *BaseFlexarListener) ExitFor_rule(ctx *For_ruleContext) {}

// EnterFor_in is called when production for_in is entered.
func (s *BaseFlexarListener) EnterFor_in(ctx *For_inContext) {}

// ExitFor_in is called when production for_in is exited.
func (s *BaseFlexarListener) ExitFor_in(ctx *For_inContext) {}

// EnterFull_for is called when production full_for is entered.
func (s *BaseFlexarListener) EnterFull_for(ctx *Full_forContext) {}

// ExitFull_for is called when production full_for is exited.
func (s *BaseFlexarListener) ExitFull_for(ctx *Full_forContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseFlexarListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseFlexarListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterDo_while_statement is called when production do_while_statement is entered.
func (s *BaseFlexarListener) EnterDo_while_statement(ctx *Do_while_statementContext) {}

// ExitDo_while_statement is called when production do_while_statement is exited.
func (s *BaseFlexarListener) ExitDo_while_statement(ctx *Do_while_statementContext) {}

// EnterSwitch_statement is called when production switch_statement is entered.
func (s *BaseFlexarListener) EnterSwitch_statement(ctx *Switch_statementContext) {}

// ExitSwitch_statement is called when production switch_statement is exited.
func (s *BaseFlexarListener) ExitSwitch_statement(ctx *Switch_statementContext) {}

// EnterSwitch_case is called when production switch_case is entered.
func (s *BaseFlexarListener) EnterSwitch_case(ctx *Switch_caseContext) {}

// ExitSwitch_case is called when production switch_case is exited.
func (s *BaseFlexarListener) ExitSwitch_case(ctx *Switch_caseContext) {}

// EnterTry_statement is called when production try_statement is entered.
func (s *BaseFlexarListener) EnterTry_statement(ctx *Try_statementContext) {}

// ExitTry_statement is called when production try_statement is exited.
func (s *BaseFlexarListener) ExitTry_statement(ctx *Try_statementContext) {}

// EnterCatch_statement is called when production catch_statement is entered.
func (s *BaseFlexarListener) EnterCatch_statement(ctx *Catch_statementContext) {}

// ExitCatch_statement is called when production catch_statement is exited.
func (s *BaseFlexarListener) ExitCatch_statement(ctx *Catch_statementContext) {}

// EnterFinally_statement is called when production finally_statement is entered.
func (s *BaseFlexarListener) EnterFinally_statement(ctx *Finally_statementContext) {}

// ExitFinally_statement is called when production finally_statement is exited.
func (s *BaseFlexarListener) ExitFinally_statement(ctx *Finally_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseFlexarListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseFlexarListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElif_statement is called when production elif_statement is entered.
func (s *BaseFlexarListener) EnterElif_statement(ctx *Elif_statementContext) {}

// ExitElif_statement is called when production elif_statement is exited.
func (s *BaseFlexarListener) ExitElif_statement(ctx *Elif_statementContext) {}

// EnterElse_statement is called when production else_statement is entered.
func (s *BaseFlexarListener) EnterElse_statement(ctx *Else_statementContext) {}

// ExitElse_statement is called when production else_statement is exited.
func (s *BaseFlexarListener) ExitElse_statement(ctx *Else_statementContext) {}

// EnterStruct_func is called when production struct_func is entered.
func (s *BaseFlexarListener) EnterStruct_func(ctx *Struct_funcContext) {}

// ExitStruct_func is called when production struct_func is exited.
func (s *BaseFlexarListener) ExitStruct_func(ctx *Struct_funcContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseFlexarListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseFlexarListener) ExitFunc(ctx *FuncContext) {}

// EnterAnonymous_func is called when production anonymous_func is entered.
func (s *BaseFlexarListener) EnterAnonymous_func(ctx *Anonymous_funcContext) {}

// ExitAnonymous_func is called when production anonymous_func is exited.
func (s *BaseFlexarListener) ExitAnonymous_func(ctx *Anonymous_funcContext) {}

// EnterFunc_param is called when production func_param is entered.
func (s *BaseFlexarListener) EnterFunc_param(ctx *Func_paramContext) {}

// ExitFunc_param is called when production func_param is exited.
func (s *BaseFlexarListener) ExitFunc_param(ctx *Func_paramContext) {}

// EnterFunc_callback is called when production func_callback is entered.
func (s *BaseFlexarListener) EnterFunc_callback(ctx *Func_callbackContext) {}

// ExitFunc_callback is called when production func_callback is exited.
func (s *BaseFlexarListener) ExitFunc_callback(ctx *Func_callbackContext) {}

// EnterFunc_param_callback is called when production func_param_callback is entered.
func (s *BaseFlexarListener) EnterFunc_param_callback(ctx *Func_param_callbackContext) {}

// ExitFunc_param_callback is called when production func_param_callback is exited.
func (s *BaseFlexarListener) ExitFunc_param_callback(ctx *Func_param_callbackContext) {}

// EnterFunc_param_rule is called when production func_param_rule is entered.
func (s *BaseFlexarListener) EnterFunc_param_rule(ctx *Func_param_ruleContext) {}

// ExitFunc_param_rule is called when production func_param_rule is exited.
func (s *BaseFlexarListener) ExitFunc_param_rule(ctx *Func_param_ruleContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseFlexarListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseFlexarListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterFunc_call_params is called when production func_call_params is entered.
func (s *BaseFlexarListener) EnterFunc_call_params(ctx *Func_call_paramsContext) {}

// ExitFunc_call_params is called when production func_call_params is exited.
func (s *BaseFlexarListener) ExitFunc_call_params(ctx *Func_call_paramsContext) {}

// EnterFunc_return is called when production func_return is entered.
func (s *BaseFlexarListener) EnterFunc_return(ctx *Func_returnContext) {}

// ExitFunc_return is called when production func_return is exited.
func (s *BaseFlexarListener) ExitFunc_return(ctx *Func_returnContext) {}

// EnterFunc_body is called when production func_body is entered.
func (s *BaseFlexarListener) EnterFunc_body(ctx *Func_bodyContext) {}

// ExitFunc_body is called when production func_body is exited.
func (s *BaseFlexarListener) ExitFunc_body(ctx *Func_bodyContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseFlexarListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseFlexarListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterVariable_declaration is called when production variable_declaration is entered.
func (s *BaseFlexarListener) EnterVariable_declaration(ctx *Variable_declarationContext) {}

// ExitVariable_declaration is called when production variable_declaration is exited.
func (s *BaseFlexarListener) ExitVariable_declaration(ctx *Variable_declarationContext) {}

// EnterVariable_assign is called when production variable_assign is entered.
func (s *BaseFlexarListener) EnterVariable_assign(ctx *Variable_assignContext) {}

// ExitVariable_assign is called when production variable_assign is exited.
func (s *BaseFlexarListener) ExitVariable_assign(ctx *Variable_assignContext) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BaseFlexarListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BaseFlexarListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterAssing is called when production assing is entered.
func (s *BaseFlexarListener) EnterAssing(ctx *AssingContext) {}

// ExitAssing is called when production assing is exited.
func (s *BaseFlexarListener) ExitAssing(ctx *AssingContext) {}

// EnterType is called when production type is entered.
func (s *BaseFlexarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseFlexarListener) ExitType(ctx *TypeContext) {}

// EnterFinal_type is called when production final_type is entered.
func (s *BaseFlexarListener) EnterFinal_type(ctx *Final_typeContext) {}

// ExitFinal_type is called when production final_type is exited.
func (s *BaseFlexarListener) ExitFinal_type(ctx *Final_typeContext) {}
