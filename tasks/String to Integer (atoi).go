package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(myAtoi("42"))
}

var re = regexp.MustCompile(`\s*([+-]?)(\d*)`)

func myAtoi(s string) int {
	matches := re.FindStringSubmatch(s)

	res := 0

	for _, e := range matches[2] {
		res = res*10 + (int(e) - int('0'))

		if res > 1<<31-1 && matches[1] != "-" {
			return 1<<31 - 1
		}
		if res > 1<<31 && matches[1] == "-" {
			return -(1 << 31)
		}
	}

	if matches[1] == "-" {
		res *= -1
	}

	return res
}
