package main

import (
	"fmt"
)

func main() {
	str := "sum 2 3 4"
	lex := lex(str)
	fmt.Println(lex)
	a, b := parse(lex)
	// use fmf sprintf to interpolate the string after transpile
	transpile(a, b)
}
