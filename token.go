package gothon

import "fmt"

type TokenType string

const (
	// Single-character tokens.
	LEFT_PAREN  TokenType = "left_paren"
	RIGHT_PAREN TokenType = "right_paren"
	LEFT_BRACE  TokenType = "left_brace"
	RIGHT_BRACE TokenType = "right_brace"

	COMMA     TokenType = "comma"
	DOT       TokenType = "dot"
	MINUS     TokenType = "minus"
	PLUS      TokenType = "plus"
	SEMICOLON TokenType = "semicolon"
	SLASH     TokenType = "slash"
	STAR      TokenType = "star"
	DIVIDE    TokenType = "divide"

	// One or two character tokens.
	BANG         TokenType = "bang"
	BANG_EQUAL   TokenType = "bang_equal"
	EQUAL        TokenType = "equal"
	EQUAL_EQUAL  TokenType = "equal_equal"
	GREATER      TokenType = "greater"
	GREATER_THAN TokenType = "greater_than"
	LESS         TokenType = "less"
	LESS_EQUAL   TokenType = "less_equal"

	// Literals.
	IDENTIFIER TokenType = "identifier"
	STRING     TokenType = "str"
	NUMBER     TokenType = "num"

	// Keywords.
	FUN   TokenType = "function"
	CLASS TokenType = "class"

	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"

	TRUE  TokenType = "true"
	FALSE TokenType = "false"

	IF    TokenType = "if"
	ELSE  TokenType = "else"
	FOR   TokenType = "for"
	WHILE TokenType = ""

	OR  TokenType = "or"
	AND TokenType = "and"

	VAR TokenType = "var"

	NIL TokenType = "nil"
	EOF TokenType = "eof"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(t TokenType, lexeme string, literal any, line int) Token {
	return Token{
		Type:    t,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("type: %s, lexeme: \"%s\", literal: %v at line: %d", t.Type, t.Lexeme, t.Literal, t.Line)
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
