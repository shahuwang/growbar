package main

import ()

func referIfString(v *Value) {
	if v.typ == CRB_STRING_VALUE {
		referString(v.string_value)
	}
}

func referString(str CRBString) {
	str.ref_count++
}
