package main

import (
	// "fmt"
	"os"
)

const NATIVE_LIB_NAME = "grobar.lang.file"

type Interpreter struct {
	variable            *Variable
	function_list       *FunctionDefinition
	statement_list      *StatementList
	current_line_number int
	lex                 *GrowLex
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

	}
FUNC_END:
	return result
}
