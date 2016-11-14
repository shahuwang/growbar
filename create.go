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

func allocStatement(typ StatementType) *Statement {
	st := new(Statement)
	st.typ = typ
	st.line_number = getCurrentInterpreter().current_line_number
	return st
}

func createContinueStatement() *Statement {
	//TODO
	return allocStatement(CONTINUE_STATEMENT)
}

func createBreakStatement() *Statement {
	// TODO
	return allocStatement(BREAK_STATEMENT)
}

func createBlock(st *StatementList) *Block {
	block := new(Block)
	block.statement_list = st
	return block
}

func createExpressionStatement(exp *Expression) *Statement {
	st := allocStatement(EXPRESSION_STATEMENT)
	st.expresion_s = exp
	return st
}

func createStatementList(st *Statement) *StatementList {
	sl := new(StatementList)
	sl.statement = st
	return sl
}

func chainStatementList(sl *StatementList, st *Statement) *StatementList {
	if sl == nil {
		return createStatementList(st)
	}
	pos := sl
	for {
		if pos.next == nil {
			break
		}
		pos = pos.next
	}
	return sl
}

func createForStatement(init *Expression, cond *Expression, post *Expression, block *Block) *Statement {
	// TODO
	return allocStatement(FOR_STATEMENT)
}

func createReturnStatement(expression *Expression) *Statement {
	// TODO
	return allocStatement(RETURN_STATEMENT)
}

func createGlobalStatement(identifier_list *IdentifierList) *Statement {
	// TODO
	return allocStatement(GLOBAL_STATEMENT)
}

func createIfStatement(cond *Expression, then *Block, elsiflist *Elsif, elseBlock *Block) *Statement {
	// TODO
	return allocStatement(IF_STATEMENT)
}

func createWhileStatement(cond *Expression, block *Block) *Statement {
	//TODO
	return allocStatement(WHILE_STATEMENT)
}

func createGlobalIdentifier(identifier string) *IdentifierList {
	// TODO
	return new(IdentifierList)
}

func chainIdentifier(il *IdentifierList, identifier string) *IdentifierList {
	// TODO
	return new(IdentifierList)
}

func chainElsifList(list *Elsif, add *Elsif) *Elsif {
	// TODO
	return new(Elsif)
}

func createElsif(expr *Expression, block *Block) *Elsif {
	// TODO
	return new(Elsif)
}

func createArgumentList(expr *Expression) *ArgumentList {
	// TODO
	return new(ArgumentList)
}

func chainArgumentList(list *ArgumentList, expr *Expression) *ArgumentList {
	//TODO
	return new(ArgumentList)
}

func createParameter(identifier string) *ParameterList {
	// TODO
	return new(ParameterList)
}

func chainParameter(list *ParameterList, identifier string) *ParameterList {
	// TODO
	return new(ParameterList)
}

func createFunctionCallExpression(funcName string, argument *ArgumentList) *Expression {
	// TODO
	return new(Expression)
}

func functionDefine(identifier string, pl *ParameterList, block *Block) {
	// TODO
}
