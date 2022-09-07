package parser

import (
	"onemelone/JackAnalyzer/node_type"
	"onemelone/JackAnalyzer/tokenizer"
	"onemelone/JackAnalyzer/util"
)

type Parser struct {
	root           *TreeNode
	curNode        *TreeNode
	inputTokenizer *tokenizer.Tokenizer
}

func NewParser(inputTokenizer *tokenizer.Tokenizer) *Parser {
	return &Parser{
		inputTokenizer: inputTokenizer,
	}
}

// Parse Return the root of the parse-tree.
func (p *Parser) Parse() *TreeNode {
	p.compileClass()
	return p.root
}

func (p *Parser) compileClass() {
	p.root = &TreeNode{
		NodeType: node_type.Class,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}

	// class keyword
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Keyword() != "class" {
		panic("parse error, wrong structure")
	}
	p.root.Sons = append(p.root.Sons, &TreeNode{
		NodeType: node_type.Keyword,
		Value:    util.Str2Ptr(p.inputTokenizer.Keyword()),
		Sons:     nil,
	})

	// class name
	p.inputTokenizer.Advance()
	p.root.Sons = append(p.root.Sons, &TreeNode{
		NodeType: node_type.Identifier,
		Value:    util.Str2Ptr(p.inputTokenizer.Identifier()),
		Sons:     nil,
	})

	// {
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Symbol() != "{" {
		panic("parse error, wrong structure")
	}
	p.root.Sons = append(p.root.Sons, &TreeNode{
		NodeType: node_type.Symbol,
		Value:    util.Str2Ptr(p.inputTokenizer.Symbol()),
		Sons:     nil,
	})

	// set current node
	p.curNode = p.root

	// try to parse class var dec or subroutine dec
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Keyword() == "static" || p.inputTokenizer.Keyword() == "field" {
		p.compileClassVarDec()
	} else if p.inputTokenizer.Keyword() == "constructor" || p.inputTokenizer.Keyword() == "function" ||
		p.inputTokenizer.Keyword() == "method" {
		p.compileSubroutineDec()
	} else {
		panic("parse error, wrong structure")
	}
}

func (p *Parser) compileClassVarDec() {
	classVarDecNode := &TreeNode{
		NodeType: node_type.ClassVarDec,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, classVarDecNode)

	// keyword field / static
	classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
		NodeType: node_type.Keyword,
		Value:    util.Str2Ptr(p.inputTokenizer.Keyword()),
		Sons:     nil,
	})

	// keyword / identifier - int / char / bool / class name
	classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
		NodeType: p.inputTokenizer.TokenType(),
		Value:    util.Str2Ptr(p.inputTokenizer.Token()),
		Sons:     nil,
	})

	// var name
	classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
		NodeType: node_type.Identifier,
		Value:    util.Str2Ptr(p.inputTokenizer.Identifier()),
		Sons:     nil,
	})

	p.inputTokenizer.Advance()

	for p.inputTokenizer.Symbol() == "," {
		// append ","
		classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
			NodeType: node_type.Symbol,
			Value:    util.Str2Ptr(p.inputTokenizer.Symbol()),
			Sons:     nil,
		})

		// append var name
		p.inputTokenizer.Advance()
		classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
			NodeType: node_type.Identifier,
			Value:    util.Str2Ptr(p.inputTokenizer.Identifier()),
			Sons:     nil,
		})

		// get next one
		p.inputTokenizer.Advance()
	}

	// append ";"
	classVarDecNode.Sons = append(classVarDecNode.Sons, &TreeNode{
		NodeType: node_type.Symbol,
		Value:    util.Str2Ptr(p.inputTokenizer.Symbol()),
		Sons:     nil,
	})

	// try to parse class var dec or subroutine dec
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Keyword() == "static" || p.inputTokenizer.Keyword() == "field" {
		p.compileClassVarDec()
	} else if p.inputTokenizer.Keyword() == "constructor" || p.inputTokenizer.Keyword() == "function" ||
		p.inputTokenizer.Keyword() == "method" {
		p.compileSubroutineDec()
	} else {
		panic("parse error, wrong structure")
	}
}

func (p *Parser) compileSubroutineDec() {
}

func (p *Parser) compileVarDec() {
}

func (p *Parser) compileStatements() {
}

func (p *Parser) compileLet() {
}

func (p *Parser) compileIf() {
}

func (p *Parser) compileWhile() {
}

func (p *Parser) compileDo() {
}

func (p *Parser) compileReturn() {
}

func (p *Parser) compileTerm() {
}

func (p *Parser) compileExpressionList() {
}
