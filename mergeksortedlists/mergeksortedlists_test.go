package mergeksortedlists

import "testing"

func TestGetSortedLists(t *testing.T) {

	// initialize the K sorted lists

	lists := make([]*ListNode, 3)

	// index 0
	lists[0] = &ListNode{
		Val:  1,
		Next: nil,
	}
	lists[0].Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	lists[0].Next.Next = &ListNode{
		Val:  5,
		Next: nil,
	}
	// printList(t, lists[0])

	// index 1
	lists[1] = &ListNode{
		Val:  1,
		Next: nil,
	}
	lists[1].Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	lists[1].Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	// printList(t, lists[1])

	// index 2
	lists[2] = &ListNode{
		Val:  2,
		Next: nil,
	}
	lists[2].Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	// printList(t, lists[2])

	resultList := getSortedList(lists)
	printPtr := resultList
	for printPtr != nil {
		t.Logf("%d", printPtr.Val)
		printPtr = printPtr.Next
	}
}

func printList(t *testing.T, head *ListNode) {

	printPtr := head
	for printPtr != nil {
		t.Logf("%d", printPtr.Val)
		printPtr = printPtr.Next
	}
}
