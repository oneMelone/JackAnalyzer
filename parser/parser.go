package parser

import (
	"onemelone/JackAnalyzer/node_type"
	"onemelone/JackAnalyzer/tokenizer"
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
	p.parseKeyword(p.root)

	// class name
	p.inputTokenizer.Advance()
	p.parseIdentifier(p.root)

	// {
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Symbol() != "{" {
		panic("parse error, wrong structure")
	}
	p.parseSymbol(p.root)

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

	// }
	p.parseSymbol(p.root)
}

func (p *Parser) compileClassVarDec() {
	classVarDecNode := &TreeNode{
		NodeType: node_type.ClassVarDec,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, classVarDecNode)

	// keyword field / static
	p.parseKeyword(classVarDecNode)
	// keyword / identifier - int / char / bool / class name
	p.inputTokenizer.Advance()
	p.parseGeneral(classVarDecNode)
	// var name
	p.inputTokenizer.Advance()
	p.parseIdentifier(classVarDecNode)

	p.inputTokenizer.Advance()
	for p.inputTokenizer.Symbol() == "," {
		// append ","
		p.parseSymbol(classVarDecNode)
		// append var name
		p.inputTokenizer.Advance()
		p.parseIdentifier(classVarDecNode)
		// get next one
		p.inputTokenizer.Advance()
	}

	// append ";"
	p.parseSymbol(classVarDecNode)
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
	subroutineDecNode := &TreeNode{
		NodeType: node_type.SubroutineDec,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, subroutineDecNode)
	storeCurNode := p.curNode
	p.curNode = subroutineDecNode

	// keyword constructor / function method
	p.parseKeyword(subroutineDecNode)
	// type
	p.inputTokenizer.Advance()
	p.parseGeneral(subroutineDecNode)
	// subroutine name
	p.inputTokenizer.Advance()
	p.parseIdentifier(subroutineDecNode)
	// left (
	p.inputTokenizer.Advance()
	p.parseSymbol(subroutineDecNode)
	// parameter list
	p.inputTokenizer.Advance()
	p.compileParamList()
	// right )
	p.parseSymbol(subroutineDecNode)

	p.compileSubroutineBody()

	p.curNode = storeCurNode
	if p.inputTokenizer.Token() != "}" {
		p.compileSubroutineDec()
	}
}

func (p *Parser) compileParamList() {
	paramListNode := &TreeNode{
		NodeType: node_type.ParameterList,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, paramListNode)

	for p.inputTokenizer.Token() != ")" {
		p.parseGeneral(paramListNode)
		p.inputTokenizer.Advance()
	}
}

func (p *Parser) compileSubroutineBody() {
	subroutineBodyNode := &TreeNode{
		NodeType: node_type.SubroutineBody,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, subroutineBodyNode)
	p.curNode = subroutineBodyNode

	p.inputTokenizer.Advance()
	if p.inputTokenizer.Symbol() != "{" {
		panic("parse error, wrong structure")
	}
	p.parseSymbol(subroutineBodyNode)

	p.inputTokenizer.Advance()
	for true {
		if p.inputTokenizer.Token() == "}" {
			p.parseSymbol(subroutineBodyNode)
			return
		}

		if p.inputTokenizer.Keyword() == "var" {
			p.compileVarDec()
		} else {
			p.compileStatements()
		}
	}
}

func (p *Parser) compileVarDec() {
	varDecNode := &TreeNode{
		NodeType: node_type.VarDec,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, varDecNode)

	// var
	p.parseKeyword(varDecNode)
	// type
	p.inputTokenizer.Advance()
	p.parseGeneral(varDecNode)
	// varName
	p.inputTokenizer.Advance()
	p.parseIdentifier(varDecNode)

	for p.inputTokenizer.Advance() && p.inputTokenizer.Symbol() != ";" {
		// ,
		p.parseSymbol(varDecNode)
		// varName
		p.parseIdentifier(varDecNode)
	}

	// ;
	p.parseSymbol(varDecNode)
	p.inputTokenizer.Advance()
}

func (p *Parser) compileStatements() {
	statementsNode := &TreeNode{
		NodeType: node_type.Statements,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, statementsNode)
	p.curNode = statementsNode
	p.branchStatementParse()
}

