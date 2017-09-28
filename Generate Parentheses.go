package main

var res []string

func generateParenthesis(n int) []string {
	res = nil
	recGP(n*2, 0, 0, 0, "")
	return res
}

func recGP(n int, i int, close int, nl int, temp string) {
	if i == n {
		res = append(res, temp)
		return
	}
	if n > 0 && nl < n/2 {
		recGP(n, i+1, close+1, nl+1, temp+"(")
	}
	if n > 0 && close > 0 {
		recGP(n, i+1, close-1, nl, temp+")")
	}
}

func main() {
	res := generateParenthesis(4)
	for i := 0; i < len(res); i++ {
		print(res[i] + "\n")
	}
}
