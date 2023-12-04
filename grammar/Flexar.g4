grammar Flexar;

//! Lexer

// Structure
IMPORT
    : 'import'
    ;

NAMESPACE
    : 'namespace'
    ;

EXPORT
    : 'export'
    ;

USE
    : 'use'
    ;

AS
    : 'as'
    ;

BOOLEAN
    : 'true'
    | 'false'
    ;

THIS
    : 'this'
    ;

NEW
    : 'new'
    ;

// Enum

ENUM
    : 'enum'
    ;

STRUCT
    : 'struct'
    ;


// Function

FUNC
    : 'func'
    ;

// POO

CLASS
    : 'class'
    ;

INTERFACE
    : 'interface'
    ;

EXTENDS
    : 'extends'
    ;

IMPLEMENTS
    : 'implements'
    ;

PUBLIC
    : 'public'
    ;

PRIVATE
    : 'private'
    ;

PROTECTED
    : 'protected'
    ;

STATIC
    : 'static'
    ;

ABSTRACT
    : 'abstract'
    ;

FINAL
    : 'final'
    ;

OVERRIDE
    : 'override'
    ;

// Control

IF
    : 'if'
    ;

ELSE
    : 'else'
    ;

ELIF
    : 'elif'
    ;

SWITCH
    : 'switch'
    ;

CASE
    : 'case'
    ;

DEFAULT
    : 'default'
    ;

WHILE
    : 'while'
    ;

DO
    : 'do'
    ;

FOR
    : 'for'
    ;

BREAK
    : 'break'
    ;

CONTINUE
    : 'continue'
    ;

RETURN
    : 'return'
    ;

CONST
    : 'const'
    ;

READONLY
    : 'readonly'
    ;

IN
    : 'in'
    ;

IS
    : 'is'
    ;

// Exception

TRY
    : 'try'
    ;

CATCH
    : 'catch'
    ;

FINALLY
    : 'finally'
    ;

// Type

INT : 'int';
INT8 : 'int8';
INT16 : 'int16';
INT32 : 'int32';
INT64 : 'int64';
UINT : 'uint';
UINT8 : 'uint8';
UINT16 : 'uint16';
UINT32 : 'uint32';
UINT64 : 'uint64';
FLOAT : 'float';
FLOAT32 : 'float32';
FLOAT64 : 'float64';
BOOL : 'bool';
CHAR : 'char';
STRING : 'string';
INF : 'INF';
TUPLE : 'tuple';
MAP : 'map';
ERROR : 'error';
NULL : 'null';
ANY : 'any';
DYN : 'dyn';

// Text

