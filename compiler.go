package main

import (
	"fmt"
)

func main() {
	str := "sum sum 2"
	lex := lex(str)
	parse(lex)
	fmt.Println(lex)
}
