package main

import ( //   w   w  w.    bo  o    k 2 s  .    c  o m
	"fmt"
	"unicode"
	"unicode/utf8"
)

// TokenType represents the type of token
type TokenType int

const (
	TokenIdentifier TokenType = iota
	TokenNumber
	TokenOperator
	TokenWhitespace
)

// Token represents a token
type Token struct {
	Type  TokenType
	Value string
}

// Lexer tokenizes input text
type Lexer struct {
	input  string
	pos    int
	tokens []Token
}

// NewLexer creates a new lexer
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		pos:    0,
		tokens: []Token{},
	}
}

// Tokenize tokenizes input text
func (l *Lexer) Tokenize() []Token {
	for l.pos < len(l.input) {
		r, _ := utf8.DecodeRuneInString(l.input[l.pos:])
		switch {
		case unicode.IsSpace(r):
			l.consumeWhitespace()
		case unicode.IsLetter(r):
			l.consumeIdentifier()
		case unicode.IsDigit(r):
			l.consumeNumber()
		default:
			l.consumeOperator()
		}
	}
	return l.tokens
}

// consumeWhitespace consumes whitespace characters
func (l *Lexer) consumeWhitespace() {
	start := l.pos
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}
	l.tokens = append(l.tokens, Token{Type: TokenWhitespace, Value: l.input[start:l.pos]})
}

// consumeIdentifier consumes identifier tokens
func (l *Lexer) consumeIdentifier() {
	start := l.pos
	for l.pos < len(l.input) && (unicode.IsLetter(rune(l.input[l.pos])) || unicode.IsDigit(rune(l.input[l.pos]))) {
		l.pos++
	}
	l.tokens = append(l.tokens, Token{Type: TokenIdentifier, Value: l.input[start:l.pos]})
}

// consumeNumber consumes number tokens
func (l *Lexer) consumeNumber() {
	start := l.pos
	for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
		l.pos++
	}
	l.tokens = append(l.tokens, Token{Type: TokenNumber, Value: l.input[start:l.pos]})
}

// consumeOperator consumes operator tokens
func (l *Lexer) consumeOperator() {
	start := l.pos
	l.pos++
	l.tokens = append(l.tokens, Token{Type: TokenOperator, Value: l.input[start:l.pos]})
}

func main() {
	// Example 2: Implementing a custom language lexer using the `token` package

	// Input text to tokenize
	input := "x = 42 + y"

	// Create a new lexer and tokenize the input
	lexer := NewLexer(input)
	tokens := lexer.Tokenize()

	// Print the tokens
	for _, token := range tokens {
		fmt.Printf("Type: %d, Value: %s\n", token.Type, token.Value)
	}
}
