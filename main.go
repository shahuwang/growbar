package main

import (
	// "fmt"
	"os"
)

func main() {
	file, err := os.Open("test.gb")
	if err != nil {
		panic(err)
	}
	g := NewLexer(file)
	GrowParse(g)
	// ipt := getCurrentInterpreter()
	// sl := ipt.statement_list
	// fmt.Printf("%+v\n", sl.statement.expresion_s.int_value)
	// fmt.Printf("%+v\n", sl.next.statement.expresion_s)
}
