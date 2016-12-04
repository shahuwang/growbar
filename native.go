package main

import (
	"fmt"
	"os"
)

var st_native_lib_info NativePointerInfo = NativePointerInfo{
	name: "growbar.lang.file",
}

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

func nvFopenProc(ipt *Interpreter, arg_count int, args []Value) Value {
	if arg_count < 2 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_FEW_ERR)
	} else if arg_count > 2 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_MANY_ERR)
	}
	v := args[0]
	v2 := args[1]
	if v.typ != CRB_STRING_VALUE || v2.typ != CRB_STRING_VALUE {
		runtimeError(ipt.current_line_number, FOPEN_ARGUMENT_TYPE_ERR)
	}
	var mode int
	var value Value
	switch v2.string_value.str {
	case "r":
		mode = os.O_RDONLY
	case "r+", "rb+":
		mode = os.O_RDONLY | os.O_WRONLY
	case "rt+":
		mode = os.O_RDONLY | os.O_TRUNC
	case "w", "wb":
		mode = os.O_WRONLY | os.O_CREATE
	case "w+", "wb+":
		mode = os.O_WRONLY | os.O_RDONLY | os.O_CREATE
	case "wt+":
		mode = os.O_WRONLY | os.O_RDONLY | os.O_TRUNC
	case "a":
		mode = os.O_APPEND | os.O_CREATE
	case "a+", "ab+":
		mode = os.O_APPEND | os.O_RDONLY | os.O_CREATE
	case "at+":
		mode = os.O_APPEND | os.O_RDONLY | os.O_TRUNC
	}
	f, err := os.OpenFile(v.string_value.str, mode, 0660)
	if err != nil {
		value.typ = CRB_NATIVE_POINTER_VALUE
	} else {
		value.typ = CRB_NATIVE_POINTER_VALUE
		value.native_pointer = *new(NativePointer)
		value.native_pointer.info = &st_native_lib_info
		value.native_pointer.pointer = f
	}
	return value
}
