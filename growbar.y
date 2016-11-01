%{
package main
import(
  "fmt"
)
%}
%union {
}

%token <expression> INT_LITERAL
%token <expression> DOUBLE_LITERAL
%token <expression> STRING_LITERAL
%token <identifier> IDENTIFIER
%token FUNCTION IF ELSE ELSIF WHILE FOR RETURN_T BREAK CONTINUE NULL_T
        LP RP LC RC SEMICOLON COMMA ASSIGN LOGICAL_AND LOGICAL_OR EQ NE
        GT GE LT LE ADD SUB MUL DIV MOD TRUE_T FALSE_T GLOBAL_T
%%
translation_unit
    : additive
    | translation_unit additive
    ;
additive
    : ADD
    {
        fmt.Println("is ADD")
    }
    | SUB
    {
        fmt.Println("is SUB")
    }
    | MUL
    {
        fmt.Println("is MUL")
    }
    | DIV
    {
        fmt.Println("is DIV")
    }
    | MOD
    {
        fmt.Println("is MOD")
    }
    | LOGICAL_AND
    {
        fmt.Println("is AND")
    }
    | LOGICAL_OR
    {
        fmt.Println("is OR")
    }
    | ASSIGN
    {
        fmt.Println("is ASSIGN")
    }
    | EQ
    {
        fmt.Println("is EQ")
    }
    | NE
    {
        fmt.Println("is NE")
    }
    | GT
    {
        fmt.Println("is GT")
    }
    | GE
    {
        fmt.Println("is GE")
    }
    | LT
    {
        fmt.Println("is LT")
    }
    | LE
    {
        fmt.Println("is LE")
    }
    | INT_LITERAL
    {
        fmt.Println("is INT_LITERAL")
    }
    | DOUBLE_LITERAL
    {
        fmt.Println("is DOUBLE_LITERAL")
    }
%%
