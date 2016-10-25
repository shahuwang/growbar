package main

import (
	"testing"
)

func TestReflect(t *testing.T) {
	v := new(CRBValue)
	pinfo := NativePointerInfo{name: "hello"}
	pointer := NativePointer{info: &pinfo}
	v.u = pointer
	p, ok := v.u.(NativePointer)
	if ok {
		p.pointer = 12
	}
}
