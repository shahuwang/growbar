package main

import (
// "reflect"
)

type ValueType int

const (
	CRB_BOOLEAN_VALUE ValueType = iota + 1
	CRB_INT_VALUE
	CRB_DOUBLE_VALUE
	CRB_STRING_VALUE
	CRB_NATIVE_POINTER_VALUE
	CRB_NULL_VALUE
)

type NativePointerInfo struct {
	name string
}

type NativePointer struct {
	info    *NativePointerInfo
	pointer interface{}
}
type Value struct {
	typ ValueType
	u   interface{}
}
