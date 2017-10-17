package main

func search(nums []int, target int) int {
	lef, mid, rig := 0, 0, len(nums)-1
	for lef <= rig {
		mid = (lef + rig) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if nums[mid] < nums[lef] && target >= nums[lef] {
				rig = mid - 1
			} else {
				lef = mid + 1
			}
		} else {
			if nums[mid] > nums[rig] && target <= nums[rig] {
				lef = mid + 1
			} else {
				rig = mid - 1
			}
		}
	}
	return -1
}

func main() {
	print(search([]int{5, 1, 3}, 3))
}
