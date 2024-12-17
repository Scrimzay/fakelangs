package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	input := "10000011" + "00000001" + "10000110" + "00000011" + "10000001"  // Represents: cheese IDENT = INT ;

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TOKEN_CHEESE, "10000011"},
		{TOKEN_IDENT, "00000001"},
		{TOKEN_ENCHILADA, "10000110"},
		{TOKEN_INT, "00000011"},
		{TOKEN_SEMICOLON, "10000001"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}