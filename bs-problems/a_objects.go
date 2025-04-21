package main

import "fmt"

func MaxK(n int) int {
	low := 0
	high := n

	var ans int

	for low <= high {
		mid := (low + high) / 2
		size := (mid * (mid + 1) * (mid + 5)) / 6 - 1

		if size <= n {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func RunA() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(MaxK(n))
}
