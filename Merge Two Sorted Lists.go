package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var i int
	pr := &ListNode{0, nil}
	pl1, pl2 := l1, l2
	for i = 0; pl1 != nil && pl2 != nil; i++ {
		if pl1.Val < pl2.Val {
			pr.Next = &ListNode{pl1.Val, nil}
			pl1 = pl1.Next
		} else {
			pr.Next = &ListNode{pl2.Val, nil}
			pl2 = pl2.Next
		}
		pr = pr.Next
		if i == 0 {
			res = pr
		}
	}
	for ; pl1 != nil; i++ {
		pr.Next = &ListNode{pl1.Val, nil}
		pl1 = pl1.Next
		pr = pr.Next
		if i == 0 {
			res = pr
		}
	}
	for ; pl2 != nil; i++ {
		pr.Next = &ListNode{pl2.Val, nil}
		pl2 = pl2.Next
		pr = pr.Next
		if i == 0 {
			res = pr
		}
	}
	return res
}

func main() {
	var l1, l2 *ListNode
	l1, l2 = nil, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		print(res.Val)
		print('\n')
	}
}
