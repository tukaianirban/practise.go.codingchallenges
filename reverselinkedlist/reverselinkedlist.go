/*
Problem definition: reverse a linked list
Acceptable: the original linked list can be modified
*/
package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {

	oldHead := head
	newHead := head

	for head.Next != nil {

		// points to the node being moved
		newHead = head.Next

		// next of head is the uberNext node
		head.Next = newHead.Next

		// node to be moved points to head
		newHead.Next = oldHead

		oldHead = newHead
	}

	return newHead
}
