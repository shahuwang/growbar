package main

import (
	"fmt"
	"math"
	"os"
)

const NATIVE_LIB_NAME = "grobar.lang.file"

type Interpreter struct {
	variable            *Variable
	function_list       *FunctionDefinition
	statement_list      *StatementList
	current_line_number int
	lex                 *Growlex
}

var INTER *Interpreter = NewInterpreter()

func getCurrentInterpreter() *Interpreter {
	return INTER
}

func NewInterpreter() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.current_line_number = 1
	return interpreter
}

func (ipt *Interpreter) Compile(fp *os.File) {
	ipt.lex.File = fp
	// if GrowParse(ipt.lex) {
	// 	fmt.Errorf("error")
	// 	os.Exit(1)
	// }
}

func (ipt *Interpreter) Interpret() {
	ipt.AddStdFp()
	ipt.ExecuteStatementList(nil, ipt.statement_list)
}

func (ipt *Interpreter) AddStdFp() {
	v := newFp(os.Stdin)
	ipt.AddGlobalVariable("STDIN", v)
	v = newFp(os.Stdout)
	ipt.AddGlobalVariable("STDOUT", v)
	v = newFp(os.Stderr)
	ipt.AddGlobalVariable("STDERR", v)
}

func newFp(f *os.File) *Value {
	v := new(Value)
	p := NativePointer{}
	p.info = &NativePointerInfo{name: NATIVE_LIB_NAME}
	p.pointer = f
	v.native_pointer = p
	v.typ = CRB_NATIVE_POINTER_VALUE
	return v
}

func (ipt *Interpreter) AddGlobalVariable(identifier string, value *Value) {
	val := new(Variable)
	val.name = identifier
	val.value = *value
	val.next = ipt.variable
	ipt.variable = val
}

func (ipt *Interpreter) ExecuteStatementList(env *LocalEnvironment, stlist *StatementList) StatementResult {
	var pos *StatementList = stlist
	result := StatementResult{typ: NORMAL_STATEMENT_RESULT}
	for {
		if pos == nil {
			break
		}
		result = ipt.ExecuteStatement(env, pos.statement)
		PrintResult(&result)
		if result.typ != NORMAL_STATEMENT_RESULT {
			return result
		}
		pos = pos.next
	}
	return result
}

func (ipt *Interpreter) ExecuteStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	switch statement.typ {
	case EXPRESSION_STATEMENT:
		return ipt.ExecuteExpressionStatement(env, statement)
	case GLOBAL_STATEMENT:
		return ipt.ExecuteGlobalStatement(env, statement)
	case IF_STATEMENT:
		return ipt.ExecuteIfStatement(env, statement)
	case WHILE_STATEMENT:
		return ipt.ExecuteWhileStatement(env, statement)
	case FOR_STATEMENT:
		return ipt.ExecuteForStatement(env, statement)
	case RETURN_STATEMENT:
		return ipt.ExecuteReturnStatement(env, statement)
	case BREAK_STATEMENT:
		return ipt.ExecuteBreakStatement(env, statement)
	case CONTINUE_STATEMENT:
		return ipt.ExecuteContinueStatement(env, statement)
	case STATEMENT_TYPE_COUNT_PLUS_1:
		fallthrough
	default:
		errmsg, err := fmt.Printf("bad case ...%d", statement.typ)
		if err != nil {
			panic(err)
		}
		panic(errmsg)
	}
}

func (ipt *Interpreter) ExecuteExpressionStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	result := StatementResult{typ: NORMAL_STATEMENT_RESULT}
	v := ipt.evalExpression(env, statement.expresion_s)
	if v.typ == CRB_STRING_VALUE {
		releaseString(v.string_value)
	}
	result.return_value = v
	return result
}

func (ipt *Interpreter) ExecuteGlobalStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	var pos *IdentifierList
	result := StatementResult{typ: NORMAL_STATEMENT_RESULT}
	if env == nil {
		runtimeError(statement.line_number, GLOBAL_STATEMENT_IN_TOPLEVEL_ERR)
	}
	pos = statement.identifier_list
	for {
		if pos == nil {
			break
		}

	}
	return result
}

