package main

import (
	"fmt"
	"time"
)

// brute force search
func twoSumBrute(nums []int, target int) []int {

	for i, a := range nums {

		for j, b := range nums[i+1:] {

			if a+b == target {
				return []int{i, i + j + 1}
			}
		}
	}

	return []int{0, 0}
}

// hashmap
func twoSumMap(nums []int, target int) []int {

	evaulated := make(map[int]int)

	for i, a := range nums {

		if val, ok := evaulated[target-a]; ok {

			return []int{val, i}
		}

		evaulated[a] = i
	}

	panic("no solution")
	return []int{0, 0}
}

// do not copy //

func main() {

	run([]int{2, 7, 11, 15, 8}, 9)
	run([]int{1, 3, 11, 0, 5}, 5)
	run([]int{-1, -2, -3, -4, -11, 3}, -8)
}

func run(nums []int, target int) {

	start := time.Now()
	answer := twoSumMap(nums, target)
	duration := time.Since(start)

	fmt.Println(answer)
	fmt.Println(duration)
	fmt.Println("~~~~~~")
}
