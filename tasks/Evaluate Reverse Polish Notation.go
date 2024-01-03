package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 2)

	for _, e := range tokens {
		switch e {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		default:
			num, _ := strconv.Atoi(e)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
