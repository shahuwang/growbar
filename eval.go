package main

import ()

func evalIntExpression(v int) Value {
	value := Value{typ: CRB_INT_VALUE, int_value: v}
	return value
}
