// Package lexer will turn source code into tokens
package lexer

import (
	"fmt"

	"github.com/ragibasif/oro/token"
)

// lexer == tokenizer == scanner
// identifiers == variable names, etc...
// keywords == var, fn, for, etc...

type Lexer struct {
	source   string
	position int
	tokens   []token.Token
}

func (l *Lexer) String() string {
	return fmt.Sprintf("%s %v", l.source, l.tokens)
}

func (l *Lexer) ScanToken() token.Token {

	c := l.nextChar()

	var tok token.Token

	switch c {
	case '=':
		tok = token.NewToken(token.Assignment, c)
	case '+':
		tok = token.NewToken(token.Plus, c)
	case '-':
		tok = token.NewToken(token.Minus, c)
	case '*':
		tok = token.NewToken(token.Multiply, c)
	case '/':
		tok = token.NewToken(token.Divide, c)
	case '(':
		tok = token.NewToken(token.LeftParentheses, c)
	case ')':
		tok = token.NewToken(token.RightParentheses, c)
	case '{':
		tok = token.NewToken(token.LeftCurlyBrace, c)
	case '}':
		tok = token.NewToken(token.RightCurlyBrace, c)
	case ',':
		tok = token.NewToken(token.Comma, c)
	case ';':
		tok = token.NewToken(token.Semicolon, c)
	default:
		tok = token.NewToken(token.Unknown, c)
	}

	return tok
}

func (l *Lexer) withinBounds() bool {
	return l.position < len(l.source)
}

func (l *Lexer) nextChar() byte {
	l.position++
	return l.source[l.position-1]
}

func (l *Lexer) ScanTokens() []token.Token {
	for l.withinBounds() {
		l.tokens = append(l.tokens, l.ScanToken())
	}

	tok := token.Token{Type: token.EOF, Literal: ""}
	l.tokens = append(l.tokens, tok)
	return l.tokens
}

func NewLexer(source string) *Lexer {
	l := new(Lexer)
	l.source = source
	return l
}
