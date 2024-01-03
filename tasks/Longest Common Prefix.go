package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	length := 0

loop1:
	for i := 0; ; i++ {
		var letter byte

		if i < len(strs[0]) {
			letter = strs[0][i]
		} else {
			break
		}

		for _, s := range strs[1:] {
			if !(i < len(s) && s[i] == letter) {
				break loop1
			}
		}

		length++
	}

	return strs[0][:length]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
