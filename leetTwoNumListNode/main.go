package main

import (
	"fmt"
)

/**********************************************************************************************
*
* Excercise
*
***********************************************************************************************/

// Link: https://leetcode.com/problems/add-two-numbers/

// Description:

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

/**********************************************************************************************
*
* Solution Body
*
***********************************************************************************************/

/**
 * @description Returns a list of combined linked nodes in reverse order
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}

	tail := head
	remainder := 0
	for tail != nil {

		a, b := 0, 0

		if l1 != nil {

			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {

			b = l2.Val
			l2 = l2.Next
		}

		tail.Val = a + b + remainder

		if tail.Val >= 10 {

			remainder = 1
			tail.Val = tail.Val - 10
		} else {

			remainder = 0
		}

		if l1 != nil || l2 != nil {

			tail.Next = &ListNode{}
			tail = tail.Next
		} else if remainder > 0 {

			tail.Next = &ListNode{Val: remainder}
			break
		} else {

			break
		}
	}

	return head
}

/**
 * @description Creates a node
 */
func makeNode(arr []int) *ListNode {

	head := &ListNode{arr[0], nil}

	var prev = head
	for _, a := range arr[1:] {

		prev.Next = &ListNode{a, nil}
		prev = prev.Next
	}

	return head
}

/**********************************************************************************************
*
* Non-solution code for running locally
*
***********************************************************************************************/

/**
 * @description A ListNode struct
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @description Main execution block
 */
func main() {

	l1 := makeNode1([]int{6, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := makeNode1([]int{5, 6, 9})

	sum := addTwoNumbers(l1, l2)

	fmt.Printf("%+v\n", sum)
	fmt.Println("~~~~~~")

}

/**
 * @description Node creator to make local testing possible
 */
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
