package main

import "fmt"

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement(a, 3))
	fmt.Println(a)
}

func removeElement(nums []int, val int) int {
	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

	return p
}
