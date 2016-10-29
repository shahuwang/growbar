package main

import (
	"bufio"
	"os"
)

type GrowLex struct {
	File   *os.File
	Pos    int
	Start  int
	Width  int
	reader bufio.Reader
}

func (g *GrowLex) Next() rune {
	r, n, err := g.reader.ReadRune()
}

type StateFn func(*GrowLex) StateFn

func StateAction(g *GrowLex) StateFn {

}
