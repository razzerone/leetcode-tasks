package main

func main() {

}

//TODO

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < n+1; i++ {
		dp[i+1] += dp[i]
		if i+2 < n+1 {
			dp[i+2] += dp[i]
		}
	}

	return dp[n]
}
