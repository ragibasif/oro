// Package token defines the tokens that the source code will be converted to
package token

import "fmt"

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("%d %s", t.Type, t.Literal)
}

func NewToken(tt TokenType, l byte) Token {
	t := new(Token)
	t.Type = tt
	t.Literal = string(l)
	return *t
}

const (
	EOF TokenType = iota

	// Keywords
	Variable // var
	Function // fn
	Return   // return
	Print    // print
	If       // if
	Else     // else
	For      // for
	While    // while
	True     // true
	False    // false
	Null     // null
	Break    // break
	Continue // continue

	// Literals
	Identifier
	String
	Number
	Bool // true, false

	// Arithmetic Operators
	Plus     // +
	Minus    // -
	Multiply // *
	Divide   // /
	Modulo   // %
	Negate   // -

	// Bitwise Operators
	BitwiseAnd        // &
	BitwiseOr         // |
	BitwiseNot        // ~
	BitwiseXor        // ^
	BitwiseRightShift // >>
	BitwiseLeftShift  // <<

	// Relational Operators
	Eqality      // ==
	InEquality   // !=
	LessThan     // <
	GreaterThan  // >
	LessEqual    // <=
	GreaterEqual // >=

	// Logical Operators
	LogicalAnd // &&
	LogicalOr  // ||
	LogicalNot // !

	// Assignment Operators
	Assignment                  // =
	PlusAssignment              // +=
	MinusAssignment             // -=
	MultiplyAssignment          // *=
	DivideAssignment            // /=
	ModuloAssignment            // %=
	BitwiseAndAssignment        // &=
	BitwiseOrAssignment         // |=
	BitwiseXorAssignment        // ^=
	BitwiseLeftShiftAssignment  // <<=
	BitwiseRightShiftAssignment // >>=

	// Delimiters
	Question           // ?
	Colon              // :
	Semicolon          // ;
	Dot                // .
	Comma              // ,
	LeftParentheses    // (
	RightParentheses   // )
	LeftCurlyBrace     // {
	RightCurlyBrace    // }
	LeftSquareBracket  // [
	RightSquareBracket // ]

	// Escape Sequences
	EscapeBell           // \a
	EscapeBackSpace      // \b
	EscapeFormFeed       // \f
	EscapeNewLine        // \n
	EscapeCarriageReturn // \r
	EscapeHorizontalTab  // \t
	EscapeVerticalTab    // \v
	EscapeBackSlash      // \\
	EscapeQuote          // \"
	EscapeApostrophe     // \'

	// Others
	Comment // #

	// Internal
	Unknown // unknown tokens
	Count   // number of tokens
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"var":    Variable,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return Identifier
}
