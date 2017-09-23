package main

import "container/list"

func letterCombinations(digits string) []string {
	var res []string
	var dig2let map[byte][]string
	dig2let = make(map[byte][]string)
	dig2let['1'] = []string{}
	dig2let['2'] = []string{"a", "b", "c"}
	dig2let['3'] = []string{"d", "e", "f"}
	dig2let['4'] = []string{"g", "h", "i"}
	dig2let['5'] = []string{"j", "k", "l"}
	dig2let['6'] = []string{"m", "n", "o"}
	dig2let['7'] = []string{"p", "q", "r", "s"}
	dig2let['8'] = []string{"t", "u", "v"}
	dig2let['9'] = []string{"w", "x", "y", "z"}
	resl := list.New()
	for i := 0; i < len(digits); i++ {
		if resl.Len() == 0 {
			for j := 0; j < len(dig2let[digits[i]]); j++ {
				resl.PushBack(dig2let[digits[i]][j])
			}
		} else {
			rlen := resl.Len()
			for j := 0; j < rlen; j++ {
				var temp string
				temp = resl.Front().Value.(string)
				resl.Remove(resl.Front())
				for k := 0; k < len(dig2let[digits[i]]); k++ {
					resl.PushBack(temp + string(dig2let[digits[i]][k]))
				}
			}
		}
	}
	for resl.Len() != 0 {
		res = append(res, resl.Front().Value.(string))
		resl.Remove(resl.Front())
	}
	return res
}

func main() {
	res := letterCombinations("12345678")
	for i := 0; i < len(res); i++ {
		print(res[i] + "\n")
	}
}
