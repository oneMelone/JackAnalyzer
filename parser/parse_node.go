package parser

import (
	"onemelone/JackAnalyzer/node_type"
	"onemelone/JackAnalyzer/util"
	"strconv"
)

func (p *Parser) parseKeyword(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: node_type.Keyword,
		Value:    util.Str2Ptr(p.inputTokenizer.Keyword()),
		Sons:     nil,
	})
}

func (p *Parser) parseSymbol(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: node_type.Symbol,
		Value:    util.Str2Ptr(p.inputTokenizer.Symbol()),
		Sons:     nil,
	})
}

func (p *Parser) parseIntConst(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: node_type.IntegerConstant,
		Value:    util.Str2Ptr(strconv.Itoa(p.inputTokenizer.IntVal())),
		Sons:     nil,
	})
}

func (p *Parser) parseStrConst(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: node_type.StringConstant,
		Value:    util.Str2Ptr(p.inputTokenizer.StringVal()),
		Sons:     nil,
	})
}

func (p *Parser) parseIdentifier(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: node_type.Identifier,
		Value:    util.Str2Ptr(p.inputTokenizer.Identifier()),
		Sons:     nil,
	})
}

func (p *Parser) parseGeneral(node *TreeNode) {
	node.Sons = append(node.Sons, &TreeNode{
		NodeType: p.inputTokenizer.TokenType(),
		Value:    util.Str2Ptr(p.inputTokenizer.Token()),
		Sons:     nil,
	})
}
