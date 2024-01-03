package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1563847412))
}

func reverse(x int) int {
	digits := make([]int, 0, 10)
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}

	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	res := 0
	for _, e := range digits {
		res = res*10 + e
		if res > 1<<31-1 {
			return 0
		}
	}

	if isNeg {
		if res > 1<<31 {
			return 0
		}
		return -res
	}

	return res
}
