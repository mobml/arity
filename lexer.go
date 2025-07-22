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
	line   string //raw input
	tokens []Token
}

func NewLexer(input string) Lexer {
	return Lexer{line: input, tokens: make([]Token, 0)}
}

func (l *Lexer) SetLine(input string) {
	l.line = input
}

func (l *Lexer) ScanTokens() {
	source := l.line
	start := 0
	current := 0
	for current < len(source) {
		c := source[current]
		switch c {
		case ' ', '\n', '\r':
			current++
		case '(':
			l.addToken(newToken(string(c), LEFT_PAREN))
			current++
		case ')':
			l.addToken(newToken(string(c), RIGHT_PAREN))
			current++
		case '+':
			l.addToken(newToken(string(c), PLUS))
			current++
		case '-':
			l.addToken(newToken(string(c), MINUS))
			current++
		case '*':
			l.addToken(newToken(string(c), STAR))
			current++
		case '/':
			l.addToken(newToken(string(c), SLASH))
			current++
		default:

			if unicode.IsDigit(rune(c)) {
				start = current
				hasDot := false
				for current < len(source) {
					ch := rune(source[current])
					if ch == '.' {
						if hasDot {
							break
						}
						if current+1 < len(source) && unicode.IsDigit(rune(source[current+1])) {
							hasDot = true
							current++
						} else {
							break
						}
					} else if unicode.IsDigit(ch) {
						current++
					} else {
						break
					}
				}
				lexeme := source[start:current]
				num, err := strconv.ParseFloat(lexeme, 64)
				if err != nil {
					fmt.Printf("Error al convertir número: %s\n", lexeme)
				} else {
					l.addToken(Token{
						Lexeme:  lexeme,
						Type:    NUMBER,
						Literal: num,
					})
				}
			}
		}
	}
}

func (l *Lexer) CleanLexer() {
	l.line = ""
	l.tokens = []Token{}
}

func (l *Lexer) Show() {
	fmt.Printf("source: %s\n", l.line)
	fmt.Printf("tokens: [\n")
	for _, t := range l.tokens {
		fmt.Printf("\t('%s', %s)\n", t.Lexeme, t.Type.String())
	}
	fmt.Printf("]\n")
}

func (t TokenType) String() string {
	switch t {
	case PLUS:
		return "PLUS"
	case NUMBER:
		return "NUMBER"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	}
	return "UNKNOWN"
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}
