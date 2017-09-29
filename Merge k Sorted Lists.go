package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result, pre *ListNode
	result = &ListNode{0, nil}
	pre = result
	for {
		end := 0
		max := 0
		for ; end < len(lists); end++ {
			if lists[end] != nil {
				max = end
				break
			}
		}
		if end == len(lists) {
			break
		}
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && lists[i].Val < lists[max].Val {
				max = i
			}
		}
		pre.Next = lists[max]
		pre = pre.Next
		lists[max] = lists[max].Next
	}
	return result.Next
}

func main() {
	var lists []*ListNode
	lists = append(lists, &ListNode{1, nil})
	lists = append(lists, &ListNode{0, nil})
	res := mergeKLists(lists)
	print(res.Val)
}
