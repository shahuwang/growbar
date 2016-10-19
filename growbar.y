%union {
    identifier String  
    parameter_list *ParameterList
    argument_list  *ArgumentList
    expression *Expression
    statement *Statement
    statement_list *StatementList
    block *Block
    elsif *Elsif
    identifier_list *IdentifierList
}

%token <expression> INT_LITERAL
%token <expression> DOUBLE_LITERAL
%token <expression> STRING_LITERAL
%token <identifier> IDENTIFIER
%token FUNCTION IF ELSE ELSIF WHILE FOR RETURN_T BREAK CONTINUE NULL_T
        LP RP LC RC SEMICOLON COMMA ASSIGN LOGICAL_AND LOGICAL_OR EQ NE
        GT GE LT LE ADD SUB MUL DIV MOD TRUE_T FALSE_T GLOBAL_T

%type <parameter_list> parameter_list
%type <argument_list> argument_list
%type <expression> expression expression_opt logical_and_expression
        logical_or_expression equality_expression relational_expression
        additive_expression multiplicative_expression unary_expression primary_expression

%type <statement> statement global_statement if_statement while_statement
        for_statement return_statement break_statement continue_statement

%type <statement_list> statement_list
%type <block> block
%type <elsif> elsif elsif_list
%type <identifier_list> identifier_list
