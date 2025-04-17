package main

import (
	"testing"
)

func TestA(t *testing.T) {
	cases := []struct {
		line string
		words []string
		expected bool
	}{
		{"bomba", []string{"ba", "bom"}, true},
		{"IloveGo", []string{"loveG", "I", "paper", "Go", "hate", "love"}, true},
		{"abacababba", []string{"abacab", "aba", "ab", "ca", "cab", "bab", "ba"}, true},
		{"we need to cook", []string{"cook", "we", "oooh", "Jessie?"}, false},
	}

	for i, tc := range cases {
		result := ConsistsOf(tc.line, tc.words)
		if result != tc.expected {
			t.Errorf("CASE %d:expected %t, got %t\n", i, tc.expected, result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}
