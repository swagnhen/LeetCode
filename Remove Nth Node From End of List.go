package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lmap []*ListNode
	var pre *ListNode
	pre = head
	for pre != nil {
		lmap = append(lmap, pre)
		pre = pre.Next
	}
	if n == len(lmap) {
		if len(lmap) == 1 {
			pre = nil
		} else {
			pre = lmap[1]
		}
		return pre
	}
	pre = lmap[len(lmap)-n-1]
	pre.Next = pre.Next.Next
	return head
}

func main() {

}
