package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte //current char under examination
}

func New(input string) *Lexer { //Will return a struct of type lexer
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is ASCII code for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
