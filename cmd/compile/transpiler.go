package main

import "strconv"

var ostMap = map[string]string{
	"sum": "+",
	"mul": "*",
	"sub": "-",
	"div": "/",
}

// takes in AST and produce go code
func transpile(ast Node, ast2 Num) {

}

func transpileNode(ast *Node, ast2 *Num) (*Node, int) {
	if ast != nil {
		// return transpileOp(ast), nil
	} else {
		return nil, transpileNum(ast2)
	}
}

// func transpileOp(ast *Node) *Node {

// }

func transpileNum(ast2 *Num) int {
	i, err := strconv.Atoi(ast2.Val)
	if err != nil {
		// error
	}
	return i
}