func (ipt *Interpreter) ExecuteIfStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	return StatementResult{}
}

func (ipt *Interpreter) ExecuteWhileStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	return StatementResult{}
}

func (ipt *Interpreter) ExecuteForStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	result := StatementResult{typ: NORMAL_STATEMENT_RESULT}
	// var cond Value
	return result
}

func (ipt *Interpreter) ExecuteReturnStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	return StatementResult{typ: RETURN_STATEMENT_RESULT}
}

func (ipt *Interpreter) ExecuteContinueStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	return StatementResult{typ: CONTINUE_STATEMENT_RESULT}
}

func (ipt *Interpreter) ExecuteBreakStatement(env *LocalEnvironment, statement *Statement) StatementResult {
	return StatementResult{typ: BREAK_STATEMENT_RESULT}
}

// func (ipt *Interpreter) EvalExpression(env *LocalEnvironment, expr *Expression) Value {
// 	return ipt.EvalExpression(env, expr)
// }

func (ipt *Interpreter) evalStringExpression(str string) Value {
	var v Value
	v.typ = CRB_STRING_VALUE
	v.string_value = ipt.literalToCrbString(str)
	return v
}

func (ipt *Interpreter) literalToCrbString(str string) *CRBString {
	ret := new(CRBString)
	ret.str = str
	ret.is_literal = true
	ret.ref_count = 1
	return ret
}

func (ipt *Interpreter) evalExpression(env *LocalEnvironment, expr *Expression) Value {
	var v Value
	switch expr.Type {
	case INT_EXPRESSION:
		v = evalIntExpression(expr.int_value)
	case DOUBLE_EXPRESSION:
		v = evalDoubleExpression(expr.double_value)
	case BOOLEAN_EXPRESSION:
		v = evalBooleanExpression(expr.boolean_value)
	case STRING_EXPRESSION:
		v = ipt.evalStringExpression(expr.string_value)
	case IDENTIFIER_EXPRESSION:
		v = ipt.evalIdentifierExpression(env, expr)
	case ASSIGN_EXPRESSION:
		v = ipt.evalAssignExpression(env, expr.assign_expression.variable, expr.assign_expression.operand)
	case ADD_EXPRESSION, SUB_EXPRESSION, MUL_EXPRESSION,
		DIV_EXPRESSION, MOD_EXPRESSION, EQ_EXPRESSION,
		NE_EXPRESSION, GT_EXPRESSION, GE_EXPRESSION,
		LT_EXPRESSION, LE_EXPRESSION:
		v = ipt.evalBinaryExpression(env, expr.Type, expr.binary_expression.left, expr.binary_expression.right)
	case LOGICAL_OR_EXPRESSION, LOGICAL_AND_EXPRESSION:
		v = ipt.evalLogicalAndOrExpression(env, expr.Type, expr.binary_expression.left, expr.binary_expression.right)
	case MINUS_EXPRESSION:
		v = ipt.evalMinusExpression(env, expr.minus_expression)
	case FUNCTION_CALL_EXPRESSION:
		v = ipt.evalFunctionCallExpression(env, expr)
	case NULL_EXPRESSION:
		v = evalNullExpression()
	default:
		msg := fmt.Sprintf("bad case. type .. %d\n", expr.Type)
		panic(msg)
	}
	return v
}

func (ipt *Interpreter) evalFunctionCallExpression(env *LocalEnvironment, expr *Expression) Value {
	var value Value
	var fn *FunctionDefinition
	var identifier string = expr.function_call_expression.identifier
	fn = searchFunction(identifier)
	if fn == nil {
		runtimeError(expr.line_number, FUNCTION_NOT__FOUNT_ERR)
	}
	switch fn.typ {
	case CROWBAR_FUNCTION_DEFINITION:
		value = ipt.callGrowbarFunction(env, expr, fn)
	case NATIVE_FUNCTION_DEFINITION:
		value = ipt.callNativeFunction(env, expr, fn.native_f.proc)
	default:
		msg := fmt.Sprintf("bad case ..%d\n", fn.typ)
		panic(msg)
	}
	return value
}

