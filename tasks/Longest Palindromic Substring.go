package main

func main() {

}

func longestPalindrome(s string) string {
	s2 := make([]byte, 2*len(s)+1)

	for i := range s {
		s2[2*i+1] = s[i]
	}

	p := make([]int, len(s2))
	r, c := 0, 0

	for i := range s2 {
		mir := 2*c - i

		if i < r {
			p[i] = min(r-i, p[mir])
		}

		for i+1+p[i] < len(s2) && i-1-p[i] >= 0 && s2[i+1+p[i]] == s2[i-1-p[i]] {
			p[i]++
		}

		if i+p[i] > r {
			r, c = i+p[i], i
		}
	}

	center, length := 0, 0

	for i, e := range p {
		if e > length {
			center, length = i, e
		}
	}

	start := (center - length) / 2

	return s[start : start+length]
}

//func longestPalindrome(s string) string {
//	dp := make([][]bool, len(s))
//	l, r := 0, 0
//
//	for i := range dp {
//		dp[i] = make([]bool, len(s))
//		dp[i][i] = true
//	}
//
//	for i := 0; i < len(s)-1; i++ {
//		if s[i] == s[i+1] {
//			dp[i][i+1] = true
//			l, r = i, i+1
//		}
//	}
//
//	for d := 2; d < len(s); d++ {
//		for i := 0; i < len(s)-d; i++ {
//			j := i + d
//			if s[i] == s[j] && dp[i+1][j-1] {
//				dp[i][j] = true
//				l, r = i, j
//			}
//		}
//	}
//
//	return s[l : r+1]
//}
