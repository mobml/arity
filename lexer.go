package main

import (
	"fmt"
	"strconv"
	"unicode"
)

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
			l.advance()
		case '(':
			l.addToken(newToken(string(c), LEFT_PAREN))
			l.advance()
		case ')':
			l.addToken(newToken(string(c), RIGHT_PAREN))
			l.advance()
		case '+':
			l.addToken(newToken(string(c), PLUS))
			l.advance()
		case '-':
			l.addToken(newToken(string(c), MINUS))
			l.advance()
		case '*':
			l.addToken(newToken(string(c), STAR))
			l.advance()
		case '/':
			l.addToken(newToken(string(c), SLASH))
			l.advance()
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

func (l *Lexer) peek() rune {
	if l.current < len(l.line) {
		return rune(l.line[l.current])
	}
	return 0
}

func (l *Lexer) peekNext() rune {
	if l.current+1 < len(l.line) {
		return rune(l.line[l.current+1])
	}
	return 0
}

func (l *Lexer) hasNext() bool {
	return l.current < len(l.line)
}

func (l *Lexer) advance() {
	l.current++
}

func (l *Lexer) isValidDecimalStart() bool {
	return l.peek() == '.' && unicode.IsDigit(l.peekNext())
}

func (l *Lexer) consumeDigits() {
	for l.hasNext() && unicode.IsDigit(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) tokenizeNumber() error {
	start := l.current

	l.consumeDigits()

	if l.isValidDecimalStart() {
		l.advance()
		l.consumeDigits()
		if l.peek() == '.' {
			return fmt.Errorf("Invalid input: multiple dots")
		}
	} else if l.peek() == '.' {
		return fmt.Errorf("Invalid input: dot not followed by digit")
	}

	lexeme := l.line[start:l.current]
	number, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		return fmt.Errorf("Can not parse: %f", number)
	}
	l.addToken(Token{
		Lexeme:  lexeme,
		Type:    NUMBER,
		Literal: number,
	})
	return nil
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}
