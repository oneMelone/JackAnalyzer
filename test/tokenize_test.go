package test

import (
	"fmt"
	"io"
	"onemelone/JackAnalyzer/tokenizer"
	"os"
	"testing"
)

func TestTokenizer(t *testing.T) {
	input, err := os.Open("test_input.jack")
	if err != nil {
		t.Fatalf("Open input file error %+v", err)
	}

	bytes, err := io.ReadAll(input)
	if err != nil {
		t.Fatalf("Read error: %+v", err)
	}

	testTokenizer := tokenizer.NewTokenizer(bytes)

	if !testTokenizer.HasMoreTokens() {
		t.Fatalf("Has no tokens at start!")
	}

	testTokenizer.Advance()
	if testTokenizer.Keyword() != "class" {
		t.Fatalf("First token is `class`, but got %s", testTokenizer.Keyword())
	}

	testTokenizer.Advance()
	if testTokenizer.Identifier() != "Test" {
		t.Fatalf("2th token is `Test`, but got %s", testTokenizer.Identifier())
	}

	testTokenizer.Advance()
	if testTokenizer.Symbol() != "{" {
		t.Fatalf("3th token is `{`, but got %s", testTokenizer.Symbol())
	}

	for i := 0; i < 25; i++ {
		testTokenizer.Advance()
	}
	if testTokenizer.StringVal() != "test string" {
		t.Fatalf("28th token is `test string`, but got %s", testTokenizer.StringVal())
	}

	printToken := tokenizer.NewTokenizer(bytes)
	for printToken.Advance() {
		fmt.Println(printToken.Token())
	}
}
