package parse

import (
	"fmt"
)

// Pos represents a position in the original input text
type Pos struct {
	count  int
	line   int
	column int
}

func NewStartPos() *Pos {
	return &Pos{count: 0, line: 1, column: 1}
}
func NewPos(count, line, column int) *Pos {
	return &Pos{count: count, line: line, column: column}
}

// Item represents a token or string.
type Item struct {
	typ ItemType
	pos *Pos
	val string
}

func NewItem(typ ItemType, pos *Pos, val string) *Item {
	return &Item{typ: typ, pos: pos, val: val}
}

func (i Item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ < itemKeyword:
		return fmt.Sprintf("<%s>", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

// ItemType identifies the type of lex items.
type ItemType int

const eof = -1
const (
	itemUnknown ItemType = iota
	itemError
	itemEOF
	itemOperator // operator
	itemBool     // boolean constant
	itemNumber   // simple number
	itemString   // quoted string includes quotes
	itemVariable //	variable
	itemKeyword  // used only to delimit the keywords
	itemIf       // if keyword
	itemElse     // else keyword
)

var key = map[string]ItemType{
	"if":   itemIf,
	"else": itemElse,
}

type Scanner struct {
	src []rune
	pos *Pos
}

func NewScanner(src string) *Scanner {
	return &Scanner{src: []rune(src), pos: NewStartPos()}
}

func (s *Scanner) Scan() (item *Item) {
	s.skipWhiteSpace()
	switch ch := s.peek(); {
	case isOperator(ch):
		item = NewItem(itemOperator, s.pos, s.scanOperator())
	case isLetter(ch):
		item = NewItem(itemString, s.pos, s.scanLetter())
	case isDigit(ch):
		item = NewItem(itemNumber, s.pos, s.scanNumber())
	case ch == eof:
		item = NewItem(itemEOF, s.pos, "<<EOF>>")
	default:
		item = NewItem(itemUnknown, s.pos, "<<UNKNOW>>")
	}
	return
}

func (s *Scanner) peek() rune {
	if !s.reachEof() {
		return s.src[s.pos.count]
	} else {
		return eof
	}
}

func (s *Scanner) next() {
	if !s.reachEof() {
		if s.peek() == '\n' {
			s.pos.line++
			s.pos.column = 1
		} else {
			s.pos.column++
		}
		s.pos.count++
	}
}

func (s *Scanner) reachEof() bool {
	return len(s.src) <= s.pos.count
}

func (s *Scanner) skipWhiteSpace() {
	for isWhiteSpace(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanNumber() string {
	var num []rune
	for isDigit(s.peek()) {
		num = append(num, s.peek())
		s.next()
	}
	return string(num)
}

func (s *Scanner) scanLetter() string {
	var letter []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		letter = append(letter, s.peek())
		s.next()
	}
	return string(letter)
}

func (s *Scanner) scanOperator() string {
	var op []rune
	for isOperator(s.peek()) {
		op = append(op, s.peek())
		s.next()
	}
	return string(op)
}

func (s *Scanner) getPosition() *Pos {
	return &Pos{count: s.pos.count, line: s.pos.line, column: s.pos.column}
}

func isOperator(ch rune) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

func isLetter(ch rune) bool {
	const (
		SMALL_START    = 'a'
		SMALL_END      = 'z'
		CAPITAL_START  = 'A'
		CAPITAL_END    = 'Z'
		HIRAGANA_START = 0x3041
		HIRAGANA_END   = 0x3093
		KATAKANA_START = 0x30A1
		KATAKANA_END   = 0x30F6
		KANJI_START    = 0x4E00
		KANJI_END      = 0x9FA0
	)
	small := NewRuneRange(SMALL_START, SMALL_END)
	capital := NewRuneRange(CAPITAL_START, CAPITAL_END)
	hiragana := NewRuneRange(HIRAGANA_START, HIRAGANA_END)
	katakana := NewRuneRange(KATAKANA_START, KATAKANA_END)
	kanji := NewRuneRange(KANJI_START, KANJI_END)
	return small.Contains(ch) || capital.Contains(ch) || ch == '_' || hiragana.Contains(ch) || katakana.Contains(ch) || kanji.Contains(ch)
}

func isDigit(ch rune) bool {
	return NewRuneRange('0', '9').Contains(ch)
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
