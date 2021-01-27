package main

import (
	"fmt"
	"strconv"
)

var ostMap = map[string]string{
	"sum": "+",
	"mul": "*",
	"sub": "-",
	"div": "/",
}

// takes in AST and produce go code
func transpile(ast *Node, ast2 *Num) {

}

func transpileNode(ast *Node, ast2 *Num) (*Node, *Num) {
	if ast != nil {
		return transpileOp(ast), nil
	} else {
		return nil, transpileNum(ast2)
	}
}

func transpileOp(ast *Node) {
	// use map and fmtsprintf to interpolate
	// iterate through num/node
	var arr []interface{}
	for i, v := range ast.Exp {
		arr = append(arr, transpileNode(v))
	}
	fmt.Println(arr)
}

func transpileNum(ast2 *Num) int {
	i, _ := strconv.Atoi(ast2.Val)
	return i
}
