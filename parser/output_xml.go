package parser

import (
	"bufio"
	"onemelone/JackAnalyzer/node_type"
)

// OutputXML Write xml format file, with provided root node.
func OutputXML(writer *bufio.Writer, root *TreeNode, indent int) {
	for i := 0; i < 4*indent; i++ {
		_, err := writer.WriteString(" ")
		if err != nil {
			panic("write error")
		}
	}

	start, end := genTagByType(root.NodeType)
	_, err := writer.WriteString(start)
	if err != nil {
		panic("write error")
	}

	if root.Value != nil {
		// leaf
		_, err = writer.WriteString(" " + *root.Value + " ")
		if err != nil {
			panic("write error")
		}
	} else {
		// not leaf
		_, err = writer.WriteString("\n")
		if err != nil {
			panic("write error")
		}

		for _, node := range root.Sons {
			OutputXML(writer, node, indent+1)
			_, err = writer.WriteString("\n")
			if err != nil {
				panic("write error")
			}
		}

		for i := 0; i < 4*indent; i++ {
			_, err := writer.WriteString(" ")
			if err != nil {
				panic("write error")
			}
		}
	}

	_, err = writer.WriteString(end)
	if err != nil {
		panic("write error")
	}

	err = writer.Flush()
	if err != nil {
		panic("flush error")
	}
}

func genTagByType(nodeType int) (string, string) {
	switch nodeType {
	case node_type.Keyword:
		return "<keyword>", "</keyword>"
	case node_type.Symbol:
		return "<symbol>", "</symbol>"
	case node_type.IntegerConstant:
		return "<integerConstant>", "</integerConstant>"
	case node_type.StringConstant:
		return "<stringConstant>", "</stringConstant>"
	case node_type.Identifier:
		return "<identifier>", "</identifier>"
	case node_type.Class:
		return "<class>", "</class>"
	case node_type.ClassVarDec:
		return "<classVarDec>", "</classVarDec>"
	case node_type.SubroutineDec:
		return "<subroutineDec>", "</subroutineDec>"
	case node_type.ParameterList:
		return "<parameterList>", "</parameterList>"
	case node_type.SubroutineBody:
		return "<subroutineBody>", "</subroutineBody>"
	case node_type.VarDec:
		return "<varDec>", "</varDec>"
	case node_type.Statements:
		return "<statements>", "</statements>"
	case node_type.LetStatement:
		return "<letStatement>", "</letStatement>"
	case node_type.IfStatement:
		return "<ifStatement>", "</ifStatement>"
	case node_type.DoStatement:
		return "<doStatement>", "</doStatement>"
	case node_type.ReturnStatement:
		return "<returnStatement>", "</returnStatement>"
	case node_type.Expression:
		return "<expression>", "</expression>"
	case node_type.Term:
		return "<term>", "</term>"
	case node_type.ExpressionList:
		return "<expressionList>", "</expressionList>"
	default:
		panic("wrong type number!")
	}
}
