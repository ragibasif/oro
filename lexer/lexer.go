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

	l.skipWhitespace()
	c = l.peek()

	var tok token.Token

	// Keywords
	// Variable // var
	// Function // fn
	// Return   // return
	// Print    // print
	// If       // if
	// Else     // else
	// For      // for
	// While    // while
	// True     // true
	// False    // false
	// Null     // null
	// Break    // break
	// Continue // continue

	// Literals
	// String
	// Bool // true, false

	// // Escape Sequences
	// EscapeBell           // \a
	// EscapeBackSpace      // \b
	// EscapeFormFeed       // \f
	// EscapeNewLine        // \n
	// EscapeCarriageReturn // \r
	// EscapeHorizontalTab  // \t
	// EscapeVerticalTab    // \v
	// EscapeBackSlash      // \\
	// EscapeQuote          // \"
	// EscapeApostrophe     // \'

	switch c {

	// Arithmetic Operators
	case '+':
		tok = token.NewToken(token.Plus, c)
	case '-':
		tok = token.NewToken(token.Minus, c)
	case '*':
		tok = token.NewToken(token.Multiply, c)
	case '/':
		tok = token.NewToken(token.Divide, c)
	case '%':
		tok = token.NewToken(token.Modulo, c)

	// Assignment Operators
	// PlusAssignment              // +=
	// MinusAssignment             // -=
	// MultiplyAssignment          // *=
	// DivideAssignment            // /=
	// ModuloAssignment            // %=
	// BitwiseAndAssignment        // &=
	// BitwiseOrAssignment         // |=
	// BitwiseXorAssignment        // ^=
	// BitwiseLeftShiftAssignment  // <<=
	// BitwiseRightShiftAssignment // >>=
	case '=':
		tok = token.NewToken(token.Assignment, c)

	// Delimiters
	case '(':
		tok = token.NewToken(token.LeftParenthesis, c)
	case ')':
		tok = token.NewToken(token.RightParenthesis, c)
	case '{':
		tok = token.NewToken(token.LeftCurlyBrace, c)
	case '}':
		tok = token.NewToken(token.RightCurlyBrace, c)
	case '[':
		tok = token.NewToken(token.LeftSquareBracket, c)
	case ']':
		tok = token.NewToken(token.RightSquareBracket, c)
	case ',':
		tok = token.NewToken(token.Comma, c)
	case ';':
		tok = token.NewToken(token.Semicolon, c)
	case ':':
		tok = token.NewToken(token.Colon, c)
	case '?':
		tok = token.NewToken(token.Question, c)
	case '.':
		tok = token.NewToken(token.Dot, c)

	// Logical Operators
	// LogicalAnd // &&
	// LogicalOr  // ||
	case '!':
		tok = token.NewToken(token.LogicalNot, c)

	// Relational Operators
	// Eqality      // ==
	// InEquality   // !=
	// LessEqual    // <=
	// GreaterEqual // >=
	case '<':
		tok = token.NewToken(token.LessThan, c)
	case '>':
		tok = token.NewToken(token.GreaterThan, c)

	// Bitwise Operators
	// BitwiseRightShift // >>
	// BitwiseLeftShift  // <<
	case '&':
		tok = token.NewToken(token.BitwiseAnd, c)
	case '|':
		tok = token.NewToken(token.BitwiseOr, c)
	case '~':
		tok = token.NewToken(token.BitwiseNot, c)
	case '^':
		tok = token.NewToken(token.BitwiseXor, c)

	// Others
	case '#':
		tok = token.NewToken(token.Comment, c)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

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

	l.step()
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
	if !l.withinBounds() {
		return 0
	}
	l.position++
	return l.source[l.position-1]
}

func (l *Lexer) withinBounds() bool {
	return l.position < len(l.source)
}
