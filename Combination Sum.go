package main

import (
	"fmt"
	"sort"
)

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = nil
	sort.Ints(candidates)
	getSum(candidates, target, []int{}, 0)
	return result
}

func getSum(candidates []int, target int, set []int, pos int) {
	for i := pos; i < len(candidates); i++ {
		if candidates[i] <= target {
			temp := make([]int, len(set))
			copy(temp, set)
			if candidates[i] < target {
				getSum(candidates, target-candidates[i], append(temp, candidates[i]), i)
			} else if candidates[i] == target {
				result = append(result, append(temp, candidates[i]))
			}
		}
	}
	return
}

func main() {
	fmt.Print(combinationSum([]int{2, 3, 7}, 18))
}
