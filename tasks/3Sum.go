package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// TODO

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return dfs(nil, nums, 3)
}

func dfs(acc, nums []int, n int) [][]int {
	if n == 0 {
		if acc[0]+acc[1]+acc[2] == 0 {
			return [][]int{acc}
		}
		return nil
	}

	var res [][]int

	for i := 0; i < len(nums)-n+1; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, dfs(append(acc, nums[i]), nums[i+1:], n-1)...)
	}
	return res
}
