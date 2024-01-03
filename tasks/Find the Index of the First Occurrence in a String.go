package main

import "strings"

func main() {

}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr1(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
