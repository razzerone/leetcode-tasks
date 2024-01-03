package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("asasasdfghhy"))
}

func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0
	res := 0
	maxres := 0
	unique := make(map[byte]struct{}, 28)
	for i < len(s) {
		if _, ok := unique[s[i]]; ok {
			delete(unique, s[j])
			j++
			if res > maxres {
				maxres = res
			}
			res--
		} else {
			unique[s[i]] = struct{}{}
			res++
			i++
		}
	}

	if res > maxres {
		return res
	}

	return maxres
}
