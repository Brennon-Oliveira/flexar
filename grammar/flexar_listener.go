// Code generated from ./grammar/Flexar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Flexar

import "github.com/antlr4-go/antlr/v4"

// FlexarListener is a complete listener for a parse tree produced by FlexarParser.
type FlexarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgram_rule is called when entering the program_rule production.
	EnterProgram_rule(c *Program_ruleContext)

	// EnterImport_group is called when entering the import_group production.
	EnterImport_group(c *Import_groupContext)

	// EnterImport_rule is called when entering the import_rule production.
	EnterImport_rule(c *Import_ruleContext)

	// EnterImport_namespace is called when entering the import_namespace production.
	EnterImport_namespace(c *Import_namespaceContext)

	// EnterNamespace_name is called when entering the namespace_name production.
	EnterNamespace_name(c *Namespace_nameContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterNamespace_call is called when entering the namespace_call production.
	EnterNamespace_call(c *Namespace_callContext)

	// EnterClass is called when entering the class production.
	EnterClass(c *ClassContext)

	// EnterClass_extends is called when entering the class_extends production.
	EnterClass_extends(c *Class_extendsContext)

	// EnterClass_implements is called when entering the class_implements production.
	EnterClass_implements(c *Class_implementsContext)

	// EnterClass_body is called when entering the class_body production.
	EnterClass_body(c *Class_bodyContext)

	// EnterClass_attribute is called when entering the class_attribute production.
	EnterClass_attribute(c *Class_attributeContext)

	// EnterPrivacy_modifier is called when entering the privacy_modifier production.
	EnterPrivacy_modifier(c *Privacy_modifierContext)

	// EnterClass_method is called when entering the class_method production.
	EnterClass_method(c *Class_methodContext)

	// EnterAbstract_method is called when entering the abstract_method production.
	EnterAbstract_method(c *Abstract_methodContext)

	// EnterClass_modifier is called when entering the class_modifier production.
	EnterClass_modifier(c *Class_modifierContext)

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterClass_new_instance is called when entering the class_new_instance production.
	EnterClass_new_instance(c *Class_new_instanceContext)

	// EnterMethod_call is called when entering the method_call production.
	EnterMethod_call(c *Method_callContext)

	// EnterAttribute_call is called when entering the attribute_call production.
	EnterAttribute_call(c *Attribute_callContext)

	// EnterInterface is called when entering the interface production.
	EnterInterface(c *InterfaceContext)

	// EnterInterface_extends is called when entering the interface_extends production.
	EnterInterface_extends(c *Interface_extendsContext)

	// EnterInterface_body is called when entering the interface_body production.
	EnterInterface_body(c *Interface_bodyContext)

	// EnterInterface_body_rule is called when entering the interface_body_rule production.
	EnterInterface_body_rule(c *Interface_body_ruleContext)

	// EnterInterface_method is called when entering the interface_method production.
	EnterInterface_method(c *Interface_methodContext)

	// EnterInterface_attribute is called when entering the interface_attribute production.
	EnterInterface_attribute(c *Interface_attributeContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterStruct_body is called when entering the struct_body production.
	EnterStruct_body(c *Struct_bodyContext)

	// EnterStruct_attribute is called when entering the struct_attribute production.
	EnterStruct_attribute(c *Struct_attributeContext)

	// EnterEnum is called when entering the enum production.
	EnterEnum(c *EnumContext)

	// EnterEnum_body is called when entering the enum_body production.
	EnterEnum_body(c *Enum_bodyContext)

	// EnterEnum_attribute is called when entering the enum_attribute production.
	EnterEnum_attribute(c *Enum_attributeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComposed_value is called when entering the composed_value production.
	EnterComposed_value(c *Composed_valueContext)

	// EnterList_value is called when entering the list_value production.
	EnterList_value(c *List_valueContext)

	// EnterMap_value is called when entering the map_value production.
	EnterMap_value(c *Map_valueContext)

	// EnterTuple_value is called when entering the tuple_value production.
	EnterTuple_value(c *Tuple_valueContext)

	// EnterNamed_tuple_value is called when entering the named_tuple_value production.
	EnterNamed_tuple_value(c *Named_tuple_valueContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterOr_expression is called when entering the or_expression production.
	EnterOr_expression(c *Or_expressionContext)

	// EnterXor_expression is called when entering the xor_expression production.
	EnterXor_expression(c *Xor_expressionContext)

	// EnterNot_expression is called when entering the not_expression production.
	EnterNot_expression(c *Not_expressionContext)

	// EnterExpression_math is called when entering the expression_math production.
	EnterExpression_math(c *Expression_mathContext)

	// EnterTerm_math is called when entering the term_math production.
	EnterTerm_math(c *Term_mathContext)

	// EnterFactor_math is called when entering the factor_math production.
	EnterFactor_math(c *Factor_mathContext)

	// EnterBitwise_math is called when entering the bitwise_math production.
	EnterBitwise_math(c *Bitwise_mathContext)

	// EnterShift_math is called when entering the shift_math production.
	EnterShift_math(c *Shift_mathContext)

	// EnterUnary_math is called when entering the unary_math production.
	EnterUnary_math(c *Unary_mathContext)

	// EnterBefore_unary is called when entering the before_unary production.
	EnterBefore_unary(c *Before_unaryContext)

	// EnterAfter_unary is called when entering the after_unary production.
	EnterAfter_unary(c *After_unaryContext)

	// EnterMath_value is called when entering the math_value production.
	EnterMath_value(c *Math_valueContext)

	// EnterParenthesis_expression is called when entering the parenthesis_expression production.
	EnterParenthesis_expression(c *Parenthesis_expressionContext)

	// EnterBitwise_operator is called when entering the bitwise_operator production.
	EnterBitwise_operator(c *Bitwise_operatorContext)

	// EnterShift_operator is called when entering the shift_operator production.
	EnterShift_operator(c *Shift_operatorContext)

	// EnterTerm_operator is called when entering the term_operator production.
	EnterTerm_operator(c *Term_operatorContext)

	// EnterFactor_operator is called when entering the factor_operator production.
	EnterFactor_operator(c *Factor_operatorContext)

	// EnterComparision_operator is called when entering the comparision_operator production.
	EnterComparision_operator(c *Comparision_operatorContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNew_scope is called when entering the new_scope production.
	EnterNew_scope(c *New_scopeContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterFor_rule is called when entering the for_rule production.
	EnterFor_rule(c *For_ruleContext)

	// EnterFor_in is called when entering the for_in production.
	EnterFor_in(c *For_inContext)

	// EnterFull_for is called when entering the full_for production.
	EnterFull_for(c *Full_forContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterDo_while_statement is called when entering the do_while_statement production.
	EnterDo_while_statement(c *Do_while_statementContext)

	// EnterSwitch_statement is called when entering the switch_statement production.
	EnterSwitch_statement(c *Switch_statementContext)

	// EnterSwitch_case is called when entering the switch_case production.
	EnterSwitch_case(c *Switch_caseContext)

	// EnterTry_statement is called when entering the try_statement production.
	EnterTry_statement(c *Try_statementContext)

	// EnterCatch_statement is called when entering the catch_statement production.
	EnterCatch_statement(c *Catch_statementContext)

	// EnterFinally_statement is called when entering the finally_statement production.
	EnterFinally_statement(c *Finally_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterElif_statement is called when entering the elif_statement production.
	EnterElif_statement(c *Elif_statementContext)

	// EnterElse_statement is called when entering the else_statement production.
	EnterElse_statement(c *Else_statementContext)

	// EnterStruct_func is called when entering the struct_func production.
	EnterStruct_func(c *Struct_funcContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterAnonymous_func is called when entering the anonymous_func production.
	EnterAnonymous_func(c *Anonymous_funcContext)

	// EnterFunc_param is called when entering the func_param production.
	EnterFunc_param(c *Func_paramContext)

	// EnterFunc_callback is called when entering the func_callback production.
	EnterFunc_callback(c *Func_callbackContext)

	// EnterFunc_param_callback is called when entering the func_param_callback production.
	EnterFunc_param_callback(c *Func_param_callbackContext)

	// EnterFunc_param_rule is called when entering the func_param_rule production.
	EnterFunc_param_rule(c *Func_param_ruleContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterFunc_call_params is called when entering the func_call_params production.
	EnterFunc_call_params(c *Func_call_paramsContext)

	// EnterFunc_return is called when entering the func_return production.
	EnterFunc_return(c *Func_returnContext)

	// EnterFunc_body is called when entering the func_body production.
	EnterFunc_body(c *Func_bodyContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterVariable_declaration is called when entering the variable_declaration production.
	EnterVariable_declaration(c *Variable_declarationContext)

	// EnterVariable_assign is called when entering the variable_assign production.
	EnterVariable_assign(c *Variable_assignContext)

	// EnterVariable_name is called when entering the variable_name production.
	EnterVariable_name(c *Variable_nameContext)

	// EnterAssing is called when entering the assing production.
	EnterAssing(c *AssingContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterFinal_type is called when entering the final_type production.
	EnterFinal_type(c *Final_typeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgram_rule is called when exiting the program_rule production.
	ExitProgram_rule(c *Program_ruleContext)

	// ExitImport_group is called when exiting the import_group production.
	ExitImport_group(c *Import_groupContext)

	// ExitImport_rule is called when exiting the import_rule production.
	ExitImport_rule(c *Import_ruleContext)

	// ExitImport_namespace is called when exiting the import_namespace production.
	ExitImport_namespace(c *Import_namespaceContext)

	// ExitNamespace_name is called when exiting the namespace_name production.
	ExitNamespace_name(c *Namespace_nameContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitNamespace_call is called when exiting the namespace_call production.
	ExitNamespace_call(c *Namespace_callContext)

	// ExitClass is called when exiting the class production.
	ExitClass(c *ClassContext)

	// ExitClass_extends is called when exiting the class_extends production.
	ExitClass_extends(c *Class_extendsContext)

	// ExitClass_implements is called when exiting the class_implements production.
	ExitClass_implements(c *Class_implementsContext)

	// ExitClass_body is called when exiting the class_body production.
	ExitClass_body(c *Class_bodyContext)

	// ExitClass_attribute is called when exiting the class_attribute production.
	ExitClass_attribute(c *Class_attributeContext)

	// ExitPrivacy_modifier is called when exiting the privacy_modifier production.
	ExitPrivacy_modifier(c *Privacy_modifierContext)

	// ExitClass_method is called when exiting the class_method production.
	ExitClass_method(c *Class_methodContext)

	// ExitAbstract_method is called when exiting the abstract_method production.
	ExitAbstract_method(c *Abstract_methodContext)

	// ExitClass_modifier is called when exiting the class_modifier production.
	ExitClass_modifier(c *Class_modifierContext)

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitClass_new_instance is called when exiting the class_new_instance production.
	ExitClass_new_instance(c *Class_new_instanceContext)

	// ExitMethod_call is called when exiting the method_call production.
	ExitMethod_call(c *Method_callContext)

	// ExitAttribute_call is called when exiting the attribute_call production.
	ExitAttribute_call(c *Attribute_callContext)

	// ExitInterface is called when exiting the interface production.
	ExitInterface(c *InterfaceContext)

	// ExitInterface_extends is called when exiting the interface_extends production.
	ExitInterface_extends(c *Interface_extendsContext)

	// ExitInterface_body is called when exiting the interface_body production.
	ExitInterface_body(c *Interface_bodyContext)

	// ExitInterface_body_rule is called when exiting the interface_body_rule production.
	ExitInterface_body_rule(c *Interface_body_ruleContext)

	// ExitInterface_method is called when exiting the interface_method production.
	ExitInterface_method(c *Interface_methodContext)

	// ExitInterface_attribute is called when exiting the interface_attribute production.
	ExitInterface_attribute(c *Interface_attributeContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitStruct_body is called when exiting the struct_body production.
	ExitStruct_body(c *Struct_bodyContext)

	// ExitStruct_attribute is called when exiting the struct_attribute production.
	ExitStruct_attribute(c *Struct_attributeContext)

	// ExitEnum is called when exiting the enum production.
	ExitEnum(c *EnumContext)

	// ExitEnum_body is called when exiting the enum_body production.
	ExitEnum_body(c *Enum_bodyContext)

	// ExitEnum_attribute is called when exiting the enum_attribute production.
	ExitEnum_attribute(c *Enum_attributeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComposed_value is called when exiting the composed_value production.
	ExitComposed_value(c *Composed_valueContext)

	// ExitList_value is called when exiting the list_value production.
	ExitList_value(c *List_valueContext)

	// ExitMap_value is called when exiting the map_value production.
	ExitMap_value(c *Map_valueContext)

	// ExitTuple_value is called when exiting the tuple_value production.
	ExitTuple_value(c *Tuple_valueContext)

	// ExitNamed_tuple_value is called when exiting the named_tuple_value production.
	ExitNamed_tuple_value(c *Named_tuple_valueContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitOr_expression is called when exiting the or_expression production.
	ExitOr_expression(c *Or_expressionContext)

	// ExitXor_expression is called when exiting the xor_expression production.
	ExitXor_expression(c *Xor_expressionContext)

	// ExitNot_expression is called when exiting the not_expression production.
	ExitNot_expression(c *Not_expressionContext)

	// ExitExpression_math is called when exiting the expression_math production.
	ExitExpression_math(c *Expression_mathContext)

	// ExitTerm_math is called when exiting the term_math production.
	ExitTerm_math(c *Term_mathContext)

	// ExitFactor_math is called when exiting the factor_math production.
	ExitFactor_math(c *Factor_mathContext)

	// ExitBitwise_math is called when exiting the bitwise_math production.
	ExitBitwise_math(c *Bitwise_mathContext)

	// ExitShift_math is called when exiting the shift_math production.
	ExitShift_math(c *Shift_mathContext)

	// ExitUnary_math is called when exiting the unary_math production.
	ExitUnary_math(c *Unary_mathContext)

	// ExitBefore_unary is called when exiting the before_unary production.
	ExitBefore_unary(c *Before_unaryContext)

	// ExitAfter_unary is called when exiting the after_unary production.
	ExitAfter_unary(c *After_unaryContext)

	// ExitMath_value is called when exiting the math_value production.
	ExitMath_value(c *Math_valueContext)

	// ExitParenthesis_expression is called when exiting the parenthesis_expression production.
	ExitParenthesis_expression(c *Parenthesis_expressionContext)

	// ExitBitwise_operator is called when exiting the bitwise_operator production.
	ExitBitwise_operator(c *Bitwise_operatorContext)

	// ExitShift_operator is called when exiting the shift_operator production.
	ExitShift_operator(c *Shift_operatorContext)

	// ExitTerm_operator is called when exiting the term_operator production.
	ExitTerm_operator(c *Term_operatorContext)

	// ExitFactor_operator is called when exiting the factor_operator production.
	ExitFactor_operator(c *Factor_operatorContext)

	// ExitComparision_operator is called when exiting the comparision_operator production.
	ExitComparision_operator(c *Comparision_operatorContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNew_scope is called when exiting the new_scope production.
	ExitNew_scope(c *New_scopeContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitFor_rule is called when exiting the for_rule production.
	ExitFor_rule(c *For_ruleContext)

	// ExitFor_in is called when exiting the for_in production.
	ExitFor_in(c *For_inContext)

	// ExitFull_for is called when exiting the full_for production.
	ExitFull_for(c *Full_forContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitDo_while_statement is called when exiting the do_while_statement production.
	ExitDo_while_statement(c *Do_while_statementContext)

	// ExitSwitch_statement is called when exiting the switch_statement production.
	ExitSwitch_statement(c *Switch_statementContext)

	// ExitSwitch_case is called when exiting the switch_case production.
	ExitSwitch_case(c *Switch_caseContext)

	// ExitTry_statement is called when exiting the try_statement production.
	ExitTry_statement(c *Try_statementContext)

	// ExitCatch_statement is called when exiting the catch_statement production.
	ExitCatch_statement(c *Catch_statementContext)

	// ExitFinally_statement is called when exiting the finally_statement production.
	ExitFinally_statement(c *Finally_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitElif_statement is called when exiting the elif_statement production.
	ExitElif_statement(c *Elif_statementContext)

	// ExitElse_statement is called when exiting the else_statement production.
	ExitElse_statement(c *Else_statementContext)

	// ExitStruct_func is called when exiting the struct_func production.
	ExitStruct_func(c *Struct_funcContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitAnonymous_func is called when exiting the anonymous_func production.
	ExitAnonymous_func(c *Anonymous_funcContext)

	// ExitFunc_param is called when exiting the func_param production.
	ExitFunc_param(c *Func_paramContext)

	// ExitFunc_callback is called when exiting the func_callback production.
	ExitFunc_callback(c *Func_callbackContext)

	// ExitFunc_param_callback is called when exiting the func_param_callback production.
	ExitFunc_param_callback(c *Func_param_callbackContext)

	// ExitFunc_param_rule is called when exiting the func_param_rule production.
	ExitFunc_param_rule(c *Func_param_ruleContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitFunc_call_params is called when exiting the func_call_params production.
	ExitFunc_call_params(c *Func_call_paramsContext)

	// ExitFunc_return is called when exiting the func_return production.
	ExitFunc_return(c *Func_returnContext)

	// ExitFunc_body is called when exiting the func_body production.
	ExitFunc_body(c *Func_bodyContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitVariable_declaration is called when exiting the variable_declaration production.
	ExitVariable_declaration(c *Variable_declarationContext)

	// ExitVariable_assign is called when exiting the variable_assign production.
	ExitVariable_assign(c *Variable_assignContext)

	// ExitVariable_name is called when exiting the variable_name production.
	ExitVariable_name(c *Variable_nameContext)

	// ExitAssing is called when exiting the assing production.
	ExitAssing(c *AssingContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitFinal_type is called when exiting the final_type production.
	ExitFinal_type(c *Final_typeContext)
}
