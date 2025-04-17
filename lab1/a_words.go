package main

import "fmt"

func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	for i := range prefix {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func remove[T comparable](list []T, value T) []T {
	for i, v := range list {
		if v == value {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func ConsistsOf(s string, words []string) bool {
	for _, word := range words {
		if word == s {
			return true
		}
		if startsWith(s, word) {
			var new_words []string = remove(words, word)
			if ConsistsOf(s[len(word):], new_words) {
				return true
			}
		}
	}
	return false
}

func main() {
	var line string
	var n int
	var words []string
	var word string
	fmt.Scanln(&line)
	fmt.Scanln(&n)
	for range n {
		fmt.Scan(&word)
		words = append(words, word)
	}
	
	res := ConsistsOf(line, words)
	if res {
		for _, w := range words {
			fmt.Printf("%s ", w)
		}
		fmt.Println()
	} else {
		fmt.Println(line)
	}
}
