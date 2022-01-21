package main

import (
	"math/rand"
)

/**********************************************************************************************
*
* Excercise
*
***********************************************************************************************/

// Link: https://leetcode.com/problems/linked-list-random-node/

// Description:

// Given a singly linked list, return a random node's value from the linked list.
// Each node must have the same probability of being chosen.

// Implement the Solution class:

//     Solution(ListNode head) Initializes the object with the head of the singly-linked list head.
//     int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of
//     the list should be equally likely to be chosen.

/**********************************************************************************************
*
* Solution Body
*
***********************************************************************************************/

/**
 * @description Returns a random list nodes value
 */
func (this *Solution) GetRandom() int {

	var node = this.head
	var selected = rand.Intn(this.chainDepth)

	for i := 0; i < selected; i++ {
		node = node.Next
	}

	return node.Val
}

/**
 * @description Determines how many links there are in the chain
 */
func getDepth(node *ListNode) int {

	var i int = 1

	for node.Next != nil {
		node = node.Next
		i++
	}

	return i
}

/**
 * @description Solution definition
 */
type Solution struct {
	head       *ListNode
	chainDepth int
}

/**
 * @description Solution constructor
 */
func Constructor(head *ListNode) Solution {

	return Solution{head, getDepth(head)}
}

/**********************************************************************************************
*
* Non-solution code for running locally
*
***********************************************************************************************/

func main() {

	// DO NOTHING
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
