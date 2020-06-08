package string

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Hello", "olleH"},
		{"¶", "¶"},
		{"", ""},
	}

	for _, c := range tests {
		s := Reversible(c.s)
		got := s.Reverse()
		if string(got) != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
