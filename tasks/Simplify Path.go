package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.FieldsFunc("/home//foo/", func(r rune) bool {
		return r == '/'
	})

	fmt.Printf("%q", res)
}

func simplifyPath(path string) string {
	p := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	if len(p) == 0 {
		return "/"
	}

	resp := make([]string, 0, 1)

	for _, e := range p {
		if e == "." || (e == ".." && len(resp) == 0) {
			continue
		} else if e == ".." {
			resp = resp[:len(resp)-1]
		} else {
			resp = append(resp, e)
		}
	}

	return "/" + strings.Join(resp, "/")
}