func (ipt *Interpreter) callNativeFunction(env *LocalEnvironment, expr *Expression, proc NativeFuncProc) Value {
	var value Value
	var arg_count int = 0
	var arg_p *ArgumentList = expr.function_call_expression.argument
	var args []Value = make([]Value, 0)
	var i int = 0
	for {
		if arg_p == nil {
			break
		}
		arg_count++
		arg_p = arg_p.next
	}
	arg_p = expr.function_call_expression.argument
	for {
		if arg_p == nil {
			break
		}
		args[i] = ipt.evalExpression(env, arg_p.expression)
		i++
		arg_p = arg_p.next
	}
	value = proc(ipt, arg_count, args)
	i = 0
	for {
		if i >= arg_count {
			break
		}
		releaseIfString(&args[i])
		i++
	}
	//TODO
	// MEM_free(args)
	return value
}

func (ipt *Interpreter) callGrowbarFunction(env *LocalEnvironment, expr *Expression, fn *FunctionDefinition) Value {
	var value Value
	var result StatementResult
	var arg_p *ArgumentList = expr.function_call_expression.argument
	var param_p *ParameterList = fn.growbar_f.parameter
	var local_env *LocalEnvironment = allocLocalEnvironment()
	for {
		if arg_p == nil {
			break
		}
		if param_p == nil {
			runtimeError(expr.line_number, ARGUMENT_TOO_MANY_ERR)
		}
		var arg_val Value = ipt.evalExpression(env, arg_p.expression)
		addLocalVariable(local_env, param_p.name, &arg_val)
		arg_p = arg_p.next
		param_p = param_p.next
	}
	if param_p != nil {
		runtimeError(expr.line_number, ARGUMENT_TOO_FEW_ERR)
	}
	result = ipt.ExecuteStatementList(local_env, fn.growbar_f.block.statement_list)
	if result.typ == RETURN_STATEMENT_RESULT {
		value = result.return_value
	} else {
		value.typ = CRB_NULL_VALUE
	}
	ipt.disposeLocalEnvironment(env)
	return value
}

func (ipt *Interpreter) disposeLocalEnvironment(env *LocalEnvironment) {
	//TODO
	//释放内存
}

func (ipt *Interpreter) evalLogicalAndOrExpression(env *LocalEnvironment, operator ExpressionType, left *Expression, right *Expression) Value {
	var lval Value
	var rval Value
	var result Value
	result.typ = CRB_BOOLEAN_VALUE
	lval = ipt.evalExpression(env, left)
	if lval.typ != CRB_BOOLEAN_VALUE {
		runtimeError(left.line_number, NOT_BOOLEAN_TYPE_ERR)
	}
	if operator == LOGICAL_AND_EXPRESSION {
		if !lval.boolean_value {
			result.boolean_value = false
			return result
		}
	} else if operator == LOGICAL_OR_EXPRESSION {
		if lval.boolean_value {
			result.boolean_value = true
			return result
		}
	} else {
		msg := fmt.Sprintf("bad operator.. %d\n", operator)
		panic(msg)
	}
	rval = ipt.evalExpression(env, right)
	if rval.typ != CRB_BOOLEAN_VALUE {
		runtimeError(left.line_number, NOT_BOOLEAN_TYPE_ERR)
	}
	result.boolean_value = rval.boolean_value
	return result
}

func (ipt *Interpreter) evalAssignExpression(env *LocalEnvironment, identifier string, expr *Expression) Value {
	var v Value = ipt.evalExpression(env, expr)
	var left *Variable = searchLocalVariable(env, identifier)
	if left == nil {
		left = ipt.searchGlobalVariableFromEnv(env, identifier)
	}
	if left != nil {
		releaseIfString(&left.value)
		left.value = v
		referIfString(&v)
	} else {
		if env != nil {
			addLocalVariable(env, identifier, &v)
		} else {
			ipt.AddGlobalVariable(identifier, &v)
		}
		referIfString(&v)
	}
	return v
}

