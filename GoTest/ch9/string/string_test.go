package string_test

import (
	"testing"
)

func TestString(t *testing.T) {

	var s string
	t.Log(s)
	s = "Hello"
	t.Log(s)
	t.Log(len(s))
	// s[1] = '3'  // cannot assign to s[1]
	s = "\xE4\xB8\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
		t.Logf("%[1]c %[1]x", c)
	}
}
