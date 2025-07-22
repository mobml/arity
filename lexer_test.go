package main

import (
	"reflect"
	"testing"
)

func TestSimpleLexer(t *testing.T) {
	input := "1 + 2"
	lexer := NewLexer(input)
	lexer.ScanTokens()

	expected := []Token{
		{Lexeme: "1", Type: NUMBER, Literal: 1.0},
		{Lexeme: "+", Type: PLUS, Literal: nil},
		{Lexeme: "2", Type: NUMBER, Literal: 2.0},
	}

	if len(lexer.tokens) != len(expected) {
		t.Fatalf("Expected %d token, got %d", len(expected), len(lexer.tokens))
	}

	for i, token := range expected {
		got := lexer.tokens[i]

		if token.Lexeme != got.Lexeme || token.Type != got.Type || !reflect.DeepEqual(token.Literal, got.Literal) {
			t.Errorf("Token %d mismatch. Expected %+v, got %+v", i, token, got)
		}
	}

}

func TestDecimalNumberLexer(t *testing.T) {
	input := "(1.2) * (5.5) + (2 - 5) / 10"
	lexer := NewLexer(input)
	lexer.ScanTokens()

	expected := []Token{
		{Lexeme: "(", Type: LEFT_PAREN, Literal: nil},
		{Lexeme: "1.2", Type: NUMBER, Literal: 1.2},
		{Lexeme: ")", Type: RIGHT_PAREN, Literal: nil},
		{Lexeme: "*", Type: STAR, Literal: nil},
		{Lexeme: "(", Type: LEFT_PAREN, Literal: nil},
		{Lexeme: "5.5", Type: NUMBER, Literal: 5.5},
		{Lexeme: ")", Type: RIGHT_PAREN, Literal: nil},
		{Lexeme: "+", Type: PLUS, Literal: nil},
		{Lexeme: "(", Type: LEFT_PAREN, Literal: nil},
		{Lexeme: "2", Type: NUMBER, Literal: 2.0},
		{Lexeme: "-", Type: MINUS, Literal: nil},
		{Lexeme: "5", Type: NUMBER, Literal: 5.0},
		{Lexeme: ")", Type: RIGHT_PAREN, Literal: nil},
		{Lexeme: "/", Type: SLASH, Literal: nil},
		{Lexeme: "10", Type: NUMBER, Literal: 10.0},
	}

	if len(lexer.tokens) != len(expected) {
		t.Fatalf("Expected %d token, got %d", len(expected), len(lexer.tokens))
	}

	for i, token := range expected {
		got := lexer.tokens[i]

		if token.Lexeme != got.Lexeme || token.Type != got.Type || !reflect.DeepEqual(token.Literal, got.Literal) {
			t.Errorf("Token %d mismatch. Expected %+v, got %+v", i, token, got)
		}
	}

}

func TestDotAtStart_Invalid(t *testing.T) {
	input := ".1"
	lexer := NewLexer(input)
	lexer.ScanTokens()

	if len(lexer.tokens) != 0 {
		t.Fatalf("Expected no tokens for invalid input, got %d", len(lexer.tokens))
	}
}

func TestDotAtEnd_Invalid(t *testing.T) {
	input := "1."
	lexer := NewLexer(input)
	lexer.ScanTokens()

	if len(lexer.tokens) != 0 {
		t.Fatalf("Expected no tokens for invalid input, got %d", len(lexer.tokens))
	}
}