func (ipt *Interpreter) searchGlobalVariableFromEnv(env *LocalEnvironment, identifier string) *Variable {
	pos := new(GlobalaVariableRef)
	if env == nil {
		return ipt.searchGlobalVariable(identifier)
	}
	pos = env.global_variable
	for {
		if pos == nil {
			break
		}
		if pos.variable.name == identifier {
			return pos.variable
		}
		pos = pos.next
	}
	return nil
}

func (ipt *Interpreter) searchGlobalVariable(identifier string) *Variable {
	pos := ipt.variable
	for {
		if pos == nil || pos.name == identifier {
			break
		}
	}
	return pos
}

func (ipt *Interpreter) evalIdentifierExpression(env *LocalEnvironment, expr *Expression) Value {
	vp := searchLocalVariable(env, expr.identifier)
	var v Value
	if vp != nil {
		v = vp.value
	} else {
		vp = ipt.searchGlobalVariableFromEnv(env, expr.identifier)
		if vp != nil {
			v = vp.value
		} else {
			runtimeError(expr.line_number, VARIABLE_NOT_FOUND_ERROR)
		}
	}
	referIfString(&v)
	return v
}

func (ipt *Interpreter) evalMinusExpression(env *LocalEnvironment, operand *Expression) Value {
	result := new(Value)
	val := ipt.evalExpression(env, operand)
	if val.typ == CRB_INT_VALUE {
		result.typ = CRB_INT_VALUE
		result.int_value = -operand.int_value
	} else if val.typ == CRB_DOUBLE_VALUE {
		result.typ = CRB_DOUBLE_VALUE
		result.double_value = -operand.double_value
	} else {
		runtimeError(operand.line_number, MINUS_OPERAND_TYPE_ERR)
	}
	return *result
}

func (ipt *Interpreter) evalAddExpression(env *LocalEnvironment, operand *Expression) Value {
	result := new(Value)
	val := ipt.evalExpression(env, operand)
	if val.typ == CRB_INT_VALUE {
		result.typ = CRB_INT_VALUE
		result.int_value = operand.int_value
	} else if val.typ == CRB_DOUBLE_VALUE {
		result.typ = CRB_DOUBLE_VALUE
		result.double_value = operand.double_value
	} else {
		runtimeError(operand.line_number, MINUS_OPERAND_TYPE_ERR)
	}
	return *result

}