func (p *Parser) branchStatementParse() {
	switch p.inputTokenizer.Token() {
	case "let":
		p.compileLet()
	case "if":
		p.compileIf()
	case "while":
		p.compileWhile()
	case "do":
		p.compileDo()
	case "return":
		p.compileReturn()
	case "}":
		return
	}
}

func (p *Parser) compileLet() {
	letNode := &TreeNode{
		NodeType: node_type.LetStatement,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, letNode)
	curNodeStore := p.curNode
	p.curNode = letNode

	// let
	p.parseKeyword(letNode)
	// varName
	p.inputTokenizer.Advance()
	p.parseIdentifier(letNode)
	// [ / =
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Symbol() == "[" {
		// array element access
		// [
		p.parseSymbol(letNode)
		// expression
		p.inputTokenizer.Advance()
		p.compileExpression()
		// ]
		p.inputTokenizer.Advance()
		p.parseSymbol(letNode)

		p.inputTokenizer.Advance()
	}

	// =
	p.parseSymbol(letNode)

	// expression
	p.inputTokenizer.Advance()
	p.compileExpression()

	// ;
	// p.inputTokenizer.Advance()
	p.parseSymbol(letNode)

	p.curNode = curNodeStore
	if p.inputTokenizer.Advance() {
		p.branchStatementParse()
	}
}

func (p *Parser) compileIf() {
	ifNode := &TreeNode{
		NodeType: node_type.IfStatement,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, ifNode)
	curNodeStore := p.curNode
	p.curNode = ifNode

	// if
	p.parseKeyword(ifNode)
	// (
	p.inputTokenizer.Advance()
	p.parseSymbol(ifNode)
	// expression
	p.inputTokenizer.Advance()
	p.compileExpression()
	// )
	p.inputTokenizer.Advance()
	p.parseSymbol(ifNode)
	// {
	p.inputTokenizer.Advance()
	p.parseSymbol(ifNode)
	// statements
	p.inputTokenizer.Advance()
	p.compileStatements()
	// }
	p.inputTokenizer.Advance()
	p.parseSymbol(ifNode)

	if p.inputTokenizer.Advance() && p.inputTokenizer.Keyword() == "else" {
		// else
		p.parseKeyword(ifNode)
		// {
		p.inputTokenizer.Advance()
		p.parseSymbol(ifNode)
		// statements
		p.inputTokenizer.Advance()
		p.compileStatements()
		// }
		p.inputTokenizer.Advance()
		p.parseSymbol(ifNode)
	}

	p.curNode = curNodeStore
	if p.inputTokenizer.Advance() {
		p.branchStatementParse()
	}
}

func (p *Parser) compileWhile() {
	whileNode := &TreeNode{
		NodeType: node_type.WhileStatement,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, whileNode)
	curNodeStore := p.curNode
	p.curNode = whileNode

	// while
	p.parseKeyword(whileNode)
	// (
	p.inputTokenizer.Advance()
	p.parseSymbol(whileNode)
	// expression
	p.inputTokenizer.Advance()
	p.compileExpression()
	// )
	p.inputTokenizer.Advance()
	p.parseSymbol(whileNode)
	// {
	p.inputTokenizer.Advance()
	p.parseSymbol(whileNode)
	// statements
	p.inputTokenizer.Advance()
	p.compileStatements()
	// }
	p.inputTokenizer.Advance()
	p.parseSymbol(whileNode)

	p.curNode = curNodeStore
	if p.inputTokenizer.Advance() {
		p.branchStatementParse()
	}
}

func (p *Parser) compileDo() {
	doNode := &TreeNode{
		NodeType: node_type.DoStatement,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, doNode)
	curNodeStore := p.curNode
	p.curNode = doNode

	// do
	p.parseKeyword(doNode)
	// subroutine name / class name / var name
	p.inputTokenizer.Advance()
	p.parseIdentifier(doNode)

	p.inputTokenizer.Advance()
	if p.inputTokenizer.Symbol() == "." {
		// x.method()
		p.parseSymbol(doNode)
		// subroutine name
		p.inputTokenizer.Advance()
		p.parseIdentifier(doNode)

		p.inputTokenizer.Advance()
	}

	// (expressionList)
	p.parseSymbol(doNode)
	p.inputTokenizer.Advance()
	p.compileExpressionList()
	p.inputTokenizer.Advance()
	p.parseSymbol(doNode)
	// ;
	p.inputTokenizer.Advance()
	p.parseSymbol(doNode)

	p.curNode = curNodeStore

	if p.inputTokenizer.Advance() {
		p.branchStatementParse()
	}
}

