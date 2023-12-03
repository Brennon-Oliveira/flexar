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

// Enum

ENUM
    : 'enum'
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
INF : 'INF';
TUPLE : 'tuple';
MAP : 'map';
ERROR : 'error';
NULL : 'null';
ANY : 'any';

// Text

STRING
    : '"' (ESC | ~["\\])* '"'
    ;

fragment ESC
    : '\\' [btnr"\\]
    ;

NAME
    : [a-zA-Z][a-zA-Z0-9_]*
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

IS
    : 'is'
    ;

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

// Parser

start
    : IMPORT
    ;
