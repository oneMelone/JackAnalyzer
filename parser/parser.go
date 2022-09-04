package parser

import "onemelone/JackAnalyzer/tokenizer"

type Parser struct {
	inputTokenizer *tokenizer.Tokenizer
}

func NewParser(inputTokenizer *tokenizer.Tokenizer) *Parser {
	return &Parser{
		inputTokenizer: inputTokenizer,
	}
}

// Parse Return the root of the parse-tree.
func (p *Parser) Parse() *TreeNode {
	return nil
}

func (p *Parser) compileClass() *TreeNode {
	return nil
}

func (p *Parser) compileClassVarDec() *TreeNode {
	return nil
}

func (p *Parser) compileSubroutineDec() *TreeNode {
	return nil
}

func (p *Parser) compileVarDec() *TreeNode {
	return nil
}

func (p *Parser) compileStatements() *TreeNode {
	return nil
}

func (p *Parser) compileLet() *TreeNode {
	return nil
}

func (p *Parser) compileIf() *TreeNode {
	return nil
}

func (p *Parser) compileWhile() *TreeNode {
	return nil
}

func (p *Parser) compileDo() *TreeNode {
	return nil
}

func (p *Parser) compileReturn() *TreeNode {
	return nil
}

func (p *Parser) compileTerm() *TreeNode {
	return nil
}

func (p *Parser) compileExpressionList() *TreeNode {
	return nil
}
