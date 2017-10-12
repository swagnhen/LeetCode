package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	gntPermute(nums, nil, &result)
	return result
}

func gntPermute(nums []int, array []int, result *[][]int) {
	temp := make([]int, len(nums))
	for i := 1; i <= len(nums); i++ {
		copy(temp, nums)
		gntPermute(append(temp[:i-1], temp[i:]...), append(array, nums[i-1]), result)
	}
	if len(nums) == 1 {
		*result = append(*result, append(array, nums[0]))
	}
}

func main() {
	fmt.Print(permute([]int{1, 2, 3, 4}))
}
