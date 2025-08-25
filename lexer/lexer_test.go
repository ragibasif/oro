package lexer

import (
	"testing"

	"github.com/ragibasif/oro/token"
)

func TestLexer(t *testing.T) {
	source := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Assignment, "="},
		{token.Plus, "+"},
		{token.LeftParentheses, "("},
		{token.RightParentheses, ")"},
		{token.LeftCurlyBrace, "{"},
		{token.RightCurlyBrace, "}"},
		{token.Comma, ","},
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
