package main

import ()

type Interpreter struct {
	variable            *Variable
	function_list       *FunctionDefinition
	statement_list      *StatementList
	current_line_number int
}

func NewInterpreter() *Interpreter {
	interpreter = new(Interpreter)
	interpreter.current_line_number = 1
	return interpreter
}
