package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	BRACE_OPEN           = "{"
	BRACE_CLOSE          = "}"
	SQUARE_BRACKET_OPEN  = "["
	SQUARE_BRACKET_CLOSE = "]"
	LEFT_PARENTHESIS     = "("
	RIGHT_PARENTHESIS    = ")"
	COLON   = ":"
	COMMA   = ","
	SEMICOLON = ";"
	STRING  = "STRING"
	NUMBER  = "NUMBER"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)