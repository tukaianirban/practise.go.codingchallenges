/*
Problem definition: https://leetcode.com/problems/merge-k-sorted-lists/
*/

package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getSortedList(lists []*ListNode) *ListNode {

	var resultList *ListNode = nil
	var headOfResultsList *ListNode = nil
	//  := &headOfResultsList

	flag := true

	minVal := lists[0].Val

	for flag {

		// index of the list found to have the smallest number at its head compared to the others' heads
		idx := 0
		flagMinValUnset := true

		// log.Print("Next run")
		for i := 0; i < len(lists); i++ {

			if lists[i] == nil {
				continue
			}

			if flagMinValUnset {
				minVal = lists[i].Val
				flagMinValUnset = false
				idx = i
				continue
			}

			if lists[i].Val < minVal {
				idx = i
				minVal = lists[i].Val
			}
		}

		// log.Printf("minval found=%d for list number=%d; flagMinValUnset=%t", minVal, idx, flagMinValUnset)

		if flagMinValUnset {
			flag = false
			break
		}

		if resultList == nil {
			resultList = &ListNode{
				Val: lists[idx].Val,
			}
			headOfResultsList = resultList
		} else {
			resultList.Next = &ListNode{
				Val: lists[idx].Val,
			}
			resultList = resultList.Next
		}
		// resultList.Val = lists[idx].Val
		// resultList.Next = &ListNode{}
		// resultList = resultList.Next

		lists[idx] = lists[idx].Next
		// log.Printf("value at head=%d", headOfResultsList.Val)
	}

	resultList = nil

	return headOfResultsList
}
