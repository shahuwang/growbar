package main

import (
	"os"
)

func main() {
	file, err := os.Open("test.gb")
	if err != nil {
		panic(err)
	}
	g := NewLexer(file)
	GrowParse(g)
}
