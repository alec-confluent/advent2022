package main

import (
	"fmt"
	"sort"
)

// alternate approach
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	mergedLen := len(nums1) + len(nums2)

// 	var l1, l2, flag int

// 	if mergedLen%2 != 0 {

// 		l1 = mergedLen / 2
// 		flag = 1
// 	} else {

// 		l1 = mergedLen / 2
// 		l2 = (mergedLen - 1) / 2
// 		flag = 2
// 	}

// 	fmt.Println("l1", l1)
// 	fmt.Println("l2", l2)
// 	fmt.Println("flag", flag)

// 	var last int
// 	for i := 0; i <= l1; i++ {

// 		if len(nums1) > i {

// 			last = nums1[i]
// 		}

// 		if len(nums2) > i {

// 		}
// 	}

// 	return 0
// }

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

// first working
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var answer float64

	// combine the two arrays together & sort them
	concat := append(nums1, nums2...)
	sort.Ints(concat[:])

	// if the combined array is an odd length,
	// we know we can just take the middle number
	if len(concat)%2 != 0 {

		// to get the middle index location
		// just divide the combined array length by 2
		// ...the int type will auto round up for us
		answer = float64(concat[len(concat)/2])

		// if its an even length, we need to find the middle 2 numbers
	} else {

		// same thing here, we just use the length of the
		// combined array to find the middle 2 numbers
		a := float64(concat[(len(concat)-1)/2])
		b := float64(concat[len(concat)/2])
		answer = (a + b) / 2
	}

	return answer
}
