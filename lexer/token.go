package lexer

import "fmt"

// TokenKind basically an enum containing all token types.
// TokenKind has been changed from int to string for better debugging.
type TokenKind string

// seems like we will have to set the type for every single one because if not go will think they are just strings...
const (
	// Keywords
	VarKeyword      TokenKind = "var (Keyword)"
	SetKeyword      TokenKind = "set (Keyword)"
	ToKeyword       TokenKind = "to (Keyword)"
	IfKeyword       TokenKind = "if (Keyword)"
	ElseKeyword     TokenKind = "else (Keyword)"
	TrueKeyword     TokenKind = "true (Keyword)"
	FalseKeyword    TokenKind = "false (Keyword)"
	FunctionKeyword TokenKind = "function (Keyword)"
	FromKeyword     TokenKind = "from (Keyword)"
	ForKeyword      TokenKind = "for (Keyword)"
	ReturnKeyword   TokenKind = "return (Keyword)"
	WhileKeyword    TokenKind = "while (Keyword)"
	ContinueKeyword TokenKind = "continue (keyword)"
	BreakKeyword    TokenKind = "break (Keyword)"

	// Tokens
	EOF         TokenKind = "EndOfFile"
	IdToken     TokenKind = "Identifier"
	StringToken TokenKind = "String"
	NumberToken TokenKind = "Number"

	// Symbol Tokens
	PlusToken          TokenKind = "Plus(+)"
	MinusToken         TokenKind = "Minus(-)"
	StarToken          TokenKind = "Star(*)"
	SlashToken         TokenKind = "Slash(/)"
	EqualsToken        TokenKind = "Equals(=)"
	NotToken           TokenKind = "Not(!)"
	NotEqualsToken     TokenKind = "Not Equals(!=)"
	CommaToken         TokenKind = "Comma(,)"
	GreaterThanToken   TokenKind = "GreaterThanToken(>)"
	LessThanToken      TokenKind = "LessThanToken(<)"
	GreaterEqualsToken TokenKind = "GreaterEqualsToken(>=)"
	LessEqualsToken    TokenKind = "LessEqualsToken(<=)"
	AmpersandToken     TokenKind = "AmpersandToken(&)"
	AmpersandsToken    TokenKind = "AmpersandsToken(&&)"
	PipeToken          TokenKind = "PipeToken(|)"
	PipesToken         TokenKind = "PipesToken(||)"
	HatToken           TokenKind = "HatToken(^)"

	AssignToken TokenKind = "AssignToken(<-)"

	OpenBraceToken        TokenKind = "OpenBrace({)"
	CloseBraceToken       TokenKind = "Closebrace(})"
	OpenParenthesisToken  TokenKind = "OpenParenthesis"
	CloseParenthesisToken TokenKind = "CloseParenthesis"

	BadToken TokenKind = "Token Error (BadToken)" // Naughty ;)

	Semicolon TokenKind = "Semicolon ;" // Used to separate statements (for now... )
)

// Token stores information about lexical structures in the text
type Token struct {
	Value     string
	RealValue interface{}
	Kind      TokenKind
	Line      int
	Column    int
}

// CreateToken returns a Token created from the arguments provided
func CreateToken(value string, kind TokenKind, line int, column int) Token {
	return Token{
		value, nil, kind, line, column,
	}
}

// CreateTokenReal the majority of the code base uses CreateToken, however, the Token struct has
// a "RealValue" which should store the true value of a Token. For example, NumberToken (TokenKind) created using
// CreateToken will only store its string value and not its real number value. CreateTokenReal will store the
// converted type (so NumberToken actually stores a number).
func CreateTokenReal(buffer string, real interface{}, kind TokenKind, line, column int) Token {
	return Token{
		buffer, real, kind, line, column,
	}
}

// String easy representation of a Token
func (t Token) String(pretty bool) string {
	if !pretty {
		return fmt.Sprintf("Token { value: %s, kind: %s, position: (%d, %d), real: %v}", t.Value, t.Kind, t.Line, t.Column, t.RealValue)
	} else {
		return fmt.Sprintf("Token { \n\tvalue: %s, \n\tkind: %s, \n\tposition: (%d, %d)\n}", t.Value, t.Kind, t.Line, t.Column)
	}
}
