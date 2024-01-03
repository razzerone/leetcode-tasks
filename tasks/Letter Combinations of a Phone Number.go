package main

func main() {

}

var symbols = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return rec(digits[1:], symbols[digits[0]-'2'])
}

func rec(digits string, acc []string) []string {
	if len(digits) == 0 {
		return acc
	}
	res := make([]string, 0, len(acc)*len(symbols[digits[0]-'2']))
	for _, str := range acc {
		for _, symbol := range symbols[digits[0]-'2'] {
			res = append(res, str+symbol)
		}
	}

	return rec(digits[1:], res)
}
