package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//TODO : 记得要删除掉
// type GrowSymType struct{}
type GrowLex struct {
	File   *os.File
	Pos    int
	Start  int
	Width  int
	reader bufio.Reader
}

const GEOF = -1

func NewLexer(file *os.File) *GrowLex {
	g := new(GrowLex)
	g.File = file
	g.reader = bufio.NewReader(file)
	return g
}

func (g *GrowLex) Next() rune {
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

func (g *GrowLex) Peek() rune {
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

func (g *GrowLex) Back() {
	err := g.reader.UnreadRune()
	if err != nil {
		panic(err)
	}
	g.Pos = g.Pos - g.Width
}

func (g *GrowLex) Ignore() {
	g.Start = g.Pos
}

func (g *GrowLex) Lex(lval *GrowSymType) int {
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
	case r == '=':
		g.Next()
		return ASSIGN
	case r == '>':
		g.Next()
		return GT
	case r == '<':
		g.Next()
		return LT
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
		return MOD
	}
	return 0
}