TEMPLATE_STRING
    : '`' (ESC | ~[`\\] | '${' NAME '}')* '`'
    ;

DEFAULT_STRING
    : '"' (ESC | ~["\\])* '"'
    ;

fragment ESC
    : '\\' [btnr`"\\]
    ;

DISCARD
    : '_'
    ;

NAME
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

// Character

COMMA
    : ','
    ;

DOT
    : '.'
    ;

COLON
    : ':'
    ;

DOUBLE_COLON
    : '::'
    ;

ARROW
    : '->'
    ;

SEMICOLON
    : ';'
    ;

OPEN_PAREN
    : '('
    ;

CLOSE_PAREN
    : ')'
    ;

OPEN_BRACE
    : '{'
    ;

CLOSE_BRACE
    : '}'
    ;

OPEN_BRACKET
    : '['
    ;

CLOSE_BRACKET
    : ']'
    ;

// Operators



// Assign

ASSIGN
    : '='
    ;

PLUS_ASSIGN
    : '+='
    ;

MINUS_ASSIGN
    : '-='
    ;

STAR_ASSIGN
    : '*='
    ;

MODULE_ASSIGN
    : '%='
    ;

DIV_ASSIGN
    : '/='
    ;

EXP_ASSIGN
    : '**='
    ;

AND_ASSIGN
    : '&='
    ;

OR_ASSIGN
    : '|='
    ;

XOR_ASSIGN
    : '^='
    ;

SHL_ASSIGN
    : '<<='
    ;

SHR_ASSIGN
    : '>>='
    ;

INF_ASSIGN
    : ':='
    ;

// Compare

EQUAL
    : '=='
    ;

NOT_EQUAL
    : '!='
    ;

LESS
    : '<'
    ;

LESS_EQUAL
    : '<='
    ;

GREATER
    : '>'
    ;

GREATER_EQUAL
    : '>='
    ;

// Logic

AND
    : '&&'
    ;

OR
    : '||'
    ;

XOR
    : '^^'
    ;

NOT
    : '!'
    ;

// Bitwise

BIT_AND
    : '&'
    ;

BIT_OR
    : '|'
    ;

BIT_XOR
    : '^'
    ;

BIT_NOT
    : '~'
    ;

SHL
    : '<<'
    ;

SHR
    : '>>'
    ;

// Arithmetic

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

STAR
    : '*'
    ;

EXP
    : '**'
    ;

MODULE
    : '%'
    ;

DIV
    : '/'
    ;

INC
    : '++'
    ;

DEC
    : '--'
    ;

// Other

QUESTION
    : '?'
    ;

RANGE
    : '..'
    ;

SPREAD
    : '...'
    ;


// Number

INT_NUM
    : [-]? [0-9]+
    ;

FLOAT_NUM
    : [-]? [0-9]+ '.' [0-9]+ 'f'
    ;

COMMENT
    : '//' ~[\r\n]* -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;


//! Parser

program
    : import_group* namespace program_rule*
    ;

program_rule
    : EXPORT? (class
    | func
    | interface
    | struct
    | struct_func
    | enum)
    ;

// Import

import_group
    : IMPORT import_rule
    ;

import_rule
    : (USE? import_namespace (COMMA USE? import_namespace?)* SEMICOLON)+
    ;

import_namespace
    : namespace_name (DOUBLE_COLON NAME)? (AS NAME)?
    ;

namespace_name
    : NAME (DOT NAME)*
    ;

// Namespace

namespace
    : NAMESPACE namespace_name SEMICOLON
    ;

namespace_call
    : (namespace_name | NAME) DOUBLE_COLON (func_call | NAME)
    ;

// Class

class : ABSTRACT? CLASS NAME class_extends? class_implements? class_body
    ;

class_extends
    : EXTENDS (NAME | OPEN_PAREN NAME (COMMA NAME)* CLOSE_PAREN)
    ;

class_implements
    : IMPLEMENTS (NAME | OPEN_PAREN NAME (COMMA NAME)* CLOSE_PAREN)
    ;

class_body
    : OPEN_BRACE class_body_rule* CLOSE_BRACE
    ;

class_body_rule
    : class_attribute
    | constructor
    | class_method
    ;

class_attribute
    : class_modifier variable_declaration
    ;

privacy_modifier
    : PUBLIC | PRIVATE | PROTECTED
    ;

class_method
    : class_modifier func | abstract_method
    ;

abstract_method
    : ABSTRACT FUNC NAME OPEN_PAREN func_param? CLOSE_PAREN func_return?
    ;

class_modifier
    : (OVERRIDE | ABSTRACT)? READONLY? privacy_modifier? STATIC? FINAL?
    ;

constructor
    : privacy_modifier? NAME OPEN_PAREN func_param? CLOSE_PAREN func_body
    ;

class_new_instance
    : NEW (func_call | NAME OPEN_BRACE (NAME ASSIGN expression (COMMA NAME ASSIGN expression)*)? CLOSE_BRACE)
    ;

method_call
    : (THIS | variable_name) DOT func_call
    ;

attribute_call
    : (NAME | func_call) (DOT NAME | func_call)*
    ;


// Interface

interface
    : INTERFACE NAME interface_extends? interface_body
    ;

interface_extends
    : EXTENDS (NAME | OPEN_PAREN NAME (COMMA NAME)* CLOSE_PAREN)
    ;

interface_body
    : OPEN_BRACE interface_body_rule* CLOSE_BRACE
    ;

interface_body_rule
    : interface_method
    | interface_attribute
    ;

interface_method
    : FUNC NAME OPEN_PAREN func_param? CLOSE_PAREN func_return?
    ;

interface_attribute
    : NAME (COLON type)
    ;

// Struct

struct
    : STRUCT NAME struct_body
    ;

struct_body
    : OPEN_BRACE struct_attribute* CLOSE_BRACE
    ;

struct_attribute
    : variable_declaration
    ;

// Enum

enum
    : ENUM NAME enum_body
    ;

enum_body
    : OPEN_BRACE enum_attribute* CLOSE_BRACE
    ;

enum_attribute
    : NAME (ASSIGN INT_NUM)?
    ;

// Expression

expression
    : class_new_instance
    | func_call
    | anonymous_func
    | value
    | method_call
    | namespace_call
    | attribute_call
    | expression_math
    | composed_value
    | and_expression
    | expression IS type
    | expression comparision_operator expression
    | expression QUESTION expression COLON expression
    | expression RANGE expression
    | (type | NAME) OPEN_PAREN expression CLOSE_PAREN
    | NAME
    | NULL
    ;

composed_value
    : list_value | map_value | tuple_value | named_tuple_value
    ;

list_value
    : OPEN_BRACE (expression (COMMA expression)*)? CLOSE_BRACE
    ;

map_value
    : OPEN_BRACE (expression COLON expression (COMMA expression COLON expression)*)? CLOSE_BRACE
    ;

tuple_value
    : OPEN_PAREN (expression (COMMA expression)*)? CLOSE_PAREN
    ;

named_tuple_value
    : OPEN_PAREN (NAME COLON expression (COMMA NAME COLON expression)*)? CLOSE_PAREN
    ;

// Boolean

and_expression
    : and_expression AND or_expression
    | or_expression
    ;

or_expression
    : or_expression OR xor_expression
    | xor_expression
    ;

xor_expression
    : xor_expression XOR not_expression
    | not_expression
    ;

not_expression
    : NOT not_expression
    | BOOLEAN
    ;

// Math

// parenthesis_expression
expression_math
    : expression_math factor_operator term_math
    | term_math
    ;

term_math
    : term_math term_operator factor_math
    | factor_math
    ;

factor_math
    : factor_math EXP bitwise_math
    | bitwise_math
    ;

bitwise_math
    : bitwise_math bitwise_operator shift_math
    | shift_math
    ;

shift_math
    : shift_math shift_operator unary_math | unary_math
    ;

unary_math
    : before_unary | after_unary | math_value | parenthesis_expression
    ;

before_unary
    : INC math_value
    | DEC math_value
    | BIT_NOT math_value
    ;

after_unary
    : math_value INC
    | math_value DEC
    ;

math_value
    : variable_name | INT_NUM
    ;

parenthesis_expression
    : OPEN_PAREN expression_math CLOSE_PAREN
    ;

// Operator

bitwise_operator
    : BIT_AND
    | BIT_OR
    | BIT_XOR
    ;

shift_operator
    : SHL
    | SHR
    ;

term_operator
    : STAR
    | DIV
    | MODULE
    ;

factor_operator
    : PLUS
    | MINUS
    ;

comparision_operator
    : EQUAL
    | NOT_EQUAL
    | LESS
    | LESS_EQUAL
    | GREATER
    | GREATER_EQUAL
    ;

value
    : BOOLEAN | INT_NUM | FLOAT_NUM | DEFAULT_STRING | TEMPLATE_STRING
    ;

// Statement

statement
    : variable_declaration
    | variable_assign
    | expression
    | for_statement
    | if_statement
    | while_statement
    | do_while_statement
    | switch_statement
    | try_statement
    | new_scope
    | BREAK
    | CONTINUE
    ;

new_scope
    : OPEN_BRACE statement* CLOSE_BRACE
    ;

// For statement

for_statement
    : FOR OPEN_PAREN for_rule CLOSE_PAREN OPEN_BRACE statement* CLOSE_BRACE
    ;

for_rule
    : for_in
    | full_for
    ;

for_in
    :  (DISCARD | NAME) COMMA (DISCARD | NAME) IN expression
    ;

full_for
    : variable_declaration SEMICOLON expression? SEMICOLON expression?
    ;

// While statement

while_statement
    : WHILE OPEN_PAREN expression CLOSE_PAREN OPEN_BRACE statement* CLOSE_BRACE
    ;

do_while_statement
    : DO OPEN_BRACE statement* CLOSE_BRACE WHILE OPEN_PAREN expression CLOSE_PAREN
    ;

// Switch statement

switch_statement
    : SWITCH OPEN_PAREN expression CLOSE_PAREN OPEN_BRACE switch_case* CLOSE_BRACE
    ;

switch_case
    : CASE expression COLON statement*
    | DEFAULT COLON statement*
    ;

// Try statement

try_statement
    : TRY OPEN_BRACE statement* CLOSE_BRACE catch_statement? finally_statement?
    ;

catch_statement
    : CATCH OPEN_PAREN NAME COLON NAME CLOSE_PAREN OPEN_BRACE statement* CLOSE_BRACE
    ;

finally_statement
    : FINALLY OPEN_BRACE statement* CLOSE_BRACE
    ;


// If statement

if_statement
    : IF OPEN_PAREN expression CLOSE_PAREN OPEN_BRACE statement* CLOSE_BRACE elif_statement* else_statement?
    ;

elif_statement
    : ELIF OPEN_PAREN expression CLOSE_PAREN OPEN_BRACE statement* CLOSE_BRACE
    ;

else_statement
    : ELSE OPEN_BRACE statement* CLOSE_BRACE
    ;

// Function

struct_func
    : OPEN_PAREN NAME STAR? NAME CLOSE_PAREN func
    ;

func
    : FUNC NAME OPEN_PAREN func_param? CLOSE_PAREN func_return? func_body
    ;

anonymous_func
    : FUNC OPEN_PAREN func_param? CLOSE_PAREN func_return? ARROW func_body
    ;

func_param
    : func_param_rule (COMMA func_param_rule)*
    ;

func_callback
    : FUNC OPEN_PAREN func_param_callback? CLOSE_PAREN func_return?
    ;

func_param_callback
    : type (COMMA type)*
    ;

func_param_rule
    : NAME QUESTION? COLON type SPREAD? (ASSIGN value)?
    ;

func_call
    : NAME OPEN_PAREN func_call_params? CLOSE_PAREN
    ;

func_call_params
    : expression (COMMA expression)*
    ;

func_return
    : type | OPEN_PAREN type (COMMA type)* CLOSE_PAREN
    ;

func_body
    : OPEN_BRACE statement* return_statement? CLOSE_BRACE
    ;

return_statement
    : RETURN (
        expression (COMMA expression)*
        )
    ;

// Variable

variable_declaration
    : NAME (COLON type) (ASSIGN expression)?
    ;

variable_assign
    : variable_name assing expression
    ;

variable_name
    : NAME
    | THIS DOT NAME
    | DISCARD
    ;

assing
    : ASSIGN
    | PLUS_ASSIGN
    | MINUS_ASSIGN
    | STAR_ASSIGN
    | MODULE_ASSIGN
    | DIV_ASSIGN
    | EXP_ASSIGN
    | AND_ASSIGN
    | OR_ASSIGN
    | XOR_ASSIGN
    | SHL_ASSIGN
    | SHR_ASSIGN
    | INF_ASSIGN
    ;

// Type

type
    : (CONST? final_type (
        QUESTION
        | OPEN_BRACKET (INT_NUM | DYN | NAME)? CLOSE_BRACKET
        | OPEN_BRACKET final_type OPEN_PAREN (INT_NUM | DYN | NAME)? CLOSE_PAREN CLOSE_BRACKET
        )*)
    // callback
    | func_callback
    ;

final_type
    : INT | INT8 | INT16 | INT32 | INT64
          | UINT | UINT8 | UINT16 | UINT32 | UINT64
          | FLOAT | FLOAT32 | FLOAT64
          | BOOL
          | CHAR
          | STRING
          | INF
          | TUPLE
          | MAP
          | ERROR
          | NULL
          | ANY
          | NAME
    ;