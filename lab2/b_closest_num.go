package main

import (
	"fmt"
	"slices"
)

func abs(n int) int {if n < 0 {return -n}; return n}
func min(a, b int) int {if a < b {return a}; return b}

func ClosestNumber(array []int, target int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("slice should not be empty\n")
	}
	var min_diff int = abs(array[0] - target)
	var closest int = array[0]
	for _, num := range array {
		diff := abs(num - target)
		if diff < min_diff {
			min_diff = diff
			closest = num
		}
	}
	return closest, nil
}

func ClosestNumberBin(array []int, target int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("slice should not be empty\n")
	}
	nums := make([]int, len(array))
	copy(nums, array)
	slices.Sort(nums)
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return target, nil
		}
	}
	d1 := abs(nums[low] - target)
	d2 := abs(nums[high] - target)
	if d1 == d2 {
		return min(nums[low], nums[high]), nil
	} else if d1 < d2{
		return nums[low], nil
	} else {
		return nums[high], nil
	}
}

func RunB() {
	var n, x int
	numbers := make([]int, 0)
	fmt.Scanln(&n)
	for range n {
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	fmt.Scanln(&x)
	res1, _ := ClosestNumber(numbers, x)
	res2, _ := ClosestNumberBin(numbers, x)
	fmt.Printf("Linear solution: %d\n", res1)
	fmt.Printf("Binary search solution: %d\n", res2)
}
