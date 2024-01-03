package main

import "fmt"

func main() {
	fmt.Printf("%q", generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	return gen("(", 1, n, 2*n-1)
}

func gen(acc string, k, n, m int) []string {
	if k < 0 {
		return nil
	}
	if m == 0 {
		if n == 0 && k == 0 {
			return []string{acc}
		}
		return nil
	}
	open := gen(acc+"(", k+1, n, m-1)
	closed := gen(acc+")", k-1, n-1, m-1)
	return append(open, closed...)
}
