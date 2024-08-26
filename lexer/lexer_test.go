package lexer

import (
	"json-parser/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	t.Run("Test next token", func(t *testing.T) {
		input := `{}(),;`
		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.BRACE_OPEN, "{"},
			{token.BRACE_CLOSE, "}"},
			{token.LEFT_PARENTHESIS, "("},
			{token.RIGHT_PARENTHESIS, ")"},
			{token.BRACE_OPEN, ")"},
			{token.BRACE_OPEN, ","},
			{token.BRACE_OPEN, ";"},
		}
		l := New(input)
		for i, tt := range tests {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedLiteral, tok.Literal)
			}
		}
	})
}
