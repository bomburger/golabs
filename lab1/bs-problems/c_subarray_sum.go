package main

import "fmt"


func SubSum(nums []int, size, sum int) int {
	prefix_sums := make([]int, len(nums))
	var s int

	for i, x := range nums {
		s += x
		prefix_sums[i] = s
	}
	for i := 0; i < len(nums) - size; i++ {
		var start, end int
		if i > 0 { start = prefix_sums[i - 1] } else { start = 0 }
		end = prefix_sums[i + size - 1]
		if end - start == sum {
			return i
		}
	}
	return -1
}

func RunC() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}
	requests := make([][2]int, m)
	for i := range m {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		requests[i] = [2]int {a, b}
	}
	for _, r := range requests {
		fmt.Println(SubSum(nums, r[0], r[1]))
	}
}
