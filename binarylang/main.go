package main

import (
	"strings"
	"fmt"
	"strconv"
)

type TokenType string

// Pizza for printing, Cheese for var declaration, Tacos for identifying, Nuggets for number
// Enchilada for equal, Apple for add, Salmon for subtract, Icaco for inputs
// h
const (
	TOKEN_IDENT   = "00000001" // Unique binary code for IDENT
    TOKEN_EOF     = "00000010" // Unique binary code for EOF
    TOKEN_INT     = "00000011" // Unique binary code for INT
    TOKEN_ILLEGAL = "00000100" // Unique binary code for ILLEGAL
	TOKEN_SEMICOLON = "10000001" // Arbitrary unique binary code for semicolon
    TOKEN_PIZZA     = "10000010" // Arbitrary unique binary code for pizza
    TOKEN_CHEESE    = "10000011" // Arbitrary unique binary code for cheese
    TOKEN_TACOS     = "10000100" // Arbitrary unique binary code for tacos
    TOKEN_NUGGETS   = "10000101" // Arbitrary unique binary code for nuggets
    TOKEN_ENCHILADA = "10000110" // Arbitrary unique binary code for enchilada
    TOKEN_APPLE     = "10000111" // Arbitrary unique binary code for apple
    TOKEN_SALMON    = "10001000" // Arbitrary unique binary code for salmon
    TOKEN_ICACO     = "10001001" // Arbitrary unique binary code for icaco
)

const (
    _ int = iota  // iota resets in each const block
    LOWEST
    EQUALS      // ==
    LESSGREATER // > or <
    SUM         // +
    PRODUCT     // *
    PREFIX      // -X or !X
    CALL        // myFunction(X)
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int  // current position in input
	currentToken Token
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: input, position: 0} // Explicitly set position to 0
    fmt.Printf("New lexer created with input: %s\n", input) // Debug print
    // Do not call l.readBinaryString() here
    return l
}

// NextToken reads the next token from the input and returns it.
func (l *Lexer) NextToken() Token {
    var tok Token

    // Skip any whitespace characters to reach the start of the next token.
    l.skipWhitespace()

    if l.position >= len(l.input) {
        tok.Literal = ""
        tok.Type = TOKEN_EOF
        return tok
    }

	binaryString := l.readBinaryString()
    fmt.Printf("Read binary string: %s\n", binaryString) // Debug print

    tokenType := l.determineTokenType(binaryString)
    fmt.Printf("Determined token type: %s\n", tokenType) // Debug print

    // Determine the token type based on the binary string.
    tok = newToken(tokenType, binaryString)

    l.currentToken = tok

    return tok
}

// Placeholder for the determineTokenType method
func (l *Lexer) determineTokenType(binaryString string) TokenType {
    switch binaryString {
	case "00000001": // Binary representation for TOKEN_IDENT
		return TOKEN_IDENT
	case "00000010": // Binary representation for TOKEN_EOF
		return TOKEN_EOF
	case "00000011": // Binary representation for TOKEN_INT
		return TOKEN_INT
	case "00000100": // Binary representation for TOKEN_ILLEGAL
		return TOKEN_ILLEGAL
	case "10000001": // Binary representation for TOKEN_SEMICOLON
		return TOKEN_SEMICOLON
	case "10000010": // Binary representation for TOKEN_PIZZA
		return TOKEN_PIZZA
	case "10000011": // Binary representation for TOKEN_CHEESE
		return TOKEN_CHEESE
	case "10000100": // Binary representation for TOKEN_TACOS
		return TOKEN_TACOS
	case "10000101": // Binary representation for TOKEN_NUGGETS
		return TOKEN_NUGGETS
	case "10000110": // Binary representation for TOKEN_ENCHILADA
		return TOKEN_ENCHILADA
	case "10000111": // Binary representation for TOKEN_APPLE
		return TOKEN_APPLE
	case "10001000": // Binary representation for TOKEN_SALMON
		return TOKEN_SALMON
	case "10001001": // Binary representation for TOKEN_ICACO
		return TOKEN_ICACO
	// ... additional cases if any ...
	default:
		return TOKEN_ILLEGAL
	}
}

func (l *Lexer) readBinaryString() string {
    position := l.position
    if position+8 > len(l.input) {
        return ""
    }
    l.position += 8
    binaryString := l.input[position:l.position]
    fmt.Printf("Binary string read: %s\n", binaryString) // Debug print
    return binaryString
}

