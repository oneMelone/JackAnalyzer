package test

import (
	"io"
	"onemelone/JackAnalyzer/node_type"
	"onemelone/JackAnalyzer/parser"
	"onemelone/JackAnalyzer/tokenizer"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	input, err := os.Open("test_input.jack")
	if err != nil {
		t.Fatalf("Open input file error %+v", err)
	}

	bytes, err := io.ReadAll(input)
	if err != nil {
		t.Fatalf("read input err: %+v", err)
	}

	testTokenizer := tokenizer.NewTokenizer(bytes)

	testParser := parser.NewParser(testTokenizer)

	root := testParser.Parse()

	if root.NodeType != node_type.Class {
		t.Fatalf("root node should be class, but got %d", root.NodeType)
	}

	if root.Sons[0].NodeType != node_type.Keyword {
		t.Fatalf("First son should be keyword, but got %d", root.Sons[0].NodeType)
	}

	if *root.Sons[1].Value != "Test" {
		t.Fatalf("Class name is Test, but got %s", *root.Sons[1].Value)
	}

	if root.Sons[3].NodeType != node_type.ClassVarDec {
		t.Fatalf("Third son should be class var dec, but got %d", root.Sons[3].NodeType)
	}
}
