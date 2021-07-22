package main

import (
	"fmt"
)

type ListNode struct {
	value int
	Next  *ListNode
}

func sortSpecailList(head *ListNode) *ListNode {
	oddList, evenList := splitList(head)
	evenList = reverseList(evenList)
	printList(oddList)
	printList(evenList)
	head = mergeList(oddList, evenList)
	return head
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	flag := 1
	var oddList *ListNode
	var evenList *ListNode
	now := head
	odd := head
	even := head.Next
	oddList = odd
	evenList = even
	for now.Next != nil {
		if flag%2 == 1 {
			next := now.Next
			odd.Next = next.Next
			now = next
			odd = odd.Next
			flag++
		} else {
			next := now.Next
			even.Next = next.Next
			now = next
			even = even.Next
			flag++
		}
	}
	return oddList, evenList
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	now := head
	for now != nil {
		next := now.Next
		now.Next = pre
		pre = now
		now = next
	}
	return pre
}

func mergeList(headL *ListNode, headR *ListNode) *ListNode {
	var head *ListNode
	if headL.value < headR.value {
		head = headL
		headL = headL.Next
	} else {
		head = headR
		headR = headR.Next
	}
	now := head
	for headL != nil && headR != nil {
		if headL.value <= headR.value {
			now.Next = headL
			headL = headL.Next
			now = now.Next
		} else {
			now.Next = headR
			headR = headR.Next
			now = now.Next
		}
	}
	if headL != nil {
		now.Next = headL
	}else{
		now.Next = headR
	}
	return head
}

func printList(head *ListNode) {
	for now := head; now != nil; now = now.Next {
		fmt.Print(now.value)
	}
	fmt.Println("")
}

func main() {

	n7 := &ListNode{7, nil}
	n6 := &ListNode{2, n7}
	n5 := &ListNode{5, n6}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{6, n3}
	n1 := &ListNode{1, n2}
	fmt.Println("---before sort---")
	printList(n1)
	n1 = sortSpecailList(n1)
	fmt.Println("---after sort---")
	printList(n1)
}
