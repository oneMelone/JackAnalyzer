package tokenizer

type Tokenizer struct {
	bytes            []byte
	currentPosition  int
	currentToken     string
	currentTokenType int
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
//	If call this when HasMoreTokens is false, panic.
func (t *Tokenizer) Advance() {

}

// TokenType Return the type of the current token
func (t *Tokenizer) TokenType() int {
	return 0
}

// Keyword Return the keyword of the current token.
func (t *Tokenizer) Keyword() string {
	return ""
}

// Symbol Return the symbol of the current token.
func (t *Tokenizer) Symbol() string {
	return ""
}

// Identifier Return the identifier of the current token.
func (t *Tokenizer) Identifier() string {
	return ""
}

// IntVal Return the integer value of the current token.
func (t *Tokenizer) IntVal() int {
	return 0
}

// StringVal Return the string value of the current token.
func (t *Tokenizer) StringVal() string {
	return ""
}
