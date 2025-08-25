package main

import (
	"fmt"

	"github.com/ragibasif/oro/lexer"
)

func main() {
	print("Oro\n")
	l := lexer.NewLexer("Hello world")
	s := l.String()
	fmt.Println(s)
}
