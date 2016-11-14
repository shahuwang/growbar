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
	v.u = p
	v.typ = CRB_NATIVE_POINTER_VALUE
	ipt.AddGlobalVariable("STDIN", v)
	p.pointer = os.Stdout
	v.u = p
	ipt.AddGlobalVariable("STDOUT", v)
	p.pointer = os.Stderr
	v.u = p
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

func (ipt *Interpreter) EvalExpression(env *LocalEnvironment, expr *Expression) Value {
	return ipt.EvalExpression(env, expr)
}

func (ipt *Interpreter) evalExpression(env *LocalEnvironment, expr *Expression) Value {
	//TODO
	var v Value
	return v
}
