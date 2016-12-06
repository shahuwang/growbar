package main

import (
// "fmt"
)

func allocExpression(typ ExpressionType) *Expression {
	exp := new(Expression)
	exp.Type = typ
	ipt := getCurrentInterpreter()
	exp.line_number = ipt.current_line_number
	exp.binary_expression = new(BinaryExpression)
	exp.minus_expression = new(Expression)
	exp.minus_expression.Type = MINUS_EXPRESSION
	exp.assign_expression = new(AssignExpression)
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

func convertValueToExpression(v *Value) Expression {
	expr := new(Expression)
	ipt := getCurrentInterpreter()
	expr.line_number = ipt.current_line_number
	if v.typ == CRB_INT_VALUE {
		expr.Type = INT_EXPRESSION
		expr.int_value = v.int_value
	} else if v.typ == CRB_DOUBLE_VALUE {
		expr.Type = DOUBLE_EXPRESSION
		expr.double_value = v.double_value
	} else {
		expr.Type = BOOLEAN_EXPRESSION
		expr.boolean_value = v.boolean_value
	}
	return *expr
}

func createMinusExpression(operand *Expression) *Expression {
	if operand.Type == INT_EXPRESSION || operand.Type == DOUBLE_EXPRESSION {
		ipt := getCurrentInterpreter()
		v := ipt.evalMinusExpression(nil, operand)
		*operand = convertValueToExpression(&v)
		return operand
	} else {
		exp := allocExpression(MINUS_EXPRESSION)
		exp.minus_expression = operand
		return exp
	}
}

func createAddExpression(operand *Expression) *Expression {
	if operand.Type == INT_EXPRESSION || operand.Type == DOUBLE_EXPRESSION {
		ipt := getCurrentInterpreter()
		v := ipt.evalAddExpression(nil, operand)
		*operand = convertValueToExpression(&v)
		return operand
	} else {
		exp := allocExpression(ADD_EXPRESSION)
		exp.minus_expression = operand
		return exp
	}
	return operand
}

func createBinaryExpression(
	operator ExpressionType, left *Expression, right *Expression) *Expression {
	if (left.Type == INT_EXPRESSION ||
		left.Type == DOUBLE_EXPRESSION) &&
		(right.Type == INT_EXPRESSION ||
			right.Type == DOUBLE_EXPRESSION) {
		ipt := getCurrentInterpreter()
		v := ipt.evalBinaryExpression(nil, operator, left, right)
		*left = convertValueToExpression(&v)
		return left
	} else {
		exp := allocExpression(operator)
		exp.binary_expression.left = left
		exp.binary_expression.right = right
		return exp
	}
}

func allocStatement(typ StatementType) *Statement {
	st := new(Statement)
	st.typ = typ
	st.line_number = getCurrentInterpreter().current_line_number
	return st
}

func createContinueStatement() *Statement {
	return allocStatement(CONTINUE_STATEMENT)
}

func createBreakStatement() *Statement {
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
	pos.next = createStatementList(st)
	return sl
}

func createForStatement(init *Expression, cond *Expression, post *Expression, block *Block) *Statement {
	st := allocStatement(FOR_STATEMENT)
	st.for_s.init = init
	st.for_s.condition = cond
	st.for_s.post = post
	st.for_s.block = block
	return st
}

func createReturnStatement(expression *Expression) *Statement {
	st := allocStatement(RETURN_STATEMENT)
	st.return_s.return_value = expression
	return st
}

func createGlobalStatement(identifier_list *IdentifierList) *Statement {
	st := allocStatement(GLOBAL_STATEMENT)
	st.global_s.identifier_list = identifier_list
	return st
}

func createIfStatement(cond *Expression, then *Block, elsiflist *Elsif, elseBlock *Block) *Statement {
	st := allocStatement(IF_STATEMENT)
	st.if_s.condition = cond
	st.if_s.then_block = then
	st.if_s.elsif_list = elsiflist
	st.if_s.else_block = elseBlock
	return st
}

func createWhileStatement(cond *Expression, block *Block) *Statement {
	st := allocStatement(WHILE_STATEMENT)
	st.while_s.condition = cond
	st.while_s.block = block
	return st
}

func createGlobalIdentifier(identifier string) *IdentifierList {
	i_list := new(IdentifierList)
	i_list.name = identifier
	i_list.next = nil
	return i_list
}

func chainIdentifier(il *IdentifierList, identifier string) *IdentifierList {
	pos := il
	for {
		if pos == nil {
			break
		}
		pos = pos.next
	}
	pos.next = createGlobalIdentifier(identifier)
	return il
}

func chainElsifList(list *Elsif, add *Elsif) *Elsif {
	pos := list
	for {
		if pos.next == nil {
			break
		}
		pos = pos.next
	}
	pos.next = add
	return list
}

func createElsif(expr *Expression, block *Block) *Elsif {
	ei := new(Elsif)
	ei.condition = expr
	ei.block = block
	ei.next = nil
	return ei
}

func createArgumentList(expr *Expression) *ArgumentList {
	al := new(ArgumentList)
	al.expression = expr
	al.next = nil
	return al
}

func chainArgumentList(list *ArgumentList, expr *Expression) *ArgumentList {
	pos := list
	for {
		if pos.next == nil {
			break
		}
		pos = pos.next
	}
	pos.next = createArgumentList(expr)
	return list
}

func createParameter(identifier string) *ParameterList {
	p := new(ParameterList)
	p.name = identifier
	p.next = nil
	return p
}

func chainParameter(list *ParameterList, identifier string) *ParameterList {
	pos := list
	for {
		if pos.next == nil {
			break
		}
		pos = pos.next
	}
	pos.next = createParameter(identifier)
	return list
}

func createFunctionCallExpression(funcName string, argument *ArgumentList) *Expression {
	exp := allocExpression(FUNCTION_CALL_EXPRESSION)
	exp.function_call_expression.identifier = funcName
	exp.function_call_expression.argument = argument
	return exp
}

func functionDefine(identifier string, pl *ParameterList, block *Block) {
	ipt := getCurrentInterpreter()
	if searchFunction(identifier) != nil {
		compileError(ipt.current_line_number, FUNCTION_MULTIPLE_DEFINE_ERR, identifier)
	}
	f := new(FunctionDefinition)
	f.name = identifier
	f.typ = CROWBAR_FUNCTION_DEFINITION
	f.growbar_f = new(GrowbarFunction)
	f.growbar_f.parameter = pl
	f.growbar_f.block = block
	f.next = ipt.function_list
	ipt.function_list = f
}
