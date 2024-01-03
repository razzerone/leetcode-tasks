package main

import "slices"

func main() {

}

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	i, j := 0, 0
	res := 0

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			res++
			i++
			j++
		} else {
			j++
		}
	}

	return res
}
