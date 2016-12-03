package main

import (
	"fmt"
)

func getOperatorString(typ ExpressionType) string {
	switch typ {
	case ASSIGN_EXPRESSION:
		return "="
	case ADD_EXPRESSION:
		return "+"
	case SUB_EXPRESSION:
		return "-"
	case MUL_EXPRESSION:
		return "*"
	case MOD_EXPRESSION:
		return "%"
	case DIV_EXPRESSION:
		return "/"
	case LOGICAL_AND_EXPRESSION:
		return "&&"
	case LOGICAL_OR_EXPRESSION:
		return "||"
	case EQ_EXPRESSION:
		return "=="
	case NE_EXPRESSION:
		return "!="
	case GT_EXPRESSION:
		return ">"
	case GE_EXPRESSION:
		return ">="
	case LT_EXPRESSION:
		return "<"
	case LE_EXPRESSION:
		return "<="
	case MINUS_EXPRESSION:
		return "-"
	case BOOLEAN_EXPRESSION, INT_EXPRESSION,
		DOUBLE_EXPRESSION, IDENTIFIER_EXPRESSION,
		FUNCTION_CALL_EXPRESSION, NULL_EXPRESSION,
		EXPRESSION_TYPE_COUNT_PLUS_1:
		fallthrough
	default:
		msg := fmt.Sprintf("bad expression type ..%d\n", typ)
		panic(msg)
	}
}

func releaseString(str *CRBString) {
	str.ref_count--
	if str.ref_count == 0 {

	}
}

func searchLocalVariable(env *LocalEnvironment, identifier string) *Variable {
	if env == nil {
		return nil
	}
	pos := env.variable
	for {
		if pos == nil || pos.name == identifier {
			break
		}
		pos = pos.next
	}
	return pos
}

func addLocalVariable(env *LocalEnvironment, identifier string, value *Value) {
	newV := new(Variable)
	newV.name = identifier
	newV.value = *value
	newV.next = env.variable
	env.variable = newV
}

func searchFunction(name string) *FunctionDefinition {
	var pos *FunctionDefinition
	ipt := getCurrentInterpreter()
	pos = ipt.function_list
	for {
		if pos == nil {
			break
		}
		if pos.name == name {
			return pos
		}
		pos = pos.next
	}
	return pos
}

func allocLocalEnvironment() *LocalEnvironment {
	ret := new(LocalEnvironment)
	return ret
}

func PrintResult(result *StatementResult) {
	v := result.return_value
	switch v.typ {
	case CRB_INT_VALUE:
		fmt.Printf("%d\n", v.int_value)
	case CRB_DOUBLE_VALUE:
		fmt.Printf("%f\n", v.double_value)
	case CRB_BOOLEAN_VALUE:
		fmt.Printf("%t\n", v.boolean_value)
	case CRB_STRING_VALUE:
		fmt.Println(v.string_value.str)
	case CRB_NULL_VALUE:
		fmt.Println("nil")
	default:
		fmt.Printf("%+v", v)
	}
}