func (p *Parser) compileReturn() {
	returnNode := &TreeNode{
		NodeType: node_type.ReturnStatement,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, returnNode)
	curNodeStore := p.curNode
	p.curNode = returnNode

	// return
	p.parseKeyword(returnNode)

	// expression ?
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Token() == ";" {
		p.parseSymbol(returnNode)
	} else {
		p.compileExpression()
	}
	// ;
	p.parseSymbol(returnNode)

	p.curNode = curNodeStore
	if p.inputTokenizer.Advance() {
		p.branchStatementParse()
	}
}

func (p *Parser) compileTerm() {
	termNode := &TreeNode{
		NodeType: node_type.Term,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, termNode)
	curNodeStore := p.curNode
	p.curNode = termNode

	if p.inputTokenizer.TokenType() == node_type.IntegerConstant {
		p.parseIntConst(termNode)
		p.curNode = curNodeStore
		p.inputTokenizer.Advance()
		return
	}

	if p.inputTokenizer.TokenType() == node_type.StringConstant {
		p.parseStrConst(termNode)
		p.curNode = curNodeStore
		p.inputTokenizer.Advance()
		return
	}

	if p.inputTokenizer.TokenType() == node_type.Keyword {
		p.parseKeyword(termNode)
		p.curNode = curNodeStore
		p.inputTokenizer.Advance()
		return
	}

	if p.inputTokenizer.TokenType() == node_type.Symbol {
		if p.inputTokenizer.Symbol() == "(" {
			// (expression)
			p.parseSymbol(termNode)
			p.inputTokenizer.Advance()
			p.compileExpression()
			p.inputTokenizer.Advance()
			p.parseSymbol(termNode)
			p.inputTokenizer.Advance()
			p.curNode = curNodeStore
			return
		}
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.compileTerm()
		p.inputTokenizer.Advance()
		p.curNode = curNodeStore
		return
	}

	p.parseIdentifier(termNode)
	p.inputTokenizer.Advance()
	if p.inputTokenizer.Token() == "[" {
		// access array
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.compileExpression()
		p.inputTokenizer.Advance()
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.curNode = curNodeStore
		return
	}

	if p.inputTokenizer.Token() == "(" {
		// subroutine call
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.compileExpressionList()
		p.inputTokenizer.Advance()
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.curNode = curNodeStore
		return
	}

	if p.inputTokenizer.Token() == "." {
		// method call
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.parseIdentifier(termNode)
		p.inputTokenizer.Advance()
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.compileExpressionList()
		p.inputTokenizer.Advance()
		p.parseSymbol(termNode)
		p.inputTokenizer.Advance()
		p.curNode = curNodeStore
		return
	}

	p.inputTokenizer.Advance()
	p.curNode = curNodeStore
	return
}

func (p *Parser) compileExpressionList() {
	expressionListNode := &TreeNode{
		NodeType: node_type.ExpressionList,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, expressionListNode)
	curNodeStore := p.curNode
	p.curNode = expressionListNode

	p.compileExpression()
	for p.inputTokenizer.Advance() && p.inputTokenizer.Token() == "," {
		p.parseSymbol(expressionListNode)
		p.inputTokenizer.Advance()
		p.compileExpression()
	}

	p.curNode = curNodeStore
}

func (p *Parser) compileExpression() {
	expressionNode := &TreeNode{
		NodeType: node_type.Expression,
		Value:    nil,
		Sons:     make([]*TreeNode, 0),
	}
	p.curNode.Sons = append(p.curNode.Sons, expressionNode)
	curNodeStore := p.curNode
	p.curNode = expressionNode

	p.compileTerm()

	var opSet = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {},
		"&": {}, "|": {}, "<": {}, "=": {},
	}
	if _, ok := opSet[p.inputTokenizer.Token()]; ok {
		// op term
		p.parseSymbol(expressionNode)
		p.inputTokenizer.Advance()
		p.compileTerm()
	}

	p.curNode = curNodeStore
}
