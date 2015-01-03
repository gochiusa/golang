package parse

import (
	"testing"
)

func Test_aIsLetter(t *testing.T) {
	actual := isLetter('a')
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_兎IsLetter(t *testing.T) {
	actual := isLetter('兎')
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_0IsNotLetter(t *testing.T) {
	actual := isLetter('0')
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_5IsDigit(t *testing.T) {
	actual := isDigit('5')
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_aIsNotDigit(t *testing.T) {
	actual := isDigit('a')
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_aIsNotWhiteSpace(t *testing.T) {
	actual := isWhiteSpace('a') || isWhiteSpace('1') || isWhiteSpace('兎')
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_IsWhiteSpace(t *testing.T) {
	actual := isWhiteSpace(' ') && isWhiteSpace('\t') && isWhiteSpace('\n')
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadchEOF(t *testing.T) {
	s := NewScanner("ab")
	s.next()
	s.next()
	actual := s.reachEof()
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadchEOFincludeClearLine(t *testing.T) {
	s := NewScanner("a\ncdefg")
	for i := 0; i < 6; i++ {
		s.next()
	}
	actual := s.reachEof()
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	s.next()
	actual = s.reachEof()
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadchEOFincludeClearLineAndMultiByteRunes(t *testing.T) {
	s := NewScanner("ご注文はうさぎですか?\nIs the order a rabbit?")
	for i := 0; i < 33; i++ {
		s.next()
	}
	actual := s.reachEof()
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	s.next()
	actual = s.reachEof()
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNextAndPeek(t *testing.T) {
	s := NewScanner("ご注文はうさぎですか？ Is the order a rabbit? （ごちうさ/ゴチウサ）12345")
	for i := 0; i < 50; i++ {
		s.next()
	}
	actual := s.peek()
	expected := '5'

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestScanLetter(t *testing.T) {
	s := NewScanner("ご注文はうさぎですか PetitRabbits 12470円だよ")
	expected := []string{"ご注文はうさぎですか", "PetitRabbits", "12470円だよ"}
	for i, _ := range expected {
		if actual := s.scanLetter(); actual != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
		s.skipWhiteSpace()
	}
}

func TestScanNumber(t *testing.T) {
	actual := NewScanner("12470").scanNumber()
	expected := "12470"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestScanner(t *testing.T) {
	s := NewScanner("var x = 12470")
	item := s.Scan()

	// fixme: itemKeyword 追加後にチェック
	//	if expected := itemKeyword; item.typ != expected {
	//		t.Errorf("got %v\nwant %v", item.typ, expected)
	//	}
	if expected := "var"; item.val != expected {
		t.Errorf("got %v\nwant %v", item.val, expected)
	}
	if item.pos.count != 3 || item.pos.line != 1 || item.pos.column != 4 {
		t.Errorf("position error: Pos{count: %v, line: %v, column: %v}", item.pos.count, item.pos.line, item.pos.column)
	}
}
