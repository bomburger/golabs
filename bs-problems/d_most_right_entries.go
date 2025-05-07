package main

import "fmt"


// finds the last occurance of target in sorted array
// returns index or -1 if target is not in array
func RightFind(nums_sorted []int, target int) int {
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
			for i := mid + 1; i < len(nums_sorted); i++ {
				if nums_sorted[i] != target { return i - 1 }
			}
			return len(nums_sorted) - 1
*/
			if mid < len(nums_sorted) - 1 && nums_sorted[mid + 1] != mid_n {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

// returns 1 based index for each target,
// -1 if target is not present in nums
func RightFindMultiple(nums_sorted []int, targets []int) []int {
	answers := make([]int, len(targets))	
	for i, target := range targets {
		ans := RightFind(nums_sorted, target)
		if ans != -1 { ans++ }
		answers[i] = ans
	}
	return answers
}

func RunD() {
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
	fmt.Println(RightFindMultiple(as, bs))
}
