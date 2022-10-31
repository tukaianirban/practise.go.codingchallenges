package reverselinkedlist

import "testing"

func TestReverseLinkedList(t *testing.T) {

	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}

	head.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}

	head.Next.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	t.Log("original linked list :::")
	printLinkedList(t, head)

	newHead := reverseLinkedList(head)
	t.Log("reversed linked list :::")
	printLinkedList(t, newHead)
}

func printLinkedList(t *testing.T, head *ListNode) {

	movingPtr := head

	for movingPtr != nil {
		t.Logf("%d", movingPtr.Val)
		movingPtr = movingPtr.Next
	}
}
