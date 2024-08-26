package lexer

import "json-parser/token"
type Lexer struct{
	input string
	position int // Current position
	readPosition int // Next Position
	ch byte // Current character to be read
}
// The approach is according to the ASCII Code and will be upgraded
// for UTF-8.
func New(input string) *Lexer{
	l:=&Lexer{input: input}
	l.readChar()
	return l
}
func (l *Lexer)readChar(){
	if l.readPosition >= len(l.input){
		l.ch = 0
	}else{
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

}
func (l* Lexer) NextToken() token.Token{
	var tok token.Token
	switch l.ch{
	case ':':
		tok = newToken(token.COLON,l.ch)
	case ';':
		tok = newToken(token.COMMA,l.ch)
	case '(':
		tok = newToken(token.LEFT_PARENTHESIS,l.ch)
	case ')':
		tok = newToken(token.RIGHT_PARENTHESIS,l.ch)
	case '{':
		tok = newToken(token.BRACE_OPEN,l.ch)
	case '}':
		tok = newToken(token.BRACE_CLOSE,l.ch)
	case '[':
		tok = newToken(token.SQUARE_BRACKET_OPEN,l.ch)
	case ']':
		tok = newToken(token.SQUARE_BRACKET_CLOSE,l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		
	}
	l.readChar()
	return tok

}
func newToken(tokenType token.TokenType,ch byte)token.Token{
	return token.Token{Type: tokenType,Literal: string(ch)}
}