package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input      string
		want       int
		countLines bool
		countBytes bool
	}{
		{"hello world", 2, false, false},
		{"hello world again", 3, false, false},
		{"I\n love\n you so much", 3, true, false},
		{"This is just one line", 1, true, false},
		{"hello", 5, false, true},
	}

	for _, tc := range tests {
		b := bytes.NewBufferString(tc.input)
		res := count(b, tc.countLines, tc.countBytes)
		if res != tc.want {
			t.Errorf("expected %d, got %d", tc.want, res)
		}
	}
}
