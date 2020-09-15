package palitrans

import (
	"testing"
)

func TestCharAt(t *testing.T) {
	s := CharAt("abc", -1)
	if s != "" {
		t.Error(s)
		return
	}

	s = CharAt("I am 字串", 5)
	if s != "字" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 0)
	if s != "ā" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 1)
	if s != "d" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 2)
	if s != "ā" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 3)
	if s != "" {
		t.Error(s)
		return
	}
}
