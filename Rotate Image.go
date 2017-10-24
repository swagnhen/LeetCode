package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := i; j < size-i-1; j++ {
			matrix[j][size-i-1], matrix[size-i-1][size-j-1], matrix[size-j-1][i], matrix[i][j] = matrix[i][j], matrix[j][size-i-1], matrix[size-i-1][size-j-1], matrix[size-j-1][i]
		}
	}
}

func main() {
	input := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}
	rotate(input)
	fmt.Print(input)
}
