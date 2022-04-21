package lexer

import (
	"compiler/helpers"
	"compiler/token"
)

type Lexer struct {
	input            string
	cursor           int
	currentCharacter byte
}

func New(input string) *Lexer {
	return &Lexer{input: input, cursor: -1}
}

func (lexer *Lexer) readNextChar() {
	if lexer.cursor+1 >= len(lexer.input) {
		lexer.currentCharacter = 0
	} else {
		lexer.currentCharacter = lexer.input[lexer.cursor+1]
	}
	lexer.cursor += 1

}

func (lexer *Lexer) ReadNextToken() token.Token {
	lexer.skipWhiteSpace()
	lexer.readNextChar()

	var character = lexer.currentCharacter
	var nextToken token.Token

	switch character {
	case ';':
		nextToken = token.New(token.SEMICOLON, ";")
	case '(':
		nextToken = token.New(token.LPAREN, "(")
	case ')':
		nextToken = token.New(token.RPAREN, ")")
	case ',':
		nextToken = token.New(token.COMMA, ",")
	case '+':
		nextToken = token.New(token.PLUS, "+")
	case '-':
		nextToken = token.New(token.MINUS, "-")
	case '*':
		nextToken = token.New(token.ASTERISK, "*")
	case '/':
		nextToken = token.New(token.SLASH, "/")
	case '<':
		nextToken = token.New(token.LT, "<")
	case '>':
		nextToken = token.New(token.GT, ">")
	case '{':
		nextToken = token.New(token.LBRACE, "{")
	case '}':
		nextToken = token.New(token.RBRACE, "}")
	case 0:
		nextToken = token.New(token.EOF, "")

	case '=':
		if lexer.peekChar() == '=' {
			lexer.readNextChar()
			nextToken = token.New(token.EQ, "==")
		} else {
			nextToken = token.New(token.ASSIGN, "=")
		}
	case '!':
		if lexer.peekChar() == '=' {
			lexer.readNextChar()
			nextToken = token.New(token.NOT_EQ, "!=")
		} else {
			nextToken = token.New(token.BANG, "!")
		}

	default:
		if helpers.IsLetter(character) {
			tokenValue := lexer.readTokenValue(helpers.IsLetter)
			nextToken = token.New(token.DefineTokenType((tokenValue)), tokenValue)
		} else if helpers.IsDigit(character) {
			nextToken = token.New(token.INT, lexer.readTokenValue(helpers.IsDigit))
		} else {
			nextToken = token.New(token.ILLEGAL, string(lexer.currentCharacter))
		}
	}

	return nextToken
}

func (lexer *Lexer) readTokenValue(valueFilter func(byte) bool) string {
	stringStartPosition := lexer.cursor

	for valueFilter(lexer.peekChar()) {
		lexer.readNextChar()
	}

	stringEndPosition := lexer.cursor + 1

	return lexer.input[stringStartPosition:stringEndPosition]
}

func (lexer *Lexer) skipWhiteSpace() {
	for helpers.IsWhiteSpace((lexer.peekChar())) {
		lexer.readNextChar()
	}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.cursor+1 >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.cursor+1]
}
