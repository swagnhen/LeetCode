package main

import "container/list"

func isValid(s string) bool {
	sstack := list.New()
	var temp uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			sstack.PushFront(s[i])
			continue
		}
		if sstack.Len() == 0 {
			return false
		}
		temp = sstack.Front().Value.(uint8)
		sstack.Remove(sstack.Front())
		if s[i] == ')' && temp != '(' {
			return false
		}
		if s[i] == ']' && temp != '[' {
			return false
		}
		if s[i] == '}' && temp != '{' {
			return false
		}
	}
	if sstack.Len() != 0 {
		return false
	}
	return true
}

func main() {
	print(isValid("([])"))
}
