package main

func main() {

}

//todo

func canBeValid(s string, locked string) bool {
	if locked[0] == '1' && s[0] == ')' {
		return false
	}

	return generate(s[1:], locked[1:], 1)
}

func generate(s, locked string, k int) bool {
	if k < 0 || (len(s) == 0 && k != 0) {
		return false
	}
	if len(s) == 0 {
		return k == 0
	}

	if locked[0] == '1' {
		if s[0] == '(' {
			k++
		} else {
			k--
		}
		return generate(s[1:], locked[1:], k)
	}

	return generate(s[1:], locked[1:], k+1) || generate(s[1:], locked[1:], k-1)
}
