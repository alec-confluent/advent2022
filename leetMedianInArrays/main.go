package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergedLen := len(nums1) + len(nums2)

	var l1, l2, flag int

	if mergedLen%2 != 0 {

		l1 = mergedLen / 2
		flag = 1
	} else {

		l1 = mergedLen / 2
		l2 = (mergedLen - 1) / 2
		flag = 2
	}

	fmt.Println("l1", l1)
	fmt.Println("l2", l2)
	fmt.Println("flag", flag)

	// for i := 0; i <= l1; i++ {

	// 	if () {

	// 	}
	// }

	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	nums1 := []int{1}
	nums2 := []int{2, 3, 4, 5, 6}

	concat := append(nums1, nums2...)
	sort.Ints(concat[:])
	fmt.Println(concat)
	fmt.Println("~~~starting~~~")

	longest := findMedianSortedArrays(nums1, nums2)

	fmt.Println("~~~answer~~~")
	fmt.Println(longest)

}

// // first working
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	var answer float64

// 	concat := append(nums1, nums2...)
// 	sort.Ints(concat[:])

// 	// if odd
// 	if len(concat)%2 != 0 {

// 		answer = float64(concat[len(concat)/2])
// 	} else {

// 		a := float64(concat[(len(concat)-1)/2])
// 		b := float64(concat[len(concat)/2])
// 		answer = (a + b) / 2
// 	}

// 	return answer
// }
