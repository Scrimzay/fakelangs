package main

import (
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := "10000011" + "00000001" + "10000110" + "00000011" + "10000001"  // Represents: cheese IDENT = INT ;

	l := NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	testLetStatement(t, program.Statements[0], "00000001") // Check if the first statement is a let statement with the expected identifier
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

func testLetStatement(t *testing.T, s Statement, name string) {
	if s.TokenLiteral() != "10000011" { // "10000011" is the binary string for TOKEN_CHEESE
		t.Errorf("s.TokenLiteral not '10000011'. got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
	}
}