package main

import (
	"fmt"
)

func nvPrintProc(ipt *Interpreter, arg_count int, args []Value) Value {
	value := Value{typ: CRB_NULL_VALUE}
	if arg_count < 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_FEW_ERR)
	} else if arg_count > 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_MANY_ERR)
	}
	v := args[0]
	switch v.typ {
	case CRB_BOOLEAN_VALUE:
		if v.boolean_value {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	case CRB_INT_VALUE:
		fmt.Println(v.int_value)
	case CRB_DOUBLE_VALUE:
		fmt.Println(v.double_value)
	case CRB_STRING_VALUE:
		fmt.Println(v.string_value.str)
	case CRB_NULL_VALUE:
		fmt.Println("null")
	case CRB_NATIVE_POINTER_VALUE:
		fmt.Printf("%s:%p", v.native_pointer.info.name, v.native_pointer.pointer)
	}
	return value
}
