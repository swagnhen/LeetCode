package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	wt := 0
	mpos := 0
	for i := 0; i < len(height); i++ {
		if height[i] > height[mpos] {
			mpos = i
		}
	}
	for i := 0; i <= height[mpos]; i++ {
		flag := false
		pos := -1
		for j := 0; j < len(height); j++ {
			if height[j] > i {
				if flag {
					wt += j - pos - 1
					pos = j
				} else {
					flag = true
					pos = j
				}
			}
		}
	}
	return wt
}

func main() {
	print(trap([]int{}))
}
