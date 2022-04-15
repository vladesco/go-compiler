package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"compiler/token"
	"fmt"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
	errors       []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}
	parser.initialize()

	return parser
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	if !parser.expectCurrentToken(token.EOF) {
		statement := parser.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		parser.readNextToken()
	}

	return program
}

func (parser *Parser) GetParsingErrors() []string {
	return parser.errors
}

func (parser *Parser) readNextToken() {
	parser.currentToken = parser.nextToken
	parser.nextToken = parser.lexer.ReadNextToken()
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:

		return parser.parseLetStatement()
	}

	return nil
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {

	if !parser.readNewNextTokenIfOldExpect(token.IDENT) {
		return nil
	}

	statement := &ast.LetStatement{Token: parser.currentToken}

	if !parser.readNewNextTokenIfOldExpect(token.ASSIGN) {
		return nil
	}

	for !parser.expectCurrentToken(token.SEMICOLON) {
		parser.readNextToken()
	}

	return statement
}

func (parser *Parser) expectCurrentToken(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

func (parser *Parser) readNewNextTokenIfOldExpect(tokenType token.TokenType) bool {

	if parser.nextToken.Type == tokenType {
		parser.readNextToken()
		return true
	}

	parser.writeError(fmt.Sprintf("expected next token to be %s, got %s instead", string(tokenType), string(parser.nextToken.Type)))

	return false
}

func (parser *Parser) writeError(errorMessage string) {
	parser.errors = append(parser.errors, errorMessage)
}

func (parser *Parser) initialize() {
	parser.currentToken = parser.lexer.ReadNextToken()
	parser.nextToken = parser.lexer.ReadNextToken()
}
