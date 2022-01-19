package main

import (
	"fmt"
	"math/rand"
)

// optimized for memory utilization
// else we could just index each node value in an array

type Solution struct {
	head       *ListNode
	chainDepth int
}

func Constructor(head *ListNode) Solution {

	return Solution{head, getDepth(head)}
}

func (this *Solution) GetRandom() int {

	var node = this.head
	var selected = rand.Intn(this.chainDepth)

	for i := 0; i < selected; i++ {
		node = node.Next
	}

	return node.Val
}

func getDepth(node *ListNode) int {

	var i int = 1

	for node.Next != nil {
		node = node.Next
		i++
	}

	return i
}

// do not copy //

func main() {

	for i := 1; i < 20; i++ {
		fmt.Println(rand.Intn(5))
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
