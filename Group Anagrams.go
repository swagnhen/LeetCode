package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	cmap := [][26]int{}
	i := 0
	for ; i < len(strs); i++ {
		j := 0
		temp := [26]int{}
		for k := 0; k < len(strs[i]); k++ {
			temp[strs[i][k]-'a']++
		}
		for ; j < len(result); j++ {
			if temp == cmap[j] {
				result[j] = append(result[j], strs[i])
				break
			}
		}
		if j == len(result) {
			cmap = append(cmap, temp)
			result = append(result, []string{strs[i]})
		}
	}
	return result
}

func main() {
	fmt.Print(groupAnagrams([]string{"hos", "boo", "nay", "deb", "wow", "bop", "bob", "brr", "hey", "rye", "eve", "elf", "pup", "bum", "iva", "lyx", "yap", "ugh", "hem", "rod", "aha", "nam", "gap", "yea", "doc", "pen", "job", "dis", "max", "oho", "jed", "lye", "ram", "pup", "qua", "ugh", "mir", "nap", "deb", "hog", "let", "gym", "bye", "lon", "aft", "eel", "sol", "jab"}))
}
