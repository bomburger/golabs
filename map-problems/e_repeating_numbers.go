package main

import (
	"fmt"
	"slices"
)

func CountArraysForNumber(arrays [][]int, number int) int {
	var count int = 0
	for _, array := range arrays {
		if slices.Contains(array, number) {
			count++
		}
	}
	return count
}

func RepeatingNumbers(arrays [][]int, k int) []int {
	repeating := make(map[int]bool) //hashset
	for _, array := range arrays {
		for _, num := range array {
			if repeating[num] { continue }
			count := CountArraysForNumber(arrays, num)
			if count < k {
				repeating[num] = true
			}
		}
	}
	ans := make([]int, 0)
	for num, yes := range repeating {
		if !yes {
			fmt.Println("ATAS")
		}
		ans = append(ans, num)
	}
	slices.Sort(ans)
	return ans
}

func RunE() {
	var n, k, m int
	arrays := make([][]int, 0)
	fmt.Scan(&n)
	fmt.Scan(&k)
	for range n {
		fmt.Scan(&m)
		arr := make([]int, 0)
		var x int
		for range m {
			fmt.Scan(&x)
			arr = append(arr, x)
		}
		arrays = append(arrays, arr)
	}
	fmt.Println(RepeatingNumbers(arrays, k))
}
