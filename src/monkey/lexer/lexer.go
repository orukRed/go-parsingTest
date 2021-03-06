package lexer

import "go_parsing/src/monkey/token"

type Lexer struct {
	input        string
	position     int //入力における現在位置
	readPosition int //これから読み込む位置（現在の文字の次）
	ch           byte
}

// 引数の値を使って構造体をポインタで初期化
func New(input string) *Lexer {
	l := &Lexer{input: input} //new(Lexer)と同じ・
	l.readChar()
	return l
}

// 次の一文字を読んでinput文字列の現在位置を進める
// ASCII文字のみをカバー。
// Unicodeは非対応
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}

}
