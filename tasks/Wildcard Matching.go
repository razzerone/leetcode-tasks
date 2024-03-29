package main

func main() {

}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			dp[0][i+1] = dp[0][i]
		} else {
			break
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j] || dp[i][j+1] || dp[i+1][j]
			}
		}
	}

	return dp[len(s)][len(p)]
}
