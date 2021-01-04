package main

import (
	"fmt"
	"strings"
)

func lex(s string) []string {
	entries := strings.Split(s, " ")
	lex := make([]string, 0)
	for _, e := range entries {
		eTrim := strings.Trim(e, " ")
		if len(eTrim) > 0 {
			lex = append(lex, eTrim)
		}
	}
	return lex
}

func parse(t []string) {

}

func main() {
	str := "mul 3 sub 2 sum 1 3 4"
	lex := lex(str)
	fmt.Println(lex)
}
