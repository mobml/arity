package main

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	STAR
	SLASH
	LEFT_PAREN
	RIGHT_PAREN
)

type Token struct {
	Lexeme  string
	Type    TokenType
	Literal any
}

func newToken(lex string, typeToken TokenType) Token {
	return Token{Lexeme: lex, Type: typeToken, Literal: nil}
}
