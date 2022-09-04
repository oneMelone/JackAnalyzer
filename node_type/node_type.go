package node_type

// Node type
const (
	Keyword = iota
	Symbol
	IntegerConstant
	StringConstant
	Identifier

	Class
	ClassVarDec
	SubroutineDec
	ParameterList
	SubroutineBody
	VarDec

	Statements
	LetStatement
	IfStatement
	WhileStatement
	DoStatement
	ReturnStatement

	Expression
	Term
	ExpressionList
)
