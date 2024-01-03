package main

func main() {

}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prod := 1
	for i, e := range nums {
		res[i] = prod
		prod *= e
	}

	prod = 1

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= prod
		prod *= nums[i]
	}

	return res
}
