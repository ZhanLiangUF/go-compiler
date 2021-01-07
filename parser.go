package main

import (
	"fmt"
)

func parse(t []string) {
	var c *int
	c = new(int)
	*c = 0

	peek := peek(t, c)
	consume := consume(t, c)

	parseNum := parseNum(consume)
	parseOp := parseOp(consume)

	fmt.Println(peek)
	fmt.Println(parseNum)
	fmt.Println(parseOp)
}

func parseNum(consume string) *Num {
	var num *Num
	num = new(Num)
	num.Val = consume
	num.Type = "Num"
	return num
}

func parseOp(consume string) *Node {
	var node *Node
	node = new(Node)

	node.Val = consume
	node.Type = "Op"
	node.Exp = []Num{}
	return node
}

// *int means pointer to an int value
func peek(t []string, c *int) string {
	return t[*c]
}

func consume(t []string, c *int) string {
	*c++
	return t[*c]
}
