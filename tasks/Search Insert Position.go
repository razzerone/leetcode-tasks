package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 4}, 6))
}

// binary search

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := (l + r) / 2

		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
