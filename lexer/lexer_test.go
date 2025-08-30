package lexer

import (
	"fmt"
	"testing"

	"github.com/ragibasif/oro/token"
)

func TestLexer(t *testing.T) {
	source := `var five = 5;
var three = 3;

fn add(a, b) {
    return a + b;
}

var result = add(five, three);
print(result);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Variable, "var"},
		{token.Identifier, "five"},
		{token.Assignment, "="},
		{token.Number, "5"},
		{token.Semicolon, ";"},

		{token.Variable, "var"},
		{token.Identifier, "three"},
		{token.Assignment, "="},
		{token.Number, "3"},
		{token.Semicolon, ";"},

		{token.Function, "fn"},
		{token.Identifier, "add"},
		{token.LeftParentheses, "("},
		{token.Identifier, "a"},
		{token.Identifier, "b"},
		{token.RightParentheses, ")"},
		{token.LeftCurlyBrace, "{"},
		{token.Return, "return"},
		{token.Identifier, "a"},
		{token.Plus, "+"},
		{token.Identifier, "b"},
		{token.Semicolon, ";"},
		{token.RightCurlyBrace, "}"},

		{token.Variable, "var"},
		{token.Identifier, "result"},
		{token.Assignment, "="},
		{token.Identifier, "add"},
		{token.LeftParentheses, "("},
		{token.Identifier, "five"},
		{token.Identifier, "three"},
		{token.RightParentheses, ")"},
		{token.Semicolon, ";"},

		{token.Print, "print"},
		{token.LeftParentheses, "("},
		{token.Identifier, "result"},
		{token.RightParentheses, ")"},
		{token.Semicolon, ";"},

		{token.EOF, ""},
	}

	t.Run("NewLexer", func(t *testing.T) {
		l := NewLexer(source)
		expectedSource := source
		gotSource := l.source
		if expectedSource != gotSource {
			t.Fatalf("source wrong, expected=%q, got=%q", expectedSource, gotSource)
		}
	})

	t.Run("ScanToken", func(t *testing.T) {
		l := NewLexer(source)
		for i := range len(tests) - 1 {
			tok := l.ScanToken()

			if tok.Type != tests[i].expectedType {
				fmt.Println(tok.String(), tests[i].expectedType, tests[i].expectedLiteral)
				t.Fatalf("tests[%d] - TokenType wrong, expected=%q, got=%q", i, tests[i].expectedType, tok.Type)
			}

			if tok.Literal != tests[i].expectedLiteral {
				t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tests[i].expectedLiteral, tok.Literal)
			}
		}

	})

	t.Run("ScanTokens", func(t *testing.T) {
		l := NewLexer(source)
		l.ScanTokens()
		for i, tt := range tests {
			tok := l.tokens[i]

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tests[i].expectedLiteral {
				t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tests[i].expectedLiteral, tok.Literal)
			}
		}

	})

}
