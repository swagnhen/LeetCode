package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	lef, mid, rig := 0, 0, len(nums)-1
	start, end := -1, -1
	for lef <= rig {
		mid = (lef + rig) / 2
		if target == nums[mid] {
			start, end = mid, mid
			for (start-1) >= 0 && nums[start-1] == nums[start] {
				start--
			}
			for (end+1) < len(nums) && nums[end+1] == nums[end] {
				end++
			}
			break
		} else if target > nums[mid] {
			lef = mid + 1
		} else {
			rig = mid - 1
		}
	}
	return []int{start, end}
}

func main() {
	fmt.Print(searchRange([]int{1}, 1))
}
