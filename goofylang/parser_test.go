package main

import (
	"testing"
)

func TestLetStatements(t *testing.T) {
    input := `
    cheese x = 5;
    cheese y = 10;
    cheese foobar = 838383;
    `

    l := NewLexer(input)
    p := NewParser(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    // ... additional tests on the contents of program.Statements ...
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}