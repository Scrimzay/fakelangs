package main

import (
	"strings"
	"fmt"
	"strconv"
)

type TokenType int

// Pizza for printing, Cheese for var declaration, Tacos for identifying, Nuggets for number
// Enchilada for equal, Apple for add, Salmon for subtract, Icaco for inputs

const (
	TOKEN_IDENT TokenType = iota
	TOKEN_EOF
	TOKEN_INT
	TOKEN_ILLEGAL
	TOKEN_SEMICOLON
	TOKEN_PIZZA
	TOKEN_CHEESE
	TOKEN_TACOS
	TOKEN_NUGGETS
	TOKEN_ENCHILADA
	TOKEN_APPLE
	TOKEN_SALMON
	TOKEN_ICACO
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
	readPosition int  // current reading position in input
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Check if we reached the end of the input
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL" character, signifies end of input
		l.ch = 0
	} else {
		//Read the next char
		l.ch = l.input[l.readPosition]
	}
	// Update the current position
	l.position = l.readPosition
	//Move to next position
	l.readPosition++
}

// NextToken reads the next token from the input and returns it.
func (l *Lexer) NextToken() Token {
	var tok Token

	// Skip any whitespace characters to reach the start of the next token.
	l.skipWhitespace()

	// Switch statement to handle different characters.
	switch l.ch {
	case '=':
		// If the current character is '=', create an ENCHILADA token.
		tok = newToken(TOKEN_ENCHILADA, l.ch)
	case '+':
		// If it's '+', create an APPLE token.
		tok = newToken(TOKEN_APPLE, l.ch)
	case '-':
		// If it's '-', create a SALMON token.
		tok = newToken(TOKEN_SALMON, l.ch)
    case ';':
        tok = newToken(TOKEN_SEMICOLON, l.ch)
	case 0:
		// If it's the end of the input (0), create an EOF (End Of File) token.
		tok.Literal = ""
		tok.Type = TOKEN_EOF
	default:
		// If it's not a known character, check if it's an identifier or number.
		if isLetter(l.ch) {
			// If it's a letter, read the full identifier and check if it's a keyword.
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// If it's a digit, read the full number.
			tok.Literal = l.readNumber()
			tok.Type = TOKEN_INT
			return tok
		} else {
			// If it's an unknown character, create an ILLEGAL token.
			tok = newToken(TOKEN_ILLEGAL, l.ch)
		}
	}

	// Read the next character for the next call to NextToken.
	l.readChar()
	return tok
}

// skipWhitespace advances the lexer's position past any whitespace.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken creates a new Token of the given TokenType and literal character.
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads an identifier from the current position.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	// There's a typo here; it should be l.position, not 1.position
	return l.input[position:l.position]
}

// isLetter checks if the character is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readNumber reads a full number starting from the current position.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	// Again, it should be l.position, not 1.position
	return l.input[position:l.position]
}

// isDigit checks if the character is a numeric digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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

// Node is the base interface for all nodes in the abstract syntax tree (AST).
// Each node must implement these methods.
type Node interface {
    TokenLiteral() string // Returns the literal value of the token associated with this node.
    String() string       // Returns a string representation of the node.
}

// Program is the root node of every AST produced by the parser.
type Program struct {
    Statements []Statement // A sequence of statements in the program.
}

// TokenLiteral returns the literal value of the first token in the program,
// or an empty string if there are no statements.
func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

// String returns a concatenated string of all the statement strings in the program.
func (p *Program) String() string {
    var out strings.Builder
    for _, s := range p.Statements {
        out.WriteString(s.String())
    }
    return out.String()
}

// Statement interface represents all statement nodes in the AST.
type Statement interface {
    Node
    statementNode() // A marker method to differentiate other nodes from statement nodes.
}

// Expression interface represents all expression nodes in the AST.
type Expression interface {
    Node
    expressionNode() // A marker method to differentiate other nodes from expression nodes.
}

