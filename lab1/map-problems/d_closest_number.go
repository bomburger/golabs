package main

import "fmt"

func NumbersClose(arr []int, window int) bool {
	last_occurance := make(map[int]int)
	for i, num := range arr {
		last, ok := last_occurance[num]
		if ok && i - last <= window {
			return true	
		}
		last_occurance[num] = i
	}
	return false
}

func RunD() {
	var n, m int
	numbers := make([]int, 0)
	fmt.Scan(&n)
	fmt.Scan(&m)
	for range n {
		var x int
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	if NumbersClose(numbers, m) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
