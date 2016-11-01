package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Growlex struct {
	File   *os.File
	Pos    int
	Start  int
	Width  int
	reader *bufio.Reader
}

const GEOF = -1

func NewLexer(file *os.File) *Growlex {
	g := new(Growlex)
	g.File = file
	g.reader = bufio.NewReader(file)
	return g
}

func (g *Growlex) Next() rune {
	r, n, err := g.reader.ReadRune()
	if err == io.EOF {
		return GEOF
	}
	if err != nil {
		panic(err)
	}
	g.Width = n
	g.Pos = g.Pos + g.Width
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
	g.Pos = g.Pos - g.Width
}

func (g *Growlex) Ignore() {
	g.Start = g.Pos
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
		return LP
	case r == ')':
		g.Next()
		return RP
	case r == '{':
		g.Next()
		return LC
	case r == '}':
		g.Next()
		return RC
	case r == ';':
		g.Next()
		return SEMICOLON
	case r == ',':
		g.Next()
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
		return ADD
	case r == '-':
		g.Next()
		return SUB
	case r == '/':
		g.Next()
		return DIV
	case r == '%':
		g.Next()
		return MOD
	case r == '*':
		g.Next()
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
		return g.Lex(lval)
	case r == '0' || r == '1' || r == '2' ||
		r == '3' || r == '4' || r == '5' || r == '6' ||
		r == '7' || r == '8' || r == '9':
		return g.scanNumber()
	}
	return 0
}

func (g *Growlex) logicAnd() int {
	g.Next()
	r := g.Peek()
	if r == '&' {
		g.Next()
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
		return EQ
	}
	return ASSIGN
}

func (g *Growlex) notEqual() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
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
		return GE
	}
	return GT
}

func (g *Growlex) ltOrLte() int {
	g.Next()
	r := g.Peek()
	if r == '=' {
		g.Next()
		return LE
	}
	return LT
}

func (g *Growlex) winNewLine() int {
	g.Next()
	r := g.Peek()
	ipt := getCurrentInterpreter()
	if r == '\n' {
		g.Next()
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
	return 0
}

func (g *Growlex) scanNumber() int {
	digits := "0123456789"
	if g.accept("0") {
		// 只能是0或者小数
		if g.accept(".") {
			if g.accept(digits) {
				g.acceptRun(digits)
				if !isAlphaNumeric(g.Peek()) {
					return DOUBLE_LITERAL
				} else {
					g.Error("wrong number syntax")
					return -1
				}
			}
			g.Error("wrong number syntax")
			return -1
		}
		if !isAlphaNumeric(g.Peek()) {
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
				if !isAlphaNumeric(g.Peek()) {
					return DOUBLE_LITERAL
				}
			}
			g.Error("wrong number syntax")
			return -1
		}
		if !isAlphaNumeric(g.Peek()) {
			return INT_LITERAL
		}
	} else {
		msg := fmt.Sprintf("syntax7, %s", string(g.Peek()))
		g.Error(msg)
		return -1
	}
	return 0
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
