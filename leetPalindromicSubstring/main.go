package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-palindromic-substring/

// working, had to read the Expand Around Center approach
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i, _ := range s {

		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)

		length := max(l1, l2)

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandAroundCenter(s string, l int, r int) int {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func main() {

	s := "aaababad"
	// s := "racecar"

	fmt.Println("~~~starting~~~")

	answer := longestPalindrome(s)

	fmt.Println("~~~answer~~~")
	fmt.Println(answer)

}

// working but too slow
// func longestPalindrome(s string) string {

// 	answer := s[0:1]

// 	m := make(map[int]string)

// 	for i, a := range s {

// 		a1 := string(a)
// 		m[i] = a1

// 		for _, b := range s[i+1:] {

// 			a1 = a1 + string(b)
// 			p := string(b) + m[i]

// 			if a1 == p && len(p) > len(answer) {
// 				fmt.Println("Palindromic", a, p)
// 				answer = p
// 			}

// 			m[i] = p

// 			fmt.Println(string(p))
// 		}

// 	}

// 	return answer
// }