func (ipt *Interpreter) evalBinaryExpression(
	env *LocalEnvironment,
	operator ExpressionType,
	left *Expression, right *Expression) Value {
	lval := ipt.evalExpression(env, left)
	rval := ipt.evalExpression(env, right)
	var result Value
	if lval.typ == CRB_INT_VALUE && rval.typ == CRB_INT_VALUE {
		ipt.evalBinaryInt(operator, lval.int_value, rval.int_value, &result, left.line_number)
	} else if lval.typ == CRB_DOUBLE_VALUE && rval.typ == CRB_DOUBLE_VALUE {
		ipt.evalBinaryDouble(operator, lval.double_value, rval.double_value, &result, left.line_number)
	} else if lval.typ == CRB_INT_VALUE && rval.typ == CRB_DOUBLE_VALUE {
		lval.double_value = float32(lval.int_value)
		ipt.evalBinaryDouble(operator, lval.double_value, rval.double_value, &result, left.line_number)
	} else if lval.typ == CRB_DOUBLE_VALUE && rval.typ == CRB_INT_VALUE {
		rval.double_value = float32(rval.int_value)
		ipt.evalBinaryDouble(operator, lval.double_value, rval.double_value, &result, left.line_number)
	} else if lval.typ == CRB_BOOLEAN_VALUE && rval.typ == CRB_BOOLEAN_VALUE {
		result.typ = CRB_BOOLEAN_VALUE
		result.boolean_value = ipt.evalBinaryBoolean(operator, lval.boolean_value, rval.boolean_value, left.line_number)
	} else if lval.typ == CRB_STRING_VALUE && operator == ADD_EXPRESSION {
		var right_str *CRBString
		if rval.typ == CRB_INT_VALUE {
			str := fmt.Sprintf("%d", rval.int_value)
			right_str = ipt.createCrowbarString(str)
		} else if rval.typ == CRB_DOUBLE_VALUE {
			str := fmt.Sprintf("%f", rval.double_value)
			right_str = ipt.createCrowbarString(str)
		} else if rval.typ == CRB_BOOLEAN_VALUE {
			str := fmt.Sprintf("%t", rval.boolean_value)
			right_str = ipt.createCrowbarString(str)
		} else if rval.typ == CRB_STRING_VALUE {
			right_str = rval.string_value
		} else if rval.typ == CRB_NATIVE_POINTER_VALUE {
			str := fmt.Sprintf("(%s:%p)", rval.native_pointer.info.name, rval.native_pointer.pointer)
			right_str = ipt.createCrowbarString(str)
		} else if rval.typ == CRB_NULL_VALUE {
			right_str = ipt.createCrowbarString("null")
		}
		result.typ = CRB_STRING_VALUE
		result.string_value = ipt.chainString(lval.string_value, right_str)
	} else if lval.typ == CRB_STRING_VALUE && rval.typ == CRB_STRING_VALUE {
		result.typ = CRB_BOOLEAN_VALUE
		result.boolean_value = ipt.evalCompareString(operator, &lval, &rval, left.line_number)
	} else if lval.typ == CRB_NULL_VALUE || rval.typ == CRB_NULL_VALUE {
		result.typ = CRB_BOOLEAN_VALUE
		result.boolean_value = ipt.evalBinaryNull(operator, &lval, &rval, left.line_number)
	} else {
		opt_str := getOperatorString(operator)
		msg := fmt.Sprintf(" operator: %s", opt_str)
		runtimeError(left.line_number, BAD_OPERAND_TYPE_ERR, msg)
	}
	return result
}

func (ipt *Interpreter) evalCompareString(operator ExpressionType, left *Value, right *Value, line_number int) bool {
	lstr := left.string_value.str
	rstr := right.string_value.str
	switch operator {
	case EQ_EXPRESSION:
		return lstr == rstr
	case NE_EXPRESSION:
		return lstr != rstr
	case GT_EXPRESSION:
		return lstr > rstr
	case GE_EXPRESSION:
		return lstr >= rstr
	case LT_EXPRESSION:
		return lstr < rstr
	case LE_EXPRESSION:
		return lstr <= rstr
	default:
		opt_str := getOperatorString(operator)
		msg := fmt.Sprintf(" operator: %s", opt_str)
		runtimeError(line_number, BAD_OPERATOR_FOR_STRING_ERR, msg)
	}
	return true
}

func (ipt *Interpreter) evalBinaryInt(
	operator ExpressionType,
	left int, right int,
	result *Value, line_number int) {
	switch operator {
	case ADD_EXPRESSION:
		result.int_value = left + right
		result.typ = CRB_INT_VALUE
	case SUB_EXPRESSION:
		result.int_value = left - right
		result.typ = CRB_INT_VALUE
	case MUL_EXPRESSION:
		result.int_value = left * right
		result.typ = CRB_INT_VALUE
	case DIV_EXPRESSION:
		result.int_value = left / right
		result.typ = CRB_INT_VALUE
	case MOD_EXPRESSION:
		result.int_value = left % right
		result.typ = CRB_INT_VALUE
	case EQ_EXPRESSION:
		result.boolean_value = (left == right)
		result.typ = CRB_BOOLEAN_VALUE
	case NE_EXPRESSION:
		result.boolean_value = (left != right)
		result.typ = CRB_BOOLEAN_VALUE
	case GT_EXPRESSION:
		result.boolean_value = (left > right)
		result.typ = CRB_BOOLEAN_VALUE
	case GE_EXPRESSION:
		result.boolean_value = (left >= right)
		result.typ = CRB_BOOLEAN_VALUE
	case LT_EXPRESSION:
		result.boolean_value = (left < right)
		result.typ = CRB_BOOLEAN_VALUE
	case LE_EXPRESSION:
		result.boolean_value = (left <= right)
		result.typ = CRB_BOOLEAN_VALUE
	default:
		msg := fmt.Sprintf("bad case ... %d", operator)
		panic(msg)
	}
}

