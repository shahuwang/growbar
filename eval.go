package main

import ()

func evalIntExpression(v int) Value {
	value := Value{typ: CRB_INT_VALUE, int_value: v}
	return value
}

func evalDoubleExpression(v float32) Value {
	value := Value{typ: CRB_DOUBLE_VALUE, double_value: v}
	return value
}

func evalCompareString(operator ExpressionType, left *Value, right *Value, line_number int) bool {
	//TODO
	return true
}
