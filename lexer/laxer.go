package laxer

type Lexer struct {
	input string //string

	position     int //入力における現在位置
	readPosition int //これから読み込む位置（現在の文字の次）
	ch           byte
}
