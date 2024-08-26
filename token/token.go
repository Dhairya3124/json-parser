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
	COLON                = ":"
	COMMA                = ","
	STRING               = "STRING"
	NUMBER               = "NUMBER"
	ILLEGAL              = "ILLEGAL"
	EOF                  = "EOF"
)