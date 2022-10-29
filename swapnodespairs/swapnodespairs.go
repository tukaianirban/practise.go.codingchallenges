package swapnodespairs

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodePairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	log.Printf("swapping %d and %d", head.Val, head.Next.Val)
	temp := head.Next
	head.Next = temp.Next
	temp.Next = head

	log.Printf("after swapping = %d and %d", temp.Val, temp.Next.Val)

	temp.Next.Next = swapNodePairs(temp.Next.Next)

	return temp
}
