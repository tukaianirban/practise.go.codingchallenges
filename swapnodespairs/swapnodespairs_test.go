package swapnodespairs

import "testing"

func TestSwapNodePairs(t *testing.T) {

	node6 := &ListNode{
		Val:  6,
		Next: nil,
	}

	node5 := &ListNode{
		Val:  5,
		Next: node6,
	}

	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}

	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}

	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}

	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	head := swapNodePairs(node1)

	t.Logf("traversing resultant list:=")
	for head != nil {
		t.Logf("%d", head.Val)
		head = head.Next
	}
}
