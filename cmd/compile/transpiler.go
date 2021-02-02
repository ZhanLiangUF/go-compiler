package main

import (
	"strings"
)

var ostMap = map[string]string{
	"sum": "+",
	"mul": "*",
	"sub": "-",
	"div": "/",
}

// takes in AST and produce go code
// TO DO fix function parameters to match parserl
func transpile(ast *Node, ast2 *Num) string {
	a := transpileNode(ast, ast2)
	return a
}

func transpileNode(ast *Node, ast2 *Num) string {
	if ast != nil {
		return transpileOp(ast)
	} else {
		return transpileNum(ast2)
	}
}

func transpileOp(ast *Node) string {
	var arr []string
	for _, v := range ast.Exp {
		if ast.Type == "Node" {
			v := v.(*Node)
			arr = append(arr, transpileNode(v, nil))
		} else {
			v := v.(*Num)
			arr = append(arr, transpileNode(nil, v))
		}
	}
	s := strings.Join(arr, ostMap[ast.Val])
	return s
}

func transpileNum(ast2 *Num) string {
	return ast2.Val
}
