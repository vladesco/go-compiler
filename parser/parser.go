package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"compiler/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

type Parser struct {
	lexer          *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	errors         []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}
	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.infixParseFns = make(map[token.TokenType]infixParseFn)

	parser.initialize()

	return parser
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parser.expectCurrentToken(token.EOF) {

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

func (parser *Parser) registerPrefixParseFn(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

func (parser *Parser) registerInfixParseFn(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

func (parser *Parser) readNextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.ReadNextToken()
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()

	case token.RETURN:
		return parser.parseReturnStatement()

	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{BaseNode: ast.BaseNode{Token: parser.currentToken}}

	if !parser.readNextTokenIfPeekExpect(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		BaseNode: ast.BaseNode{Token: parser.currentToken},
		Value:    parser.currentToken.Literal,
	}

	if !parser.readNextTokenIfPeekExpect(token.ASSIGN) {
		return nil
	}

	for !parser.expectCurrentToken(token.SEMICOLON) {
		parser.readNextToken()
	}

	return statement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{BaseNode: ast.BaseNode{Token: parser.currentToken}}

	parser.readNextToken()

	for !parser.expectCurrentToken(token.SEMICOLON) {
		parser.readNextToken()
	}

	return statement
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		BaseNode:   ast.BaseNode{Token: parser.currentToken},
		Expression: parser.parseExpression(LOWEST),
	}

	if parser.expectPeekToken(token.SEMICOLON) {
		parser.readNextToken()
	}

	return statement
}

func (parser *Parser) parseExpression(precendance int) ast.Expression {
	prefixFn := parser.prefixParseFns[parser.currentToken.Type]

	if prefixFn == nil {
		parser.writeError(fmt.Sprintf("no prefix parse function for %s found", parser.currentToken.Literal))
		return nil
	}

	leftExp := prefixFn()

	for !parser.readNextTokenIfPeekExpect(token.SEMICOLON) && precendance < parser.peekTokenPrecedence() {
		infixFn := parser.infixParseFns[parser.peekToken.Type]

		if infixFn == nil {
			return leftExp
		}

		parser.readNextToken()
		leftExp = infixFn(leftExp)
	}

	return leftExp
}

func (parser *Parser) parseIdentifier() ast.Expression {
	expression := &ast.Identifier{
		BaseNode: ast.BaseNode{Token: parser.currentToken},
		Value:    parser.currentToken.Literal,
	}

	return expression
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{BaseNode: ast.BaseNode{Token: parser.currentToken}}

	value, err := strconv.ParseInt(parser.currentToken.Literal, 0, 64)

	if err != nil {
		parser.writeError(fmt.Sprintf("could not parse %q as integer", parser.currentToken.Literal))
		return nil
	}

	literal.Value = value

	return literal
}

func (parser *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		BaseNode: ast.BaseNode{Token: parser.currentToken},
		Operator: parser.currentToken.Literal,
	}

	parser.readNextToken()
	expression.Right = parser.parseExpression(PREFIX)

	return expression
}

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		BaseNode: ast.BaseNode{Token: parser.currentToken},
		Left:     left,
		Operator: parser.currentToken.Literal,
	}

	precedence := parser.currentTokenPrecedence()
	parser.readNextToken()
	expression.Right = parser.parseExpression(precedence)

	return expression
}

func (parser *Parser) expectCurrentToken(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

func (parser *Parser) currentTokenPrecedence() int {
	if peek, ok := precedences[parser.currentToken.Type]; ok {
		return peek
	}

	return LOWEST
}

func (parser *Parser) expectPeekToken(tokenType token.TokenType) bool {
	return parser.peekToken.Type == tokenType
}

func (parser *Parser) peekTokenPrecedence() int {
	if peek, ok := precedences[parser.peekToken.Type]; ok {
		return peek
	}

	return LOWEST
}

func (parser *Parser) readNextTokenIfPeekExpect(tokenType token.TokenType) bool {

	if parser.peekToken.Type == tokenType {
		parser.readNextToken()
		return true
	}

	parser.writeError(fmt.Sprintf("expected next token to be %s, got %s instead", string(tokenType), string(parser.peekToken.Type)))

	return false
}

func (parser *Parser) writeError(errorMessage string) {
	parser.errors = append(parser.errors, errorMessage)
}

func (parser *Parser) initialize() {
	parser.currentToken = parser.lexer.ReadNextToken()
	parser.peekToken = parser.lexer.ReadNextToken()

	parser.registerPrefixParseFn(token.IDENT, parser.parseIdentifier)
	parser.registerPrefixParseFn(token.INT, parser.parseIntegerLiteral)
	parser.registerPrefixParseFn(token.BANG, parser.parsePrefixExpression)
	parser.registerPrefixParseFn(token.MINUS, parser.parsePrefixExpression)

	parser.registerInfixParseFn(token.PLUS, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.MINUS, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.SLASH, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.ASTERISK, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.EQ, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.NOT_EQ, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.LT, parser.parseInfixExpression)
	parser.registerInfixParseFn(token.GT, parser.parseInfixExpression)
}
