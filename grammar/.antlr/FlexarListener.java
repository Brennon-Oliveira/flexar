// Generated from /workspaces/script_ease/grammar/Flexar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link FlexarParser}.
 */
public interface FlexarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link FlexarParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(FlexarParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(FlexarParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#program_rule}.
	 * @param ctx the parse tree
	 */
	void enterProgram_rule(FlexarParser.Program_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#program_rule}.
	 * @param ctx the parse tree
	 */
	void exitProgram_rule(FlexarParser.Program_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#import_group}.
	 * @param ctx the parse tree
	 */
	void enterImport_group(FlexarParser.Import_groupContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#import_group}.
	 * @param ctx the parse tree
	 */
	void exitImport_group(FlexarParser.Import_groupContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#import_rule}.
	 * @param ctx the parse tree
	 */
	void enterImport_rule(FlexarParser.Import_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#import_rule}.
	 * @param ctx the parse tree
	 */
	void exitImport_rule(FlexarParser.Import_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#import_namespace}.
	 * @param ctx the parse tree
	 */
	void enterImport_namespace(FlexarParser.Import_namespaceContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#import_namespace}.
	 * @param ctx the parse tree
	 */
	void exitImport_namespace(FlexarParser.Import_namespaceContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#namespace_name}.
	 * @param ctx the parse tree
	 */
	void enterNamespace_name(FlexarParser.Namespace_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#namespace_name}.
	 * @param ctx the parse tree
	 */
	void exitNamespace_name(FlexarParser.Namespace_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#namespace}.
	 * @param ctx the parse tree
	 */
	void enterNamespace(FlexarParser.NamespaceContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#namespace}.
	 * @param ctx the parse tree
	 */
	void exitNamespace(FlexarParser.NamespaceContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#namespace_call}.
	 * @param ctx the parse tree
	 */
	void enterNamespace_call(FlexarParser.Namespace_callContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#namespace_call}.
	 * @param ctx the parse tree
	 */
	void exitNamespace_call(FlexarParser.Namespace_callContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class}.
	 * @param ctx the parse tree
	 */
	void enterClass(FlexarParser.ClassContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class}.
	 * @param ctx the parse tree
	 */
	void exitClass(FlexarParser.ClassContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_extends}.
	 * @param ctx the parse tree
	 */
	void enterClass_extends(FlexarParser.Class_extendsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_extends}.
	 * @param ctx the parse tree
	 */
	void exitClass_extends(FlexarParser.Class_extendsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_implements}.
	 * @param ctx the parse tree
	 */
	void enterClass_implements(FlexarParser.Class_implementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_implements}.
	 * @param ctx the parse tree
	 */
	void exitClass_implements(FlexarParser.Class_implementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_body}.
	 * @param ctx the parse tree
	 */
	void enterClass_body(FlexarParser.Class_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_body}.
	 * @param ctx the parse tree
	 */
	void exitClass_body(FlexarParser.Class_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_body_rule}.
	 * @param ctx the parse tree
	 */
	void enterClass_body_rule(FlexarParser.Class_body_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_body_rule}.
	 * @param ctx the parse tree
	 */
	void exitClass_body_rule(FlexarParser.Class_body_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_attribute}.
	 * @param ctx the parse tree
	 */
	void enterClass_attribute(FlexarParser.Class_attributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_attribute}.
	 * @param ctx the parse tree
	 */
	void exitClass_attribute(FlexarParser.Class_attributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#privacy_modifier}.
	 * @param ctx the parse tree
	 */
	void enterPrivacy_modifier(FlexarParser.Privacy_modifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#privacy_modifier}.
	 * @param ctx the parse tree
	 */
	void exitPrivacy_modifier(FlexarParser.Privacy_modifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_method}.
	 * @param ctx the parse tree
	 */
	void enterClass_method(FlexarParser.Class_methodContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_method}.
	 * @param ctx the parse tree
	 */
	void exitClass_method(FlexarParser.Class_methodContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#abstract_method}.
	 * @param ctx the parse tree
	 */
	void enterAbstract_method(FlexarParser.Abstract_methodContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#abstract_method}.
	 * @param ctx the parse tree
	 */
	void exitAbstract_method(FlexarParser.Abstract_methodContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_modifier}.
	 * @param ctx the parse tree
	 */
	void enterClass_modifier(FlexarParser.Class_modifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_modifier}.
	 * @param ctx the parse tree
	 */
	void exitClass_modifier(FlexarParser.Class_modifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#constructor}.
	 * @param ctx the parse tree
	 */
	void enterConstructor(FlexarParser.ConstructorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#constructor}.
	 * @param ctx the parse tree
	 */
	void exitConstructor(FlexarParser.ConstructorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#class_new_instance}.
	 * @param ctx the parse tree
	 */
	void enterClass_new_instance(FlexarParser.Class_new_instanceContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#class_new_instance}.
	 * @param ctx the parse tree
	 */
	void exitClass_new_instance(FlexarParser.Class_new_instanceContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#method_call}.
	 * @param ctx the parse tree
	 */
	void enterMethod_call(FlexarParser.Method_callContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#method_call}.
	 * @param ctx the parse tree
	 */
	void exitMethod_call(FlexarParser.Method_callContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#attribute_call}.
	 * @param ctx the parse tree
	 */
	void enterAttribute_call(FlexarParser.Attribute_callContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#attribute_call}.
	 * @param ctx the parse tree
	 */
	void exitAttribute_call(FlexarParser.Attribute_callContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface}.
	 * @param ctx the parse tree
	 */
	void enterInterface(FlexarParser.InterfaceContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface}.
	 * @param ctx the parse tree
	 */
	void exitInterface(FlexarParser.InterfaceContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface_extends}.
	 * @param ctx the parse tree
	 */
	void enterInterface_extends(FlexarParser.Interface_extendsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface_extends}.
	 * @param ctx the parse tree
	 */
	void exitInterface_extends(FlexarParser.Interface_extendsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface_body}.
	 * @param ctx the parse tree
	 */
	void enterInterface_body(FlexarParser.Interface_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface_body}.
	 * @param ctx the parse tree
	 */
	void exitInterface_body(FlexarParser.Interface_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface_body_rule}.
	 * @param ctx the parse tree
	 */
	void enterInterface_body_rule(FlexarParser.Interface_body_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface_body_rule}.
	 * @param ctx the parse tree
	 */
	void exitInterface_body_rule(FlexarParser.Interface_body_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface_method}.
	 * @param ctx the parse tree
	 */
	void enterInterface_method(FlexarParser.Interface_methodContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface_method}.
	 * @param ctx the parse tree
	 */
	void exitInterface_method(FlexarParser.Interface_methodContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#interface_attribute}.
	 * @param ctx the parse tree
	 */
	void enterInterface_attribute(FlexarParser.Interface_attributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#interface_attribute}.
	 * @param ctx the parse tree
	 */
	void exitInterface_attribute(FlexarParser.Interface_attributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#struct}.
	 * @param ctx the parse tree
	 */
	void enterStruct(FlexarParser.StructContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#struct}.
	 * @param ctx the parse tree
	 */
	void exitStruct(FlexarParser.StructContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#struct_body}.
	 * @param ctx the parse tree
	 */
	void enterStruct_body(FlexarParser.Struct_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#struct_body}.
	 * @param ctx the parse tree
	 */
	void exitStruct_body(FlexarParser.Struct_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#struct_attribute}.
	 * @param ctx the parse tree
	 */
	void enterStruct_attribute(FlexarParser.Struct_attributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#struct_attribute}.
	 * @param ctx the parse tree
	 */
	void exitStruct_attribute(FlexarParser.Struct_attributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#enum}.
	 * @param ctx the parse tree
	 */
	void enterEnum(FlexarParser.EnumContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#enum}.
	 * @param ctx the parse tree
	 */
	void exitEnum(FlexarParser.EnumContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#enum_body}.
	 * @param ctx the parse tree
	 */
	void enterEnum_body(FlexarParser.Enum_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#enum_body}.
	 * @param ctx the parse tree
	 */
	void exitEnum_body(FlexarParser.Enum_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#enum_attribute}.
	 * @param ctx the parse tree
	 */
	void enterEnum_attribute(FlexarParser.Enum_attributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#enum_attribute}.
	 * @param ctx the parse tree
	 */
	void exitEnum_attribute(FlexarParser.Enum_attributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(FlexarParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(FlexarParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#composed_value}.
	 * @param ctx the parse tree
	 */
	void enterComposed_value(FlexarParser.Composed_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#composed_value}.
	 * @param ctx the parse tree
	 */
	void exitComposed_value(FlexarParser.Composed_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#list_value}.
	 * @param ctx the parse tree
	 */
	void enterList_value(FlexarParser.List_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#list_value}.
	 * @param ctx the parse tree
	 */
	void exitList_value(FlexarParser.List_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#map_value}.
	 * @param ctx the parse tree
	 */
	void enterMap_value(FlexarParser.Map_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#map_value}.
	 * @param ctx the parse tree
	 */
	void exitMap_value(FlexarParser.Map_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#tuple_value}.
	 * @param ctx the parse tree
	 */
	void enterTuple_value(FlexarParser.Tuple_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#tuple_value}.
	 * @param ctx the parse tree
	 */
	void exitTuple_value(FlexarParser.Tuple_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#named_tuple_value}.
	 * @param ctx the parse tree
	 */
	void enterNamed_tuple_value(FlexarParser.Named_tuple_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#named_tuple_value}.
	 * @param ctx the parse tree
	 */
	void exitNamed_tuple_value(FlexarParser.Named_tuple_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#and_expression}.
	 * @param ctx the parse tree
	 */
	void enterAnd_expression(FlexarParser.And_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#and_expression}.
	 * @param ctx the parse tree
	 */
	void exitAnd_expression(FlexarParser.And_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#or_expression}.
	 * @param ctx the parse tree
	 */
	void enterOr_expression(FlexarParser.Or_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#or_expression}.
	 * @param ctx the parse tree
	 */
	void exitOr_expression(FlexarParser.Or_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#xor_expression}.
	 * @param ctx the parse tree
	 */
	void enterXor_expression(FlexarParser.Xor_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#xor_expression}.
	 * @param ctx the parse tree
	 */
	void exitXor_expression(FlexarParser.Xor_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#not_expression}.
	 * @param ctx the parse tree
	 */
	void enterNot_expression(FlexarParser.Not_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#not_expression}.
	 * @param ctx the parse tree
	 */
	void exitNot_expression(FlexarParser.Not_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#expression_math}.
	 * @param ctx the parse tree
	 */
	void enterExpression_math(FlexarParser.Expression_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#expression_math}.
	 * @param ctx the parse tree
	 */
	void exitExpression_math(FlexarParser.Expression_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#term_math}.
	 * @param ctx the parse tree
	 */
	void enterTerm_math(FlexarParser.Term_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#term_math}.
	 * @param ctx the parse tree
	 */
	void exitTerm_math(FlexarParser.Term_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#factor_math}.
	 * @param ctx the parse tree
	 */
	void enterFactor_math(FlexarParser.Factor_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#factor_math}.
	 * @param ctx the parse tree
	 */
	void exitFactor_math(FlexarParser.Factor_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#bitwise_math}.
	 * @param ctx the parse tree
	 */
	void enterBitwise_math(FlexarParser.Bitwise_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#bitwise_math}.
	 * @param ctx the parse tree
	 */
	void exitBitwise_math(FlexarParser.Bitwise_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#shift_math}.
	 * @param ctx the parse tree
	 */
	void enterShift_math(FlexarParser.Shift_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#shift_math}.
	 * @param ctx the parse tree
	 */
	void exitShift_math(FlexarParser.Shift_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#unary_math}.
	 * @param ctx the parse tree
	 */
	void enterUnary_math(FlexarParser.Unary_mathContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#unary_math}.
	 * @param ctx the parse tree
	 */
	void exitUnary_math(FlexarParser.Unary_mathContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#before_unary}.
	 * @param ctx the parse tree
	 */
	void enterBefore_unary(FlexarParser.Before_unaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#before_unary}.
	 * @param ctx the parse tree
	 */
	void exitBefore_unary(FlexarParser.Before_unaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#after_unary}.
	 * @param ctx the parse tree
	 */
	void enterAfter_unary(FlexarParser.After_unaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#after_unary}.
	 * @param ctx the parse tree
	 */
	void exitAfter_unary(FlexarParser.After_unaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#math_value}.
	 * @param ctx the parse tree
	 */
	void enterMath_value(FlexarParser.Math_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#math_value}.
	 * @param ctx the parse tree
	 */
	void exitMath_value(FlexarParser.Math_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#parenthesis_expression}.
	 * @param ctx the parse tree
	 */
	void enterParenthesis_expression(FlexarParser.Parenthesis_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#parenthesis_expression}.
	 * @param ctx the parse tree
	 */
	void exitParenthesis_expression(FlexarParser.Parenthesis_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#bitwise_operator}.
	 * @param ctx the parse tree
	 */
	void enterBitwise_operator(FlexarParser.Bitwise_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#bitwise_operator}.
	 * @param ctx the parse tree
	 */
	void exitBitwise_operator(FlexarParser.Bitwise_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#shift_operator}.
	 * @param ctx the parse tree
	 */
	void enterShift_operator(FlexarParser.Shift_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#shift_operator}.
	 * @param ctx the parse tree
	 */
	void exitShift_operator(FlexarParser.Shift_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#term_operator}.
	 * @param ctx the parse tree
	 */
	void enterTerm_operator(FlexarParser.Term_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#term_operator}.
	 * @param ctx the parse tree
	 */
	void exitTerm_operator(FlexarParser.Term_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#factor_operator}.
	 * @param ctx the parse tree
	 */
	void enterFactor_operator(FlexarParser.Factor_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#factor_operator}.
	 * @param ctx the parse tree
	 */
	void exitFactor_operator(FlexarParser.Factor_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#comparision_operator}.
	 * @param ctx the parse tree
	 */
	void enterComparision_operator(FlexarParser.Comparision_operatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#comparision_operator}.
	 * @param ctx the parse tree
	 */
	void exitComparision_operator(FlexarParser.Comparision_operatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(FlexarParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(FlexarParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(FlexarParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(FlexarParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#new_scope}.
	 * @param ctx the parse tree
	 */
	void enterNew_scope(FlexarParser.New_scopeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#new_scope}.
	 * @param ctx the parse tree
	 */
	void exitNew_scope(FlexarParser.New_scopeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#for_statement}.
	 * @param ctx the parse tree
	 */
	void enterFor_statement(FlexarParser.For_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#for_statement}.
	 * @param ctx the parse tree
	 */
	void exitFor_statement(FlexarParser.For_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#for_rule}.
	 * @param ctx the parse tree
	 */
	void enterFor_rule(FlexarParser.For_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#for_rule}.
	 * @param ctx the parse tree
	 */
	void exitFor_rule(FlexarParser.For_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#for_in}.
	 * @param ctx the parse tree
	 */
	void enterFor_in(FlexarParser.For_inContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#for_in}.
	 * @param ctx the parse tree
	 */
	void exitFor_in(FlexarParser.For_inContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#full_for}.
	 * @param ctx the parse tree
	 */
	void enterFull_for(FlexarParser.Full_forContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#full_for}.
	 * @param ctx the parse tree
	 */
	void exitFull_for(FlexarParser.Full_forContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#while_statement}.
	 * @param ctx the parse tree
	 */
	void enterWhile_statement(FlexarParser.While_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#while_statement}.
	 * @param ctx the parse tree
	 */
	void exitWhile_statement(FlexarParser.While_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#do_while_statement}.
	 * @param ctx the parse tree
	 */
	void enterDo_while_statement(FlexarParser.Do_while_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#do_while_statement}.
	 * @param ctx the parse tree
	 */
	void exitDo_while_statement(FlexarParser.Do_while_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#switch_statement}.
	 * @param ctx the parse tree
	 */
	void enterSwitch_statement(FlexarParser.Switch_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#switch_statement}.
	 * @param ctx the parse tree
	 */
	void exitSwitch_statement(FlexarParser.Switch_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#switch_case}.
	 * @param ctx the parse tree
	 */
	void enterSwitch_case(FlexarParser.Switch_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#switch_case}.
	 * @param ctx the parse tree
	 */
	void exitSwitch_case(FlexarParser.Switch_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#try_statement}.
	 * @param ctx the parse tree
	 */
	void enterTry_statement(FlexarParser.Try_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#try_statement}.
	 * @param ctx the parse tree
	 */
	void exitTry_statement(FlexarParser.Try_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#catch_statement}.
	 * @param ctx the parse tree
	 */
	void enterCatch_statement(FlexarParser.Catch_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#catch_statement}.
	 * @param ctx the parse tree
	 */
	void exitCatch_statement(FlexarParser.Catch_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#finally_statement}.
	 * @param ctx the parse tree
	 */
	void enterFinally_statement(FlexarParser.Finally_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#finally_statement}.
	 * @param ctx the parse tree
	 */
	void exitFinally_statement(FlexarParser.Finally_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#if_statement}.
	 * @param ctx the parse tree
	 */
	void enterIf_statement(FlexarParser.If_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#if_statement}.
	 * @param ctx the parse tree
	 */
	void exitIf_statement(FlexarParser.If_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#elif_statement}.
	 * @param ctx the parse tree
	 */
	void enterElif_statement(FlexarParser.Elif_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#elif_statement}.
	 * @param ctx the parse tree
	 */
	void exitElif_statement(FlexarParser.Elif_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#else_statement}.
	 * @param ctx the parse tree
	 */
	void enterElse_statement(FlexarParser.Else_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#else_statement}.
	 * @param ctx the parse tree
	 */
	void exitElse_statement(FlexarParser.Else_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#struct_func}.
	 * @param ctx the parse tree
	 */
	void enterStruct_func(FlexarParser.Struct_funcContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#struct_func}.
	 * @param ctx the parse tree
	 */
	void exitStruct_func(FlexarParser.Struct_funcContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func}.
	 * @param ctx the parse tree
	 */
	void enterFunc(FlexarParser.FuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func}.
	 * @param ctx the parse tree
	 */
	void exitFunc(FlexarParser.FuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#anonymous_func}.
	 * @param ctx the parse tree
	 */
	void enterAnonymous_func(FlexarParser.Anonymous_funcContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#anonymous_func}.
	 * @param ctx the parse tree
	 */
	void exitAnonymous_func(FlexarParser.Anonymous_funcContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_param}.
	 * @param ctx the parse tree
	 */
	void enterFunc_param(FlexarParser.Func_paramContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_param}.
	 * @param ctx the parse tree
	 */
	void exitFunc_param(FlexarParser.Func_paramContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_callback}.
	 * @param ctx the parse tree
	 */
	void enterFunc_callback(FlexarParser.Func_callbackContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_callback}.
	 * @param ctx the parse tree
	 */
	void exitFunc_callback(FlexarParser.Func_callbackContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_param_callback}.
	 * @param ctx the parse tree
	 */
	void enterFunc_param_callback(FlexarParser.Func_param_callbackContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_param_callback}.
	 * @param ctx the parse tree
	 */
	void exitFunc_param_callback(FlexarParser.Func_param_callbackContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_param_rule}.
	 * @param ctx the parse tree
	 */
	void enterFunc_param_rule(FlexarParser.Func_param_ruleContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_param_rule}.
	 * @param ctx the parse tree
	 */
	void exitFunc_param_rule(FlexarParser.Func_param_ruleContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_call}.
	 * @param ctx the parse tree
	 */
	void enterFunc_call(FlexarParser.Func_callContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_call}.
	 * @param ctx the parse tree
	 */
	void exitFunc_call(FlexarParser.Func_callContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_call_params}.
	 * @param ctx the parse tree
	 */
	void enterFunc_call_params(FlexarParser.Func_call_paramsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_call_params}.
	 * @param ctx the parse tree
	 */
	void exitFunc_call_params(FlexarParser.Func_call_paramsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_return}.
	 * @param ctx the parse tree
	 */
	void enterFunc_return(FlexarParser.Func_returnContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_return}.
	 * @param ctx the parse tree
	 */
	void exitFunc_return(FlexarParser.Func_returnContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#func_body}.
	 * @param ctx the parse tree
	 */
	void enterFunc_body(FlexarParser.Func_bodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#func_body}.
	 * @param ctx the parse tree
	 */
	void exitFunc_body(FlexarParser.Func_bodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#return_statement}.
	 * @param ctx the parse tree
	 */
	void enterReturn_statement(FlexarParser.Return_statementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#return_statement}.
	 * @param ctx the parse tree
	 */
	void exitReturn_statement(FlexarParser.Return_statementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#variable_declaration}.
	 * @param ctx the parse tree
	 */
	void enterVariable_declaration(FlexarParser.Variable_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#variable_declaration}.
	 * @param ctx the parse tree
	 */
	void exitVariable_declaration(FlexarParser.Variable_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#variable_assign}.
	 * @param ctx the parse tree
	 */
	void enterVariable_assign(FlexarParser.Variable_assignContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#variable_assign}.
	 * @param ctx the parse tree
	 */
	void exitVariable_assign(FlexarParser.Variable_assignContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#variable_name}.
	 * @param ctx the parse tree
	 */
	void enterVariable_name(FlexarParser.Variable_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#variable_name}.
	 * @param ctx the parse tree
	 */
	void exitVariable_name(FlexarParser.Variable_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#assing}.
	 * @param ctx the parse tree
	 */
	void enterAssing(FlexarParser.AssingContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#assing}.
	 * @param ctx the parse tree
	 */
	void exitAssing(FlexarParser.AssingContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(FlexarParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(FlexarParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlexarParser#final_type}.
	 * @param ctx the parse tree
	 */
	void enterFinal_type(FlexarParser.Final_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlexarParser#final_type}.
	 * @param ctx the parse tree
	 */
	void exitFinal_type(FlexarParser.Final_typeContext ctx);
}