package main

import "container/list"

func longestValidParentheses(s string) int {
	max := 0
	sstack := list.New()
	sstack.PushFront(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sstack.PushFront(i)
		} else {
			sstack.Remove(sstack.Front())
			if sstack.Len() == 0 {
				sstack.PushFront(i)
			} else {
				if max < i-sstack.Front().Value.(int) {
					max = i - sstack.Front().Value.(int)
				}
			}
		}
	}
	return max
}

func main() {
	print(longestValidParentheses("()(())"))
}
