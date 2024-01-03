package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 2, 3}))
}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]

			k++
		}
	}

	return k
}
