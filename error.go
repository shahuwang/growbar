package main

import (
	"fmt"
	"os"
)

var compileErrorMessageFormat [5]MessageFormat = [5]MessageFormat{
	MessageFormat{"dummy"}, MessageFormat{"在(%s)附近发生语法错误"},
	MessageFormat{"不正确的字符(%s)"}, MessageFormat{"函数名重复(%s)"}, MessageFormat{"dummy"},
}

var runtimeErrorMessageFormat []MessageFormat = []MessageFormat{
	MessageFormat{"dummy"},
	MessageFormat{"找不到变量(%s)"},
	MessageFormat{"找不到函数(%s)"},
	MessageFormat{"传入的参数数量多于函数定义"},
	MessageFormat{"传入的参数数量少于函数定义"},
	MessageFormat{"条件表达式的值必须是boolean型"},
	MessageFormat{"减法运算的操作数必须是数值类型"},
	MessageFormat{"双目操作符(%s)的操作数类型不正确"},
	MessageFormat{"(%s)操作符不能用于boolean型"},
	MessageFormat{"请为fopen()函数传入文件的路径和打开方式"},
	MessageFormat{"请为fclose()函数传入文件指针"},
	MessageFormat{"请为fgets()函数传入文件指针"},
	MessageFormat{"请为fputs()函数传入文件指针和字符串"},
	MessageFormat{"null只能用于运算符 == 和 != (不能进行 %s 操作)"},
	MessageFormat{"不能被0除"},
	MessageFormat{"全局变量%s不存在"},
	MessageFormat{"不能再函数外使用global语句"},
	MessageFormat{"运算符(%s)不能用于字符串类型"},
	MessageFormat{"dummy"},
}

func errorFormat(line_number int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args)
	fmt.Fprintf(os.Stderr, "Error is in line %d, %s", line_number, msg)
}

func runtimeError(line_number int, id RuntimeError, args ...interface{}) {
	format := runtimeErrorMessageFormat[id].format
	errorFormat(line_number, format, args)
	os.Exit(1)
}

func compileError(line_number int, id CompileError, args ...interface{}) {
	format := compileErrorMessageFormat[id].format
	errorFormat(line_number, format, args)
	os.Exit(1)
}
