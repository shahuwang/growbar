%{
package main
import(
)
%}
%union {
    identifier string
    expression *Expression
    statement *Statement
    statement_list *StatementList
    block *Block
    identifier_list *IdentifierList
    elsif *Elsif
    argument_list *ArgumentList
    parameter_list *ParameterList
}

%token <expression> INT_LITERAL
%token <expression> DOUBLE_LITERAL
%token <expression> STRING_LITERAL
%token <identifier> IDENTIFIER
%token FUNCTION IF ELSE ELSIF WHILE FOR RETURN_T BREAK CONTINUE NULL_T
        LP RP LC RC SEMICOLON COMMA ASSIGN LOGICAL_AND LOGICAL_OR EQ NE
        GT GE LT LE ADD SUB MUL DIV MOD TRUE_T FALSE_T GLOBAL_T
%type <expression> primary_expression unary_expression multiplicative_expression
                additive_expression relational_expression expression expression_opt
                logical_or_expression logical_and_expression  equality_expression

%type <statement> continue_statement return_statement break_statement statement
                for_statement global_statement if_statement while_statement
%type <statement_list> statement_list
%type <block> block
%type <identifier_list> identifier_list
%type <elsif> elsif elsif_list
%type <argument_list> argument_list
%type <parameter_list> parameter_list
%%
translation_unit
    : definition_or_statement
    | translation_unit definition_or_statement
    ;
definition_or_statement
    : function_definition
    | statement
    {
        ipt := getCurrentInterpreter()
        ipt.statement_list = chainStatementList(ipt.statement_list, $1)
    }
    ;
function_definition
    : FUNCTION IDENTIFIER LP parameter_list RP block
    {
        functionDefine($2, $4, $6)
    }
    | FUNCTION IDENTIFIER LP RP block
    {
        functionDefine($2, nil, $5);
    }
    ;
parameter_list
    : IDENTIFIER
    {
        $$ = createParameter($1)
    }
    | parameter_list COMMA IDENTIFIER
    {
        $$ = chainParameter($1, $3);
    }
    ;
argument_list
    : expression
    {
        $$ = createArgumentList($1)
    }
    | argument_list COMMA expression
    {
        $$ = chainArgumentList($1, $3)
    }
    ;
statement_list
    : statement
    {
        $$ = createStatementList($1);
    }
    | statement_list statement
    {
        $$ = chainStatementList($1, $2);
    }
    ;
expression
    : logical_or_expression
    | IDENTIFIER ASSIGN expression
    {
        $$ = createAssignExpression($1, $3)
    }
logical_or_expression
    : logical_and_expression
    | logical_or_expression LOGICAL_OR logical_and_expression
    {
        $$ = createBinaryExpression(LOGICAL_OR_EXPRESSION, $1, $3)
    }
    ;
logical_and_expression
    : equality_expression
    | logical_and_expression LOGICAL_AND equality_expression
    {
        $$ = createBinaryExpression(LOGICAL_AND_EXPRESSION, $1, $3)
    }
    ;
equality_expression
    : relational_expression
    | equality_expression EQ relational_expression
    {
        $$ = createBinaryExpression(EQ_EXPRESSION, $1, $3)
    }
    | equality_expression NE relational_expression
    {
        $$ = createBinaryExpression(NE_EXPRESSION, $1, $3)
    }
    ;
relational_expression
    : additive_expression
    | relational_expression GT additive_expression
    {
        $$ = createBinaryExpression(GT_EXPRESSION, $1, $3)
    }
    | relational_expression GE additive_expression
    {
        $$ = createBinaryExpression(GE_EXPRESSION, $1, $3)
    }
    | relational_expression LT additive_expression
    {
        $$ = createBinaryExpression(LT_EXPRESSION, $1, $3)
    }
    | relational_expression LE additive_expression
    {
        $$ = createBinaryExpression(LE_EXPRESSION, $1, $3)
    }
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
    : IDENTIFIER LP argument_list RP
    {
        $$ = createFunctionCallExpression($1, $3)
    }
    | IDENTIFIER LP RP
    {
        $$ = createFunctionCallExpression($1, nil)
    }
    | LP expression RP
    {
        $$ = $2
    }
    | IDENTIFIER
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

statement
    : expression SEMICOLON
    {
        $$ = createExpressionStatement($1);
    }
    | global_statement
    | if_statement
    | while_statement
    | for_statement
    | return_statement
    | break_statement
    | continue_statement
    ;
global_statement
    : GLOBAL_T identifier_list SEMICOLON
    {
        $$ = createGlobalStatement($2);
    }
identifier_list
    : IDENTIFIER
    {
        $$ = createGlobalIdentifier($1)
    }
    | identifier_list COMMA IDENTIFIER
    {
        $$ = chainIdentifier($1, $3)
    }
    ;
if_statement
    : IF LP expression RP block
    {
        $$ = createIfStatement($3, $5, nil, nil)
    }
    | IF LP expression RP block ELSE block
    {
        $$ = createIfStatement($3, $5, nil, $7)
    }
    | IF LP expression RP block elsif_list
    {
        $$ = createIfStatement($3, $5, $6, nil)   
    }
    | IF LP expression RP block elsif_list ELSE block
    {
        $$ = createIfStatement($3, $5, $6, $8)
    }
    ;
elsif_list
    : elsif
    | elsif_list elsif
    {
        $$ = chainElsifList($1, $2)
    }
    ;
elsif
    : ELSIF LP expression RP block
    {
        $$ = createElsif($3, $5)
    }
    ;
while_statement
    : WHILE LP expression RP block
    {
        $$ = createWhileStatement($3, $5)
    }
    ;
for_statement
    : FOR LP expression_opt SEMICOLON expression_opt SEMICOLON
      expression_opt RP block
    {
        $$ = createForStatement($3, $5, $7, $9);
    }
    ;
expression_opt
    :
    {
        $$ = nil;
    }
    | expression
    ;
return_statement
    : RETURN_T expression_opt SEMICOLON
    {
        $$ = createReturnStatement($2);
    }
    ;
break_statement
    : BREAK SEMICOLON
    {
        $$ = createBreakStatement();
    }
continue_statement
    : CONTINUE SEMICOLON
    {
        $$ = createContinueStatement();
    }
    ;
block
    : LC statement_list RC
    {
        $$ = createBlock($2)
    }
    | LC RC
    {
        $$ = createBlock(nil);
    }

%%
