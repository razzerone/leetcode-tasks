package main

import "strings"

func main() {
	print(intToRoman(20))
}

type pair struct {
	first  int
	second string
}

func intToRoman(num int) string {
	list := []pair{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := &strings.Builder{}
	start := 0
	for num > 0 {
		for i := start; i < len(list); i++ {
			if num-list[i].first < 0 {
				start = i + 1
				continue
			}
			num -= list[i].first
			res.WriteString(list[i].second)
			break
		}
	}
	return res.String()
}
