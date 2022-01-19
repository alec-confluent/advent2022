package main

import (
	"fmt"
)

//
func getNumFromNode(node *ListNode) []int {

	var nums = []int{node.Val}

	for node.Next != nil {
		node = node.Next
		nums = append(nums, node.Val)
	}

	return nums
}

func getSum(l1 []int, l2 []int) []int {

	new := []int{}

	iterations := max(len(l1), len(l2))

	remainder := 0
	sum := 0
	for i := 0; i < iterations; i++ {

		if i > len(l1)-1 {

			sum = l2[i] + remainder
		} else if i > len(l2)-1 {

			sum = l1[i] + remainder
		} else {

			sum = l1[i] + l2[i] + remainder
		}

		if sum >= 10 {

			remainder = 1
			sum = sum - 10
		} else {

			remainder = 0
		}

		new = append(new, sum)
	}

	if remainder > 0 {

		new = append(new, remainder)
	}

	return new
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeNode(arr []int) *ListNode {

	head := &ListNode{arr[0], nil}

	var prev = head
	for _, a := range arr[1:] {

		prev.Next = &ListNode{a, nil}
		prev = prev.Next
	}

	return head
}

// do not copy //
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := makeNode1([]int{6, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := makeNode1([]int{5, 6, 9})

	l1arr := getNumFromNode(l1)
	l2arr := getNumFromNode(l2)
	sum := getSum(l1arr, l2arr)

	fmt.Println(l1arr)
	fmt.Println("~~~~~~")
	fmt.Println(l2arr)
	fmt.Println("~~~~~~")
	fmt.Println(sum)
	fmt.Println("~~~~~~")
	// fmt.Println(makeNode(split))
	// fmt.Println("~~~~~~")

}

func makeNode1(arr []int) *ListNode {

	head := &ListNode{arr[0], nil}

	var prev = head
	for _, a := range arr[1:] {

		prev.Next = &ListNode{a, nil}
		prev = prev.Next
	}

	return head
}

// for small numbers this is working
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	new := getNumFromNode(l1) + getNumFromNode(l2)
// 	split := strings.Split(strconv.Itoa(new), "")

//     return makeNode(split)
// }

// func getNumFromNode(node *ListNode) int {

// 	var num = strconv.Itoa(node.Val)

// 	for node.Next != nil {
// 		node = node.Next
// 		num = strconv.Itoa(node.Val) + num
// 	}

// 	var final, _ = strconv.Atoi(num)

// 	return final
// }

// func makeNode(arr []string) *ListNode {

// 	v, _ := strconv.Atoi(arr[len(arr)-1])
// 	head := &ListNode{v, nil}

// 	var prev = head
// 	for i := len(arr) - 2; i >= 0; i-- {
// 		v, _ := strconv.Atoi(arr[i])
// 		prev.Next = &ListNode{v, nil}
// 		prev = prev.Next
// 	}

// 	return head
// }
