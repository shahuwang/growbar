package main

import (
	"reflect"
)

type ExpressionType int

const (
	BOOLEAN_EXPRESSION ExpressionType = iota + 1
	INT_EXPRESSION
	DOUBLE_EXPRESSION
	STRING_EXPRESSION
	IDENTIFIER_EXPRESSION
	ASSIGN_EXPRESSION
	ADD_EXPRESSION
	SUB_EXPRESSION
	MUL_EXPRESSION
	DIV_EXPRESSION
	MOD_EXPRESSION
	EQ_EXPRESSION
	NE_EXPRESSION
	GT_EXPRESSION
	GE_EXPRESSION
	LT_EXPRESSION
	LE_EXPRESSION
	LOGICAL_AND_EXPRESSION
	LOGICAL_OR_EXPRESSION
	MINUS_EXPRESSION
	FUNCTION_CALL_EXPRESSION
	NULL_EXPRESSION
	EXPRESSION_TYPE_COUNT_PLUS_1
)

type Expression struct {
	Type              ExpressionType
	line_number       int
	boolean_value     bool
	int_value         int
	double_value      float32
	string_value      string
	identifier        string
	assign_expression *AssignExpression
	minus_expression  *Expression
	binary_expression *BinaryExpression
}

type BinaryExpression struct {
	left  *Expression
	right *Expression
}

type AssignExpression struct {
	variable string
	operand  *Expression
}

type CompileError int

const (
	PARSE_ERR = iota + 1
	CHARACTER_INVALID_ERR
	FUNCTION_MULTIPLE_DEFINE_ERR
	COMPILE_ERROR_COUNT_PLUS_1
)

type ParameterList struct {
	name string
	next *ParameterList
	u    *Parameter
}

type Parameter struct {
	value reflect.Value
	typ   reflect.Type
}

type ArgumentList struct {
	expression *Expression
	next       *ArgumentList
}

type StatementType int

const (
	EXPRESSION_STATEMENT StatementType = iota + 1
	GLOBAL_STATEMENT
	IF_STATEMENT
	WHILE_STATEMENT
	FOR_STATEMENT
	RETURN_STATEMENT
	BREAK_STATEMENT
	CONTINUE_STATEMENT
	STATEMENT_TYPE_COUNT_PLUS_1
)

type Statement struct {
	typ             StatementType
	line_number     int
	expresion_s     *Expression
	identifier_list *IdentifierList
}

type StatementList struct {
	statement *Statement
	next      *StatementList
}

type RuntimeError int

const (
	VARIABLE_NOT_FOUND_ERROR RuntimeError = iota + 1
	FUNCTION_NOT__FOUNT_ERR
	ARGUMENT_TOO_MANY_ERR
	ARGUMENT_TOO_FEW_ERR
	NOT_BOOLEAN_TYPE_ERR
	MINUS_OPERAND_TYPE_ERR
	BAD_OPERAND_TYPE_ERR
	NOT_BOOLEAN_OPERATOR_ERR
	FOPEN_ARGUMENT_TYPE_ERR
	FCLOSE_ARGUMENT_TYPE_ERR
	FGETS_ARGUMENT_TYPE_ERR
	FPUTS_ARGUMENT_TYPE_ERR
	NOT_NULL_OPERATOR_ERR
	DIVISION_BY_ZERO_ERR
	GLOBAL_VARIABLE_NOT_FOUND_ERR
	GLOBAL_STATEMENT_IN_TOPLEVEL_ERR
	BAD_OPERATOR_FOR_STRING_ERR
	RUNTIME_ERROR_COUNT_PLUS_1
)

type StatementResultType int

const (
	NORMAL_STATEMENT_RESULT StatementResultType = iota + 1
	RETURN_STATEMENT_RESULT
	BREAK_STATEMENT_RESULT
	CONTINUE_STATEMENT_RESULT
	STATEMENT_RESULT_TYPE_COUNT_PLUS_1
)

type StatementResult struct {
	typ StatementResultType
	u   interface{}
}

type Block struct {
	statement_list *StatementList
}

type Elsif struct {
	condition *Expression
	block     *Block
	next      *Elsif
}

type IdentifierList struct {
	name string
	next *IdentifierList
}

type GlobalStatement struct {
	identifier_list *IdentifierList
}

type Variable struct {
	name  string
	value Value
	next  *Variable
}

type CRBString struct {
	ref_count  int
	str        string
	is_literal bool
}

type GlobalaVariableRef struct {
	variable *Variable
	next     *GlobalaVariableRef
}

type LocalEnvironment struct {
	variable        *Variable
	global_variable GlobalaVariableRef
}

type FunctionDefinitionType int

const (
	CROWBAR_FUNCTION_DEFINITION FunctionDefinitionType = iota + 1
	NATIVE_FUNCTION_DEFINITION
)

type FunctionDefinition struct {
	name string
	typ  FunctionDefinitionType
	u    interface{}
}

type MessageFormat struct {
	format string
}

type CrowbarFunc struct {
	parameter *ParameterList
	block     *Block
}

type NativeFunc struct {
	proc *NativeFuncProc
}

type NativeFuncProc func(interpreter *Interpreter, arg_count int, args *Value) Value
