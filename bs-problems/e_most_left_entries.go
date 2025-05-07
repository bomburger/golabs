package main

import "fmt"

// finds the first occurance of target in sorted array
// returns index or -1 if target is not in array
func LeftFind(nums_sorted []int, target int) int {
	var left, right int
	left = 0
	right = len(nums_sorted) - 1
	for left < right {
		mid := (left + right) / 2
		mid_n := nums_sorted[mid]
		if mid_n > target {
			right = mid - 1
		} else if mid_n < target {
			left = mid + 1	
		} else {
/*
			for i := mid - 1; i >= 0; i-- {
				if nums_sorted[i] != target { return i + 1 }
			}
			return 0
*/
			if mid > 0 && nums_sorted[mid - 1] != mid_n {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

// returns 1 based index for each target,
// -1 if target is not present in nums
func LeftFindMultiple(nums_sorted []int, targets []int) []int {
	answers := make([]int, len(targets))	
	for i, target := range targets {
		ans := LeftFind(nums_sorted, target)
		if ans != -1 { ans++ }
		answers[i] = ans
	}
	return answers
}

func RunE() {
	var n, m int
	var as, bs []int
	fmt.Scan(&n)
	for range n {
		var x int
		fmt.Scan(&x)
		as = append(as, x)
	}
	fmt.Scan(&m)
	for range m {
		var x int
		fmt.Scan(&x)
		bs = append(bs, x)
	}
	fmt.Println(LeftFindMultiple(as, bs))
}
