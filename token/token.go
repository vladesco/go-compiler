package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifications and literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// math operators
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// delimiters
	COMMA     = "COMMA"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func New(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}

func DefineTokenType(value string) TokenType {
	if tokenType, exist := keywords[value]; exist {
		return tokenType
	}
	return IDENT

}
