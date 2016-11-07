package main

import ()

func allocExpression(typ ExpressionType) *Expression {
	exp := new(Expression)
	exp.Type = typ
	ipt := getCurrentInterpreter()
	exp.line_number = ipt.current_line_number
	return exp
}

func createAssignExpression(variable string, operand *Expression) *Expression {
	exp := allocExpression(ASSIGN_EXPRESSION)
	exp.assign_expression.variable = variable
	exp.assign_expression.operand = operand
	return exp
}

func createIdentifierExpression(identifier string) *Expression {
	exp := allocExpression(IDENTIFIER_EXPRESSION)
	exp.identifier = identifier
	return exp
}

func createBooleanExpression(value bool) *Expression {
	exp := allocExpression(BOOLEAN_EXPRESSION)
	exp.boolean_value = value
	return exp
}

func createNullExpression() *Expression {
	exp := allocExpression(NULL_EXPRESSION)
	return exp
}

func createMinusExpression(operand *Expression) *Expression {
	//TODO
	return operand
}

func createAddExpression(operand *Expression) *Expression {
	// TODO
	return operand
}

func createBinaryExpression(
	operator ExpressionType, left *Expression, right *Expression) *Expression {
	//TODO
	return left
}
