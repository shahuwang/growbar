package main

import (
	"bufio"
	"fmt"
	"io"
	// "io/ioutil"
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
		mode = os.O_RDONLY | os.O_CREATE
	case "r+", "rb+":
		mode = os.O_RDONLY | os.O_WRONLY | os.O_CREATE
	case "w", "wb":
		mode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	case "w+", "wb+":
		mode = os.O_WRONLY | os.O_RDONLY | os.O_CREATE | os.O_TRUNC
	case "a":
		mode = os.O_APPEND | os.O_CREATE
	case "a+", "ab+":
		mode = os.O_APPEND | os.O_RDONLY | os.O_CREATE
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

func nvFgetsProc(ipt *Interpreter, arg_count int, args []Value) Value {
	if arg_count < 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_FEW_ERR)
	} else if arg_count > 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_MANY_ERR)
	}
	v := args[0]
	if v.typ != CRB_NATIVE_POINTER_VALUE || !checkNativePointer(&v) {
		runtimeError(ipt.current_line_number, FGETS_ARGUMENT_TYPE_ERR)
	}
	fp := v.native_pointer.pointer.(*os.File)
	reader := bufio.NewReaderSize(fp, LINE_BUF_SIZE)
	result := make([]byte, 0)
	buf := make([]byte, LINE_BUF_SIZE)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		result = append(result, buf[:n]...)
	}
	var value Value
	if len(result) > 0 {
		value.typ = CRB_STRING_VALUE
		value.string_value = ipt.createCrowbarString(string(result))
	} else {
		value.typ = CRB_NULL_VALUE
	}
	return value
}

func nvFputProc(ipt *Interpreter, arg_count int, args []Value) Value {
	if arg_count < 2 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_FEW_ERR)
	} else if arg_count > 2 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_MANY_ERR)
	}
	v := args[0]
	v2 := args[1]
	if v.typ != CRB_STRING_VALUE || (v2.typ != CRB_NATIVE_POINTER_VALUE || !checkNativePointer(&v2)) {
		runtimeError(ipt.current_line_number, FPUTS_ARGUMENT_TYPE_ERR)
	}
	fp := v2.native_pointer.pointer.(*os.File)
	_, err := fp.WriteString(v.string_value.str)
	if err != nil {
		panic(err)
	}
	return Value{typ: CRB_NULL_VALUE}
}

func nvFclose(ipt *Interpreter, arg_count int, args []Value) Value {
	if arg_count < 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_FEW_ERR)
	} else if arg_count > 1 {
		runtimeError(ipt.current_line_number, ARGUMENT_TOO_MANY_ERR)
	}
	v := args[0]
	if v.typ != CRB_NATIVE_POINTER_VALUE || !checkNativePointer(&v) {
		runtimeError(ipt.current_line_number, FCLOSE_ARGUMENT_TYPE_ERR)
	}
	fp := v.native_pointer.pointer.(*os.File)
	err := fp.Close()
	if err != nil {
		panic(err)
	}
	return Value{typ: CRB_NULL_VALUE}
}

func checkNativePointer(value *Value) bool {
	return value.native_pointer.info == &st_native_lib_info
}
