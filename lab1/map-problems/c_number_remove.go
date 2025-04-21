package main

import "fmt"

func CountNumbers(array []int) map[int]int {
	count := make(map[int]int)
	for _, x := range array {
		count[x] += 1
	}
	return count
}

func MinRemove(numbers []int) int {
	num_count := CountNumbers(numbers)
	var max_cnt int = 0
	for num, count := range num_count {
		//if we keep num and num + 1
		//(no need to check num - 1, bcs if num-1 is present,
		// this pair will be checked anyway)
		keep_cnt := count + num_count[num + 1] 
		if keep_cnt > max_cnt {
			max_cnt = keep_cnt
		}
	}

	return len(numbers) - max_cnt
}

func RunC() {
	var n int
	arr := make([]int, 0)
	fmt.Scanln(&n)
	for range n {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(MinRemove(arr))
}
