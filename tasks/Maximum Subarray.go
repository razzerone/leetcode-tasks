package main

func main() {

}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	bestSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		if currentSum > bestSum {
			bestSum = currentSum
		}
	}

	return bestSum
}
