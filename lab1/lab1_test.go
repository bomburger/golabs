package main

import (
	"testing"
	"slices"
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
			t.Errorf("CASE %d: expected %t, got %t\n", i, tc.expected, result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}

func TestB(t *testing.T) {
	cases := []struct {
		line1 string
		line2 string
		expected bool
	}{
		{"IloveGo", "IhateGo", false},
		{"debitcard", "badcredit", true},
		{"abab", "aabbc", false},
	}

	for i, tc := range cases {
		result := IsAnagram(tc.line1, tc.line2)
		if result != tc.expected {
			t.Errorf("CASE %d: expected %t, got %t\n", i, tc.expected, result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}

func TestC(t *testing.T) {
	cases := []struct {
		numbers []int
		expected int
	}{
		{ []int{1, 2, 3, 4, 5}, 3 },
		{ []int{1, 1, 2, 3, 5, 5, 2, 2, 1, 5}, 4 },
		{ []int{3, 3, 3, 4, 4, 4, 4}, 0 },
	}

	for i, tc := range cases {
		result := MinRemove(tc.numbers)
		if result != tc.expected {
			t.Errorf("CASE %d: expected %d, got %d\n", i, tc.expected, result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}

func TestD(t *testing.T) {
	cases := []struct {
		numbers []int
		m int
		expected bool
	}{
		{ []int{8, 1, 2, 1, 3, 4}, 3, true },
		{ []int{1, 2, 3, 4, 5, 6, 7}, 4, false },
		{ []int{1, 1, 1, 1}, 0, false },
	}

	for i, tc := range cases {
		result := NumbersClose(tc.numbers, tc.m)
		if result != tc.expected {
			t.Errorf("CASE %d: expected %t, got %t\n", i, tc.expected, result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}

func TestE(t *testing.T) {
	cases := []struct {
		arrays [][]int
		k int
		expected []int
	}{
		{[][]int {{10, 20, 30}, {60, 20}, {10, 50, 60, 70}, {80}}, 2, []int {30, 50, 70, 80}},
		{[][]int {{1, 2}, {1, 3, 4}, {2}, {1, 2, 5, 6}, {3, 7}}, 3, []int {3, 4, 5, 6, 7}},
	}

	for i, tc := range cases {
		result := RepeatingNumbers(tc.arrays, tc.k)
		if !slices.Equal(result, tc.expected) {
			t.Errorf("CASE %d: expected:\n", i)
			t.Error(tc.expected)
			t.Errorf("got:\n")
			t.Error(result)
		} else {
			t.Logf("CASE %d: PASSED\n", i)
		}
	}
}
