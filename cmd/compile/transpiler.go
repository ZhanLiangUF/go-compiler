package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ostMap = map[string]string{
	"sum": "+",
	"mul": "*",
	"sub": "-",
	"div": "/",
}

// takes in AST and produce go code
func transpile(ast *Node, ast2 *Num) {
	a, b := transpileNode(ast, ast2)
	fmt.Println(a)
	fmt.Println(b)
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
	var arr []string
	for i, v := range ast.Exp {
		if ast.Type == "Node" {
			v := v.(*Node)
			a, b := transpileNode(v, nil)
			// append the result
			// arr = append(arr, transpileNode(v, nil))
		} else {
			v := v.(*Num)
			a, b := transpileNode(nil, v)
			// arr = append(arr, transpileNode(nil, v))
		}
	}
	s := strings.Join(arr, ostMap[ast.Val])
	fmt.Println(arr)
}

func transpileNum(ast2 *Num) int {
	i, _ := strconv.Atoi(ast2.Val)
	return i
}
