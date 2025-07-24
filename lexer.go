package main

import (
	"fmt"
	"strconv"
	"unicode"
)

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

type Lexer struct {
	line    string //raw input
	tokens  []Token
	current int
}

func NewLexer(input string) Lexer {
	return Lexer{line: input, tokens: make([]Token, 0), current: 0}
}

func (l *Lexer) SetLine(input string) {
	l.line = input
}

func (l *Lexer) ScanTokens() error {
	source := l.line
	for l.current < len(source) {
		c := source[l.current]
		switch c {
		case ' ', '\n', '\r':
			l.current++
		case '(':
			l.addToken(newToken(string(c), LEFT_PAREN))
			l.current++
		case ')':
			l.addToken(newToken(string(c), RIGHT_PAREN))
			l.current++
		case '+':
			l.addToken(newToken(string(c), PLUS))
			l.current++
		case '-':
			l.addToken(newToken(string(c), MINUS))
			l.current++
		case '*':
			l.addToken(newToken(string(c), STAR))
			l.current++
		case '/':
			l.addToken(newToken(string(c), SLASH))
			l.current++
		default:
			// Write algorithm to number tokens
			if unicode.IsDigit(rune(c)) {
				if err := l.tokenizeNumber(); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid character: %q", c)
			}
		}
	}
	return nil
}

func underLimit(a, b int) bool {
	return a < b
}

func (l *Lexer) tokenizeNumber() error {
	start := l.current
	source := l.line

	for underLimit(l.current, len(source)) && unicode.IsDigit(rune(source[l.current])) {
		l.current++
	}
	if underLimit(l.current, len(source)) && source[l.current] == '.' {
		if underLimit(l.current+1, len(source)) && unicode.IsDigit(rune(source[l.current+1])) {
			l.current++
			for underLimit(l.current, len(source)) && unicode.IsDigit(rune(source[l.current])) {
				l.current++
			}
			if underLimit(l.current, len(source)) && source[l.current] == '.' {
				return fmt.Errorf("Invalid input: %c", source[l.current])
			}

		} else {
			return fmt.Errorf("Invalid input: %c", source[l.current])
		}
	}

	lexeme := source[start:l.current]
	number, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		return fmt.Errorf("Connot parse: %f", number)
	}
	token := Token{
		Lexeme:  lexeme,
		Type:    NUMBER,
		Literal: number,
	}
	l.addToken(token)
	return nil
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}
