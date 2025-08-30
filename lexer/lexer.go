// Package lexer will turn source code into tokens
package lexer

import (
	"fmt"

	"github.com/ragibasif/oro/token"
)

type Lexer struct {
	source   string
	position int
	tokens   []token.Token
}

func (l *Lexer) String() string {
	return fmt.Sprintf("%s %v", l.source, l.tokens)
}

func (l *Lexer) ScanToken() token.Token {

	var c byte

	if isWhitespace(l.peek()) {
		l.skipWhitespace()
		c = l.peek()
	} else {
		c = l.step()
	}

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
		if isAlpha(c) {
			tok.Literal = l.identifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(c) {
			tok.Type = token.Number
			tok.Literal = l.number()
			return tok
		} else {
			tok = token.NewToken(token.Unknown, c)
		}
	}

	return tok
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

func (l *Lexer) number() string {
	position := l.position
	for isDigit(l.peek()) {
		l.step()
	}
	return l.source[position:l.position]
}

func (l *Lexer) identifier() string {
	position := l.position
	for isAlpha(l.peek()) {
		l.step()
	}
	fmt.Println("heer", l.source[position:l.position], position, l.position)
	return l.source[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for isNewLine(l.peek()) || isWhitespace(l.peek()) {
		l.step()
	}
}

func isNewLine(c byte) bool {
	return c == '\n'
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\r' || c == '\t'
}

func isAlpha(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (l *Lexer) peek() byte {
	if !l.withinBounds() {
		return 0
	}
	return l.source[l.position]
}

func (l *Lexer) step() byte {
	l.position++
	return l.source[l.position-1]
}

func (l *Lexer) withinBounds() bool {
	return l.position < len(l.source)
}
