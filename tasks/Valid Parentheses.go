package main

import "fmt"

func main() {
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
	st := make([]rune, 0, len(s))

	for _, e := range s {
		switch e {
		case ')':
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		case '}':
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		default:
			st = append(st, e)
		}
	}

	return len(st) == 0
}
