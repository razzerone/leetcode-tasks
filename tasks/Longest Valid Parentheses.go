package main

func main() {

}

func longestValidParentheses(s string) int {
	dp := make([][]int, len(s))
	for i := range s {
		dp[i] = make([]int, len(s))
	}

}
