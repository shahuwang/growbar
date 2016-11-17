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
