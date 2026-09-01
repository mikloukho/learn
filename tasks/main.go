package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var curr int
	var result ListNode

	for true {
		if list1.Val < list2.Val {
			curr = list1.Val
		} else {
			curr = list2.Val
		}
	}

	return &result
}
