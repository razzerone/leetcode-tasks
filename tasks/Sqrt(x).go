package main

import "fmt"

func main() {
	fmt.Println(mySqrt(3))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	l, r := 0, x

	for l < r {
		mid := (l + r) / 2

		if mid*mid > x {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if l*l != x {
		return l - 1
	}

	return l
}
