package main

func maxSubArray(nums []int) int {
	max, min, sum := 0, 0, 0
	k := 0
	for ; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
		if k == 0 {
			min = nums[k]
			continue
		}
		if min < nums[k] {
			min = nums[k]
		}
	}
	if k == len(nums) {
		return min
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}
		sum = 0
		for j := i; j < len(nums); j++ {
			if sum+nums[j] < 0 {
				if sum > max {
					max = sum
				}
				i = j
				break
			}
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	print(maxSubArray([]int{-1}))
}
