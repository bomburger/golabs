package main

import "fmt"

func CountLetters(word string) map[rune]int {
	letters := make(map[rune]int)
	for _, r := range word {
		letters[r] += 1
	}
	return letters
}

func EqualMaps[K comparable, V comparable](m1, m2 map[K]V) bool {
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v2 != v1 {
			return false
		}
	}
	for k2, v2 := range m2 {
		v1, ok := m1[k2]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func IsAnagram(s1, s2 string) bool {
	count1 := CountLetters(s1)
	count2 := CountLetters(s2)
	return EqualMaps(count1, count2)
}

func RunB() {
	var line1 string
	var line2 string
	fmt.Scanln(&line1)
	fmt.Scanln(&line2)
	if IsAnagram(line1, line2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	RunB()
}
