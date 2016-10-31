package main

import (
	"bufio"
	"io"
	"os"
	// "unicode"
	"fmt"
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
		ipt := getCurrentInterpreter()
		compileError(
			ipt.current_line_number,
			CHARACTER_INVALID_ERR, string(r))
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
	panic(s)
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
	ipt := getCurrentInterpreter()
	msg := fmt.Sprintf("line %d, & must follow another &", ipt.current_line_number)
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
	ipt := getCurrentInterpreter()
	msg := fmt.Sprintf("line %d, | must follow another |", ipt.current_line_number)
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
	ipt := getCurrentInterpreter()
	msg := fmt.Sprintf("line %d, ! must follow by =", ipt.current_line_number)
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
	msg := fmt.Sprintf("line %d, error windows new line symbol \r", ipt.current_line_number)
	g.Error(msg)
	return -1
}

func (g *Growlex) unixNewLine() int {
	g.Next()
	ipt := getCurrentInterpreter()
	ipt.current_line_number++
	return 0
}
