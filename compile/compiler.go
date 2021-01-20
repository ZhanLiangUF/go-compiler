package main

import (
	"fmt"
)

func main() {
	str := "sum 2 3 4"
	lex := lex(str)
	fmt.Println(lex)
	parse(lex)
}