// skipWhitespace advances the lexer's position past any whitespace.
func (l *Lexer) skipWhitespace() {
    for l.position < len(l.input) && (l.input[l.position] == ' ' || l.input[l.position] == '\t' || l.input[l.position] == '\n' || l.input[l.position] == '\r') {
        l.position++
    }
}

func newToken(tokenType TokenType, binaryString string) Token {
    return Token{Type: tokenType, Literal: binaryString}
}

// Placeholder functions - implement these based on your language's design
func extractIdentifierFromBinary(binaryString string) string {
    // Logic to extract the identifier from the binary string
    // ...
    return "extracted_identifier"
}

func convertBinaryToInt(binaryString string) string {
    // Logic to convert the binary string to an integer value
    // ...
    return "42" // Example integer value
}

// Example of a new method to read an identifier or number
func (l *Lexer) readToken() string {
    // Read the next binary string, which represents a complete token (like an identifier or a number)
    return l.readBinaryString()
}

// The currentTokenIs method checks if the current token is of a certain type
func (l *Lexer) currentTokenIs(tokenType TokenType) bool {
    return l.currentToken.Type == tokenType
}

// LookupIdent checks if an identifier is a keyword or just a regular identifier.
func LookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		// Mapping of keyword strings to their respective token types.
		"pizza":     TOKEN_PIZZA,
		"cheese":    TOKEN_CHEESE,
		"apple":     TOKEN_APPLE,
		"enchilada": TOKEN_ENCHILADA,
		"icaco":     TOKEN_ICACO,
		"nuggets":   TOKEN_NUGGETS,
		"salmon":    TOKEN_SALMON,
		"tacos":     TOKEN_TACOS,
	}

	if tok, ok := keywords[ident]; ok {
		// If the identifier is a keyword, return the corresponding token type.
		return tok
	}

	// If it's not a keyword, it's a regular identifier.
	return TOKEN_IDENT
}

type Node interface {
	TokenLiteral() string
	String() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out strings.Builder

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Statement interface represents all statement nodes.
type Statement interface {
	Node
	statementNode() // A marker method for statements.
}

// Expression interface represents all expression nodes.
type Expression interface {
	Node
	expressionNode() // A marker method for expressions.
}

type LetStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (ls *LetStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (i *Identifier) String() string {
	return i.Value
}

// Parser struct and methods
type Parser struct {
    lexer     *Lexer
    curToken  Token
    peekToken Token
    errors    []string
}

func NewParser(l *Lexer) *Parser {
    p := &Parser{
        lexer:  l,
        errors: []string{},
    }

    // Read two tokens, so curToken and peekToken are both set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *Program {
    program := &Program{Statements: []Statement{}}

	for p.curToken.Type != TOKEN_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
    switch p.curToken.Type {
    case TOKEN_CHEESE:
        return p.parseLetStatement() // Ensure this matches the method's actual name
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *LetStatement {
    stmt := &LetStatement{Token: p.curToken}

    // Expect a variable name after 'cheese'
    if !p.expectPeek(TOKEN_IDENT) {
        return nil
    }

    stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

    // Expect an '=' sign
    if !p.expectPeek(TOKEN_ENCHILADA) {
        return nil
    }

    // TODO: Parse the expression after '='
    // Skip the expression part for now and just move to the end of the statement
    for !p.curTokenIs(TOKEN_SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) peekError(t TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t.String(), p.peekToken.Type.String())
    p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseExpression(precedence int) Expression {
    var leftExp Expression

    switch p.curToken.Type {
    case TOKEN_INT:
        leftExp = p.parseIntegralLiteral()
    case TOKEN_IDENT:
        leftExp = p.parseIdentifier()
    // Add cases for other types of expressions
    // ...
    default:
        return nil // Or handle unexpected tokens appropriately
    }

    return leftExp
}

func (p *Parser) parseIntegralLiteral() Expression {
	lit := &IntegralLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		fmt.Println("parseIntegralLiteral: Invalid integer")
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

type IntegralLiteral struct {
	Token Token
	Value int64
}

func (il *IntegralLiteral) expressionNode() {}

func (il *IntegralLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegralLiteral) String() string {
	return il.Token.Literal
}

func (t TokenType) String() string {
    return string(t)
}