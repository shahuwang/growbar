package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Growlex struct {
	File   *os.File
	Pos    int
	Start  int
	runes  []rune
	reader *bufio.Reader
}

const GEOF = -1

var Keywords = map[string]int{
	"function": FUNCTION,
	"if":       IF,
	"else":     ELSE,
	"elsif":    ELSIF,
	"while":    WHILE,
	"for":      FOR,
	"return":   RETURN_T,
	"break":    BREAK,
	"continue": CONTINUE,
	"true":     TRUE_T,
	"false":    FALSE_T,
	"global":   GLOBAL_T,
}

func NewLexer(file *os.File) *Growlex {
	g := new(Growlex)
	g.File = file
	g.reader = bufio.NewReader(file)
	return g
}

func (g *Growlex) Next() rune {
	r, _, err := g.reader.ReadRune()
	if err == io.EOF {
		return GEOF
	}
	if err != nil {
		panic(err)
	}
	g.runes = append(g.runes, r)
	g.Pos++
	return r
}

func (g *Growlex) Peek() rune {
	b, _ := g.reader.Peek(4)
	r, n := utf8.DecodeRune(b)
	if r == utf8.RuneError && n == 0 {
		return GEOF
	}
	if r == utf8.RuneError && n == 1 {
		g.Error("invalid character")
	}
	return r
}

func (g *Growlex) Back() {
	err := g.reader.UnreadRune()
	if err != nil {
		panic(err)
	}
	g.Pos = g.Pos - 1
	g.runes = g.runes[0:g.Pos]
}

func (g *Growlex) Ignore() {
	g.Start = g.Pos
}

func (g *Growlex) ignoreLast() {
	// 用于处理字符串里面的转义字符
	g.runes = g.runes[0 : g.Pos-2]
	g.Pos = g.Pos - 1
}

func (g *Growlex) hit() string {
	if len(g.runes) == 0 {
		g.Start = 0
		g.Pos = 0
		return ""
	}
	val := g.runes[g.Start:g.Pos]
	g.Ignore()
	return string(val)
}

func (g *Growlex) Error(s string) {
	ipt := getCurrentInterpreter()
	msg := fmt.Sprintf("line %d: %s", ipt.current_line_number, s)
	panic(msg)
}

func (g *Growlex) accept(valid string) bool {
	if strings.IndexRune(valid, g.Next()) >= 0 {
		return true
	}
	g.Back()
	return false
}

func (g *Growlex) acceptRun(valid string) {
	for strings.IndexRune(valid, g.Next()) >= 0 {

	}
	g.Back()
}

func (g *Growlex) Lex(lval *GrowSymType) int {
	r := g.Peek()
	switch {
	case r == '(':
		g.Next()
		g.hit()
		return LP
	case r == ')':
		g.Next()
		g.hit()
		return RP
	case r == '{':
		g.Next()
		g.hit()
		return LC
	case r == '}':
		g.Next()
		g.hit()
		return RC
	case r == ';':
		g.Next()
		g.hit()
		return SEMICOLON
	case r == ',':
		g.Next()
		g.hit()
		return COMMA
	case r == '&':
		return g.logicAnd()
	case r == '|':
		return g.loginOr()
	case r == '=':
		return g.assignOrEqual()
	case r == '!':
		return g.notEqual()
	case r == '>':
		return g.gtOrGte()
	case r == '<':
		return g.ltOrLte()
	case r == '+':
		g.Next()
		g.hit()
		return ADD
	case r == '-':
		g.Next()
		g.hit()
		return SUB
	case r == '/':
		g.Next()
		g.hit()
		return DIV
	case r == '%':
		g.Next()
		g.hit()
		return MOD
	case r == '*':
		g.Next()
		g.hit()
		return MUL
	case r == '\r':
		if g.winNewLine() == 0 {
			return g.Lex(lval)
		}
	case r == '\n':
		if g.unixNewLine() == 0 {
			return g.Lex(lval)
		}
	case r == ' ' || r == '\t':
		g.Next()
		g.hit()
		return g.Lex(lval)
	case unicode.IsDigit(r):
		return g.scanNumber(lval)
	case isAlpha(r):
		return g.keywordOrIdentifier(lval)
	case r == '"':
		return g.scanString(lval)
	case r == '#':
		g.scanComment()
		return g.Lex(lval)
	case r == GEOF:
		return 0
	default:
		fmt.Println(r)
		g.Error("error character: " + string(r))
		return -1
	}
	return 0
}

func (g *Growlex) logicAnd() int {
	g.Next()
	r := g.Peek()
	if r == '&' {
		g.Next()
		g.hit()
		return LOGICAL_AND
	}
	msg := "& must follow another &"
	g.Error(msg)
	return -1
}

func (g *Growlex) loginOr() int {
	g.Next()
	r := g.Peek()
	if r == '|' {
		g.Next()
		g.hit()
		return LOGICAL_OR
	}
	msg := "| must follow another |"
	g.Error(msg)
	return -1
}

