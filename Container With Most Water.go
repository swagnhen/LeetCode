package main

func maxArea(height []int) int {
	max := 0
	low := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] < height[j] {
				low = height[i]
			} else {
				low = height[j]
			}
			if ((j - i) * low) > max {
				max = (j - i) * low
			}
		}
	}
	return max
}

func main() {
	print(maxArea([]int{1, 3, 7, 13}))
}