// LetStatement represents a variable declaration (e.g., "cheese x = 5;").
type LetStatement struct {
    Token Token     // The first token of the statement (TOKEN_CHEESE in this case).
    Name  *Identifier // The variable name being declared.
    Value Expression // The expression assigned to the variable.
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

// Identifier represents a variable name in the AST.
type Identifier struct {
    Token Token  // The token (TOKEN_IDENT) associated with the identifier.
    Value string // The name of the identifier.
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}

// String representation of a LetStatement (e.g., "cheese x = 5;").
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

// String representation of an Identifier.
func (i *Identifier) String() string {
    return i.Value
}

// Parser struct contains the state of the parser, including a lexer, current tokens, and errors.
type Parser struct {
    lexer     *Lexer
    curToken  Token
    peekToken Token
    errors    []string // A slice of errors encountered during parsing.
}

// NewParser creates a new Parser instance using a Lexer.
func NewParser(l *Lexer) *Parser {
    p := &Parser{
        lexer:  l,
        errors: []string{},
    }

    // Initialize curToken and peekToken.
    p.nextToken()
    p.nextToken()

    return p
}

// nextToken advances the tokens: current token becomes the peek token, and peek token is updated.
func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

// ParseProgram parses the entire program and returns the root node (Program).
func (p *Parser) ParseProgram() *Program {
    program := &Program{Statements: []Statement{}}

    // Loop through all tokens until EOF, parsing each statement.
    for p.curToken.Type != TOKEN_EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

// parseStatement directs the parsing of different types of statements based on the current token.
func (p *Parser) parseStatement() Statement {
    switch p.curToken.Type {
    case TOKEN_CHEESE:
        return p.parseLetStatement()
    // Add more cases for other types of statements.
    default:
        return nil
    }
}

// parseLetStatement parses a let statement (variable declaration).
func (p *Parser) parseLetStatement() *LetStatement {
    stmt := &LetStatement{Token: p.curToken}

    if !p.expectPeek(TOKEN_IDENT) {
        return nil
    }

    stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(TOKEN_ENCHILADA) {
        return nil
    }

    // Parsing the expression after "=" and before ";".
    p.nextToken()
    stmt.Value = p.parseExpression(LOWEST)

    // Skipping to the end of the statement (semicolon).
    for !p.curTokenIs(TOKEN_SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

// curTokenIs checks if the current token is of a given type.
func (p *Parser) curTokenIs(t TokenType) bool {
    return p.curToken.Type == t
}

// peekTokenIs checks if the next token is of a given type.
func (p *Parser) peekTokenIs(t TokenType) bool {
    return p.peekToken.Type == t
}

// peekError appends an error message when the next token is not of the expected type.
func (p *Parser) peekError(t TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t.String(), p.peekToken.Type.String())
    p.errors = append(p.errors, msg)
}

// expectPeek checks if the next token is of the expected type and advances the tokens if true.
func (p *Parser) expectPeek(t TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peekError(t)
        return false
    }
}

// Errors returns the list of parsing errors encountered.
func (p *Parser) Errors() []string {
    return p.errors
}

// parseExpression handles the parsing of expressions, with precedence taken into account.
func (p *Parser) parseExpression(precedence int) Expression {
    var leftExp Expression

    // Add cases here to handle different types of expressions based on the current token.
    switch p.curToken.Type {
    case TOKEN_INT:
        leftExp = p.parseIntegerLiteral()
    case TOKEN_IDENT:
        leftExp = p.parseIdentifier()
    default:
        return nil
    }

    return leftExp
}

// parseIntegerLiteral handles parsing of integer literals.
func (p *Parser) parseIntegerLiteral() Expression {
    lit := &IntegralLiteral{Token: p.curToken}

    value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
    if err != nil {
        fmt.Println("parseIntegerLiteral: Invalid integer")
        return nil
    }

    lit.Value = value
    return lit
}

// parseIdentifier handles parsing of identifiers.
func (p *Parser) parseIdentifier() Expression {
    return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// IntegralLiteral represents an integer literal in the AST.
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

// String method for TokenType returns a string representation of the TokenType.
func (t TokenType) String() string {
    switch t {
    case TOKEN_IDENT:
        return "TOKEN_IDENT"
    case TOKEN_EOF:
        return "TOKEN_EOF"
    // ... add cases for other token types ...
    default:
        return fmt.Sprintf("Unknown TokenType (%d)", int(t))
    }
}

