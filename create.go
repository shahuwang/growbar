package main

import ()

func allocExpression(typ ExpressionType) *Expression {
	exp := new(Expression)
	exp.Type = typ
	ipt := getCurrentInterpreter()
	exp.line_number = ipt.current_line_number
	return exp
}
