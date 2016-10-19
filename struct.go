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
	Type        ExpressionType
	line_number int
}

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
	typ         StatementType
	line_number int
	u           *StatementTag
}

type StatementTag struct {
	value reflect.Value
	typ   reflect.Type
}

type StatementList struct {
	statement *Statement
	next      *StatementList
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
