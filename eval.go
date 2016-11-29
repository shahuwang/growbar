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

func evalBooleanExpression(boolean_value bool) Value {
	var v Value
	v.typ = CRB_BOOLEAN_VALUE
	v.boolean_value = boolean_value
	return v
}

func releaseIfString(v *Value) {
	if v.typ == CRB_STRING_VALUE {
		releaseString(v.string_value)
	}
}

func evalNullExpression() Value {
	v := Value{typ: CRB_NULL_VALUE}
	return v
}
