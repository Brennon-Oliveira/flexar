grammar flexar;

//! Lexer

// Structure
IMPORT
    : 'import'
    ;

NAMESPACE
    : 'namespace'
    ;



// Function

FUNC
    : 'func'
    ;

// POO

CLASS
    : 'class'
    ;




WS
    : [ \t\r\n]+ -> skip
    ;

// Parser

start
    : IMPORT
    ;
