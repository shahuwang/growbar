package main

import (
	"fmt"
	"os"
)

type Interpreter struct {
	variable            *Variable
	function_list       *FunctionDefinition
	statement_list      *StatementList
	current_line_number int
	lex                 *GrowLex
}

func NewInterpreter() *Interpreter {
	interpreter = new(Interpreter)
	interpreter.current_line_number = 1
	return interpreter
}

func (ipt *Interpreter) Compile(fp *os.File) {
	ipt.lex.File = fp
	if GrowParse(ipt.lex) {
		fmt.Errorf("error")
		os.Exit(1)
	}
}

func (ipt *Interpreter) Interpret() {

}

func (ipt *Interpret) AddStdFp() {

}

func (ipt *Interpret) AddGlobalVariable(identifier string, value *Value) {

}
