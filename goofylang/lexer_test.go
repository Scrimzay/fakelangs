package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
    input := `cheese x = 7;
              cheese y = 3;
              cheese z = x + y;
              pizza z;`

    tests := []struct {
        expectedType    TokenType
        expectedLiteral string
    }{
        {TOKEN_CHEESE, "cheese"},
        {TOKEN_IDENT, "x"},
        {TOKEN_ENCHILADA, "="},
        {TOKEN_INT, "7"},
        {TOKEN_SEMICOLON, ";"},
        {TOKEN_CHEESE, "cheese"},
        {TOKEN_IDENT, "y"},
        {TOKEN_ENCHILADA, "="},
        {TOKEN_INT, "3"},
        {TOKEN_SEMICOLON, ";"},
        {TOKEN_CHEESE, "cheese"},
        {TOKEN_IDENT, "z"},
        {TOKEN_ENCHILADA, "="},
        {TOKEN_IDENT, "x"},
        {TOKEN_APPLE, "+"},
        {TOKEN_IDENT, "y"},
        {TOKEN_SEMICOLON, ";"},
        {TOKEN_PIZZA, "pizza"},
        {TOKEN_IDENT, "z"},
        {TOKEN_SEMICOLON, ";"},
    }

    l := NewLexer(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}