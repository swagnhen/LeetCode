package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	i++
	for j := len(nums) - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 4, 3}
	nextPermutation(nums)
	fmt.Print(nums)
}
