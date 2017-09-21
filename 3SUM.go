package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	var lo, hi int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i < len(nums) && i != 0 && nums[i] == nums[i-1] {
			continue
		}
		lo = i + 1
		hi = len(nums) - 1
		for lo < hi {
			if lo < len(nums) && lo != i+1 && nums[lo] == nums[lo-1] {
				lo++
				continue
			}
			if hi >= 0 && hi != len(nums)-1 && nums[hi] == nums[hi+1] {
				hi--
				continue
			}
			if nums[i]+nums[lo]+nums[hi] == 0 {
				result = append(result, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
			}
			if nums[i]+nums[lo]+nums[hi] > 0 {
				hi--
			}
			if nums[i]+nums[lo]+nums[hi] < 0 {
				lo++
			}
		}
	}
	return result
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j])
			print(" ")
		}
		print("\n")
	}
}
