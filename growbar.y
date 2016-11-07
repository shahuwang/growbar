%{
package main
import(
 /* "fmt" */
)
%}
%union {
    identifier string
    expression *Expression
}

%token <expression> INT_LITERAL
%token <expression> DOUBLE_LITERAL
%token <expression> STRING_LITERAL
%token <identifier> IDENTIFIER
%token FUNCTION IF ELSE ELSIF WHILE FOR RETURN_T BREAK CONTINUE NULL_T
        LP RP LC RC SEMICOLON COMMA ASSIGN LOGICAL_AND LOGICAL_OR EQ NE
        GT GE LT LE ADD SUB MUL DIV MOD TRUE_T FALSE_T GLOBAL_T
%type <expression> primary_expression unary_expression multiplicative_expression
                additive_expression
%%
translation_unit
    : additive_expression
    | translation_unit additive_expression
    ;
additive_expression
    : multiplicative_expression
    | additive_expression ADD multiplicative_expression
    {
        $$ = createBinaryExpression(ADD_EXPRESSION, $1, $3)
    }
    | additive_expression SUB multiplicative_expression
    {
        $$ = createBinaryExpression(SUB_EXPRESSION, $1, $3)
    }
    ;
multiplicative_expression
    : unary_expression
    | multiplicative_expression MUL unary_expression
    {
       $$ = createBinaryExpression(MUL_EXPRESSION, $1, $3) 
    }
    | multiplicative_expression DIV unary_expression
    {
       $$ = createBinaryExpression(DIV_EXPRESSION, $1, $3) 
    }
    | multiplicative_expression MOD unary_expression
    {
       $$ = createBinaryExpression(MOD_EXPRESSION, $1, $3) 
    }
    ;

unary_expression
    : primary_expression
    | SUB unary_expression
    {
        $$ = createMinusExpression($2)  
    }
    | ADD unary_expression
    {
        $$ = createAddExpression($2) 
    }
    ;
primary_expression
    : IDENTIFIER
    {
        $$ = createIdentifierExpression($1)  
    }
    | INT_LITERAL
    | DOUBLE_LITERAL
    | STRING_LITERAL
    | TRUE_T
    {
        $$ = createBooleanExpression(true)
    }
    | FALSE_T
    {
        $$ = createBooleanExpression(false)
    }
    | NULL_T
    {
        $$ = createNullExpression();
    }
    ;
%%