func (ipt *Interpreter) evalBinaryDouble(
	operator ExpressionType, left float32, right float32, result *Value, line_number int) {
	switch operator {
	case ADD_EXPRESSION:
		result.double_value = left + right
		result.typ = CRB_DOUBLE_VALUE
	case SUB_EXPRESSION:
		result.double_value = left - right
		result.typ = CRB_DOUBLE_VALUE
	case MUL_EXPRESSION:
		result.double_value = left * right
		result.typ = CRB_DOUBLE_VALUE
	case DIV_EXPRESSION:
		result.double_value = left / right
		result.typ = CRB_DOUBLE_VALUE
	case MOD_EXPRESSION:
		result.double_value = float32(math.Mod(float64(left), float64(right)))
		result.typ = CRB_DOUBLE_VALUE
	case EQ_EXPRESSION:
		result.boolean_value = (left == right)
		result.typ = CRB_BOOLEAN_VALUE
	case NE_EXPRESSION:
		result.boolean_value = (left != right)
		result.typ = CRB_BOOLEAN_VALUE
	case GT_EXPRESSION:
		result.boolean_value = (left > right)
		result.typ = CRB_BOOLEAN_VALUE
	case GE_EXPRESSION:
		result.boolean_value = (left >= right)
		result.typ = CRB_BOOLEAN_VALUE
	case LT_EXPRESSION:
		result.boolean_value = (left < right)
		result.typ = CRB_BOOLEAN_VALUE
	case LE_EXPRESSION:
		result.boolean_value = (left <= right)
		result.typ = CRB_BOOLEAN_VALUE
	default:
		msg := fmt.Sprintf("bad case ... %d", operator)
		panic(msg)
	}
}

func (ipt *Interpreter) evalBinaryBoolean(operator ExpressionType, left bool, right bool, line_number int) bool {
	var result bool
	if operator == EQ_EXPRESSION {
		result = (left == right)
	} else if operator == NE_EXPRESSION {
		result = (left != right)
	} else {
		op_str := getOperatorString(operator)
		runtimeError(line_number, NOT_BOOLEAN_OPERATOR_ERR, op_str)
	}
	return result
}

func (ipt *Interpreter) evalBinaryNull(operator ExpressionType, lval *Value, rval *Value, line_number int) bool {
	var result bool
	if operator == EQ_EXPRESSION {
		result = (lval.typ == CRB_NULL_VALUE && rval.typ == CRB_NULL_VALUE)
	} else if operator == NE_EXPRESSION {
		result = !(lval.typ == CRB_NULL_VALUE && rval.typ == CRB_NULL_VALUE)
	} else {
		op_str := getOperatorString(operator)
		runtimeError(line_number, NOT_NULL_OPERATOR_ERR, op_str)
	}
	return result
}

func (ipt *Interpreter) createCrowbarString(str string) *CRBString {
	ret := ipt.allocCrbString(str, false)
	ret.ref_count = 1
	return ret
}

func (ipt *Interpreter) allocCrbString(str string, is_literal bool) *CRBString {
	ret := new(CRBString)
	ret.ref_count = 0
	ret.is_literal = is_literal
	ret.str = str
	return ret
}

func (ipt *Interpreter) chainString(left *CRBString, right *CRBString) *CRBString {
	ret := ipt.createCrowbarString(left.str + right.str)
	releaseString(left)
	releaseString(right)
	return ret
}
