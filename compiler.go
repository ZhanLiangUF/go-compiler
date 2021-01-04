package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "mul 3 sub 2 sum 1 3 4"
	var entries = strings.Split(str, " ")
	lex := make([]string, 0)
	for _, e := range entries {
		eTrim := strings.Trim(e, " ")
		if len(eTrim) > 0 {
			lex = append(lex, eTrim)
		}
	}
	fmt.Println(lex)
}
