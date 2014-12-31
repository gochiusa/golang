package parse

import (
	"testing"
)

func TestRuneRangeContains_aTozContains_a(t *testing.T) {
	actual := NewRuneRange('a', 'z').Contains('a')
	expexted := true
	if actual != expexted {
		t.Errorf("got %v\nwant %v", actual, expexted)
	}
}

func TestRuneRangeContains_aTozContains_z(t *testing.T) {
	actual := NewRuneRange('a', 'z').Contains('z')
	expexted := true
	if actual != expexted {
		t.Errorf("got %v\nwant %v", actual, expexted)
	}
}

func TestRuneRangeContains日本語_あToおContains_う(t *testing.T) {
	actual := NewRuneRange('あ', 'お').Contains('う')
	expexted := true
	if actual != expexted {
		t.Errorf("got %v\nwant %v", actual, expexted)
	}
}

func TestRuneRangeContains_漢字Contains兎(t *testing.T) {
	actual := NewRuneRange(0x4E00, 0x9FA0).Contains('兎') // 16進数はUnicode漢字範囲
	expexted := true
	if actual != expexted {
		t.Errorf("got %v\nwant %v", actual, expexted)
	}
}
