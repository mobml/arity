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
	//start := 0
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
			// Write algorithm to number tokens
		}
	}
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}
