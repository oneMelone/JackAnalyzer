package tokenizer

import (
	"onemelone/JackAnalyzer/node_type"
	"regexp"
	"strconv"
	"strings"
)

type Tokenizer struct {
	bytes            []byte
	currentPosition  int
	currentToken     string
	currentTokenType int
}

var (
	rules []*regexp.Regexp
	types []int
)

func init() {
	keywords := []string{
		"class", "constructor", "function", "method",
		"field", "static", "var",
		"int", "char", "boolean", "void",
		"true", "false", "null", "this",
		"let", "do", "if", "else", "while", "return",
	}
	symbols := []string{
		"{", "}", "(", ")", "[", "]", ".", ";", "+", "-", "*", "/", "&", "|",
		"<", ">", "=", "~",
	}

	for i := 0; i < len(symbols); i++ {
		symbols[i] = regexp.QuoteMeta(symbols[i])
	}

	keywordRulePattern := "(" + strings.Join(keywords, "|") + ")"
	rules = append(rules, regexp.MustCompile(keywordRulePattern))
	types = append(types, node_type.Keyword)

	symbolsRulePattern := "(" + strings.Join(symbols, "|") + ")"
	rules = append(rules, regexp.MustCompile(symbolsRulePattern))
	types = append(types, node_type.Symbol)

	intConstRulePattern := "[0-9]+"
	rules = append(rules, regexp.MustCompile(intConstRulePattern))
	types = append(types, node_type.IntegerConstant)

	stringConstRulePattern := "\".*?\""
	rules = append(rules, regexp.MustCompile(stringConstRulePattern))
	types = append(types, node_type.StringConstant)

	identifierRulePattern := "[a-zA-Z0-9_]+"
	rules = append(rules, regexp.MustCompile(identifierRulePattern))
	types = append(types, node_type.Identifier)

	blankPattern := "\\s+"
	rules = append(rules, regexp.MustCompile(blankPattern))
	types = append(types, -1)
}

func NewTokenizer(bytes []byte) *Tokenizer {
	return &Tokenizer{
		bytes:            bytes,
		currentPosition:  0,
		currentToken:     "",
		currentTokenType: -1,
	}
}

// HasMoreTokens Return true if there is more tokens in the input.
func (t *Tokenizer) HasMoreTokens() bool {
	return t.currentPosition < len(t.bytes)
}

// Advance Take a step, make the current token to be the next one.
//	Current token is nil at first.
//	If advanced, return true; if no more token to advance, return false.
func (t *Tokenizer) Advance() bool {
	if !t.HasMoreTokens() {
		return false
	}

	for index, rule := range rules {
		loc := rule.FindIndex(t.bytes[t.currentPosition:])
		if len(loc) == 0 {
			continue
		}
		if loc[0] == 0 {
			t.currentTokenType = types[index]
			// if here is a blank token, advance again
			if t.currentTokenType == -1 {
				t.currentPosition += loc[1]
				return t.Advance()
			}
			absoluteStart := t.currentPosition
			absoluteEnd := t.currentPosition + loc[1]
			t.currentToken = string(t.bytes[absoluteStart:absoluteEnd])
			t.currentPosition = absoluteEnd
			return true
		}
	}

	panic("No rule is fit!")
}

// TokenType Return the type of the current token
func (t *Tokenizer) TokenType() int {
	return t.currentTokenType
}

// Keyword Return the keyword of the current token.
func (t *Tokenizer) Keyword() string {
	if t.currentTokenType != node_type.Keyword {
		panic("Keyword should be called only for keyword token")
	}
	return t.currentToken
}

// Symbol Return the symbol of the current token.
func (t *Tokenizer) Symbol() string {
	if t.currentTokenType != node_type.Symbol {
		panic("Symbol should be called only for symbol token")
	}
	return t.currentToken
}

// Identifier Return the identifier of the current token.
func (t *Tokenizer) Identifier() string {
	if t.currentTokenType != node_type.Identifier {
		panic("Identifier should be called only for identifier token")
	}
	return t.currentToken
}

// IntVal Return the integer value of the current token.
func (t *Tokenizer) IntVal() int {
	if t.currentTokenType != node_type.IntegerConstant {
		panic("IntVal should be called only for integer value token")
	}
	intVal, err := strconv.Atoi(t.currentToken)
	if err != nil {
		panic(err)
	}
	return intVal
}

// StringVal Return the string value of the current token.
func (t *Tokenizer) StringVal() string {
	if t.currentTokenType != node_type.StringConstant {
		panic("StringVal should be called only for string value token")
	}
	return t.currentToken[1 : len(t.currentToken)-1]
}

func (t *Tokenizer) Token() string {
	return t.currentToken
}
