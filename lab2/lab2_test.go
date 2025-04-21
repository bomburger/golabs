package main

import (
	"testing"
	"slices"
)

func TestA(t *testing.T) {
	cases := []struct {
		n int
		expected int
	} {
		{0, 0},
		{2, 1},
		{6, 2},
		{77, 6},
	}

	for i, tc := range cases {
		res := MaxK(tc.n)
		if res == tc.expected {
			t.Logf("CASE %d PASSED\n", i)
		} else {
			t.Errorf("CASE %d FAILED\n expected %d, got %d\n", i, tc.expected, res)
		}
	}
}

func TestB(t *testing.T) {
	cases := []struct {
		nums []int
		target int
		expected int
	} {
		{ []int{1, 3, 5}, 3, 3 },
		{ []int{1, 5, 10, 15}, 8, 10 },
		{ []int{1, 2, 4}, 10, 4 },
		{ []int{3, 4, 5}, 2, 3 },
		{ []int {1, 4, 6, 8}, 5, 4 },
	}

	for i, tc := range cases {
		res1, _ := ClosestNumber(tc.nums, tc.target)
		res2, _ := ClosestNumberBin(tc.nums, tc.target)
		if res1 == res2 && res1 == tc.expected {
			t.Logf("CASE %d PASSED\n", i)
		} else {
			t.Errorf("CASE %d FAILED\n expected %d, got %d and %d\n", i, tc.expected, res1, res2)
		}
	}
}

func TestC(t *testing.T) {
	cases := []struct {
		nums []int
		l int
		s int
		expected int
	} {
		{ []int{1, 3, 5, 7, 9}, 2, 4, 0 },
		{ []int{1, 3, 5, 7, 9}, 1, 3, 1 },
		{ []int{1, 3, 5, 7, 9}, 2, 13, -1 },
	}
	for i, tc := range cases {
		res := SubSum(tc.nums, tc.l, tc.s)
		if res == tc.expected {
			t.Logf("CASE %d PASSED\n", i)
		} else {
			t.Errorf("CASE %d. Expected %d, got %d", i, tc.expected, res)
		}
	}
}

func TestD(t *testing.T) {
	cases := []struct {
		as []int
		bs []int
		expected []int
	} {
		{ []int{1, 2, 2, 3, 4, 5}, []int{2, 4, 1, 6}, []int{3, 5, 1, -1} },
	}
	for i, tc := range cases {
		res := RightFindMultiple(tc.as, tc.bs)
		if slices.Equal(res, tc.expected) {
			t.Logf("CASE %d PASSED\n", i)
		} else {
			t.Errorf("CASE %d FAILED\n expected: ", i)
			t.Error(tc.expected)
			t.Errorf("\ngot: ")
			t.Error(res)
		}
	}
}

func TestE(t *testing.T) {
	cases := []struct {
		as []int
		bs []int
		expected []int
	} {
		{ []int{1, 2, 2, 3, 4, 5}, []int{2, 4, 1, 6}, []int{2, 5, 1, -1} },
	}
	for i, tc := range cases {
		res := LeftFindMultiple(tc.as, tc.bs)
		if slices.Equal(res, tc.expected) {
			t.Logf("CASE %d PASSED\n", i)
		} else {
			t.Errorf("CASE %d FAILED\n expected: ", i)
			t.Error(tc.expected)
			t.Errorf("\ngot: ")
			t.Error(res)
		}
	}
}
