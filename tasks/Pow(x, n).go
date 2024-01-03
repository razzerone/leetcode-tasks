package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10))
}

func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 != 0 {
		return x * myPow(x*x, n/2)
	}
	return myPow(x*x, n/2)
}
