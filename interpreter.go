package main

import (
	"fmt"
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

}

func (ipt *Interpreter) AddStdFp() {
	v := new(Value)
	p := NativePointer{}
	p.info = &NativePointerInfo{name: NATIVE_LIB_NAME}
	p.pointer = os.Stdin
	v.native_pointer = p
	v.typ = CRB_NATIVE_POINTER_VALUE
	ipt.AddGlobalVariable("STDIN", v)
	p.pointer = os.Stdout
	v.native_pointer = p
	v.native_pointer = p
	ipt.AddGlobalVariable("STDOUT", v)
	p.pointer = os.Stderr
	v.native_pointer = p
	ipt.AddGlobalVariable("STDERR", v)
}

func (ipt *Interpreter) AddGlobalVariable(identifier string, value *Value) {
	val := new(Variable)
	val.name = identifier
	val.value = *value
	val.next = ipt.variable
	ipt.variable.next = val
}

func (ipt *Interpreter) ExecuteStatementList(env *LocalEnvironment, stlist *StatementList) StatementResult {
	var pos *StatementList = stlist
	result := StatementResult{typ: NORMAL_STATEMENT_RESULT}
	for {
		if pos == nil {
			break
		}
		result = ipt.ExecuteStatement(env, pos.statement)
		if result.typ != NORMAL_STATEMENT_RESULT {
			return result
		}
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
	// v := ipt.EvalExpression(env, &statement.u.(Expression))
	// if v.typ == CRB_STRING_VALUE {
	// 	// crb_release_string(v.u.string_value)
	// 	// 这里不知道是否需要垃圾回收
	// }
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

func (ipt *Interpreter) evalExpression(env *LocalEnvironment, expr *Expression) Value {
	//TODO
	var v Value
	switch expr.Type {
	case INT_EXPRESSION:
		v = evalIntExpression(expr.int_value)
	case DOUBLE_EXPRESSION:
		v = evalDoubleExpression(expr.double_value)
	default:
		msg := fmt.Sprintf("bad case. type .. %d\n", expr.Type)
		panic(msg)
	}
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
		result.boolean_value = evalCompareString(operator, &lval, &rval, left.line_number)
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
	//TODO
}

func (ipt *Interpreter) evalBinaryDouble(
	operator ExpressionType, left float32, right float32, result *Value, line_number int) {
	//TODO
}

func (ipt *Interpreter) evalBinaryBoolean(operator ExpressionType, left bool, right bool, line_number int) bool {
	// TODO
	return true
}

func (ipt *Interpreter) evalBinaryNull(operator ExpressionType, lval *Value, rval *Value, line_number int) bool {
	//TODO
	return true
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
	// TODO
	return new(CRBString)
}
