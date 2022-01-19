package main

import (
	"fmt"
)

// had to read this one, sliding window algorithm
func lengthOfLongestSubstring(s string) int {

	answer := 0

	m := make(map[rune]int)

	i := 0
	for j, a := range s {

		if val, ok := m[a]; ok {

			i = max(val, i)
		}

		answer = max(answer, j-i+1)
		m[a] = j + 1
	}

	return answer
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	longest := lengthOfLongestSubstring("pwwkew")

	fmt.Println(longest)
	fmt.Println("~~~~~~")

}

// // working brute force
// func lengthOfLongestSubstring(s string) int {

// 	longest := 0
// 	sequence := make(map[rune]struct{})

// 	for i, a := range s {

// 		sequence[a] = struct{}{}

// 		for _, b := range s[i+1:] {

// 			if _, ok := sequence[b]; ok {

// 				break

// 			} else {

// 				sequence[b] = struct{}{}
// 			}
// 		}

// 		if len(sequence) > longest {

// 			longest = len(sequence)
// 		}

// 		sequence = make(map[rune]struct{})
// 	}

// 	return longest
// }