func (g *Growlex) assignOrEqual() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
		g.hit()
		return EQ
	}
	g.hit()
	return ASSIGN
}

func (g *Growlex) notEqual() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
		g.hit()
		return NE
	}
	msg := "! must follow by ="
	g.Error(msg)
	return -1
}

func (g *Growlex) gtOrGte() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
		g.hit()
		return GE
	}
	g.hit()
	return GT
}

func (g *Growlex) ltOrLte() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
		g.hit()
		return LE
	}
	g.hit()
	return LT
}

func (g *Growlex) winNewLine() int {
	g.Next()
	r := g.Peek()
	ipt := getCurrentInterpreter()
	if r == '\n' {
		g.Next()
		g.hit()
		ipt.current_line_number++
		return 0
	}
	msg := "error windows new line symbol \r"
	g.Error(msg)
	return -1
}

func (g *Growlex) unixNewLine() int {
	g.Next()
	ipt := getCurrentInterpreter()
	ipt.current_line_number++
	g.hit()
	return 0
}

func (g *Growlex) scanNumber(lval *GrowSymType) int {
	exp := allocExpression(INT_EXPRESSION)
	lval.expression = exp
	runes := make([]rune, 0)
	digits := "0123456789"
	if g.accept("0") {
		// 只能是0或者小数
		runes = append(runes, '0')
		if g.accept(".") {
			if g.accept(digits) {
				g.acceptRun(digits)
				if !isAlpha(g.Peek()) {
					num := g.hit()
					f, err := strconv.ParseFloat(num, 32)
					if err != nil {
						g.Error(err.Error())
						return -1
					}
					exp.double_value = float32(f)
					return DOUBLE_LITERAL
				} else {
					g.Error("wrong number syntax")
					return -1
				}
			}
			g.Error("wrong number syntax")
			return -1
		}
		if !isAlpha(g.Peek()) {
			num := g.hit()
			ival, err := strconv.ParseInt(num, 10, 32)
			if err != nil {
				g.Error(err.Error())
				return -1
			}
			exp.int_value = int(ival)
			return INT_LITERAL
		}
		g.Error("wrong number syntax")
		return -1
	}
	if g.accept(digits) {
		g.acceptRun(digits)
		if g.accept(".") {
			if g.accept(digits) {
				g.acceptRun(digits)
				if !isAlpha(g.Peek()) {
					num := g.hit()
					f, err := strconv.ParseFloat(num, 32)
					if err != nil {
						g.Error(err.Error())
						return -1
					}
					exp.double_value = float32(f)
					return DOUBLE_LITERAL
				}
			}
			g.Error("wrong number syntax")
			return -1
		}
		if !isAlpha(g.Peek()) {
			num := g.hit()
			ival, err := strconv.ParseInt(num, 10, 32)
			if err != nil {
				g.Error(err.Error())
				return -1
			}
			exp.int_value = int(ival)
			return INT_LITERAL
		}
	} else {
		msg := fmt.Sprintf("syntax7, %s", string(g.Peek()))
		g.Error(msg)
		return -1
	}
	return 0
}

func (g *Growlex) keywordOrIdentifier(lval *GrowSymType) int {
	for {
		g.Next()
		if !isAlphaNumeric(g.Peek()) {
			break
		}
	}
	word := g.hit()
	keyword, ok := Keywords[word]
	if ok {
		return keyword
	}
	lval.identifier = word
	return IDENTIFIER
}

func (g *Growlex) scanString(lval *GrowSymType) int {
	exp := allocExpression(STRING_EXPRESSION)
	lval.expression = exp
	g.Next()
LOOP:
	for {
		r := g.Peek()
		switch {
		case r == '\\':
			g.Next()
			r2 := g.Peek()
			switch {
			case r2 == 'n':
				// 换行转义符
				g.Next()
				g.ignoreLast()
				g.runes = append(g.runes, '\n')
			case r2 == 't':
				g.Next()
				g.ignoreLast()
				g.runes = append(g.runes, '\t')
			case r2 == '"':
				g.Next()
				g.ignoreLast()
				g.runes = append(g.runes, '"')
			default:
				g.Next()
			}
		case r == '"':
			g.Next()
			exp.string_value = g.hit()
			return STRING_LITERAL
		case r == '\r' || r == '\n':
			g.Next()
			ipt := getCurrentInterpreter()
			ipt.current_line_number++
		case r == GEOF:
			g.Next()
			break LOOP
		default:
			g.Next()
		}
	}
	g.Error("string without \" enclose")
	return -1
}

func (g *Growlex) scanComment() {
	g.Next()
	for {
		r := g.Peek()
		if r == '\n' || r == '\r' || r == GEOF {
			g.hit()
			break
		}
		g.Next()
	}
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
