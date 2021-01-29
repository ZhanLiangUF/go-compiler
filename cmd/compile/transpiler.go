package main

import (
	"fmt"
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

func transpileOp(ast *Node) string {
	var arr []string
	for i, v := range ast.Exp {
		if ast.Type == "Node" {
			v := v.(*Node)
			a, b := transpileNode(v, nil)
			// arr = append(arr, transpileNode(v, nil))
		} else {
			v := v.(*Num)
			a, b := transpileNode(nil, v)
			// arr = append(arr, transpileNode(nil, v))
		}
	}
	s := strings.Join(arr, ostMap[ast.Val])
	return s
}

func transpileNum(ast2 *Num) string {
	return ast2.Val
}
