package main

import (
	"fmt"
)

func main() {
	str := "sum sum 2"
	lex := lex(str)
	fmt.Println(lex)
	parse(lex)

}
