package main

func main() {

}

func plusOne(digits []int) []int {
	jump := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+jump > 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + jump
			return digits
		}
	}

	res := append(make([]int, 1, len(digits)+1), digits...)
	res[0] = 1
	return res
}
