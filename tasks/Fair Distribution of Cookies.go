package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
}

// TODO

func sliceMin(s []int) int {
	min := 100000
	mini := -1

	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
			mini = i
		}
	}

	return mini
}

func sliceMax(s []int) (max int) {
	for _, e := range s {
		if e > max {
			max = e
		}
	}

	return max
}

func distributeCookies(cookies []int, k int) int {
	if k >= len(cookies) {
		max := 0
		for _, e := range cookies {
			if e > max {
				max = e
			}
		}
		return max
	}

	sort.Slice(cookies, func(i, j int) bool {
		return cookies[i] > cookies[j]
	})

	childrens := make([]int, k)

	for i := 0; i < k; i++ {
		childrens[i] += cookies[i]
	}

	for i := k; i < len(cookies); i++ {
		min := sliceMin(childrens)

		childrens[min] += cookies[i]
	}

	return sliceMax(childrens)
}
