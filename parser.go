package main

import (
	"errors"
	"fmt"
	"regexp"
)

func parse(t []string) {
	var c *int
	c = new(int)
	*c = 0

	a, b := parseExp(t, c)
	fmt.Println(a)
	fmt.Println(b)
}

func parseNum(t []string, c *int) *Num {
	var num *Num
	num = new(Num)
	num.Val = consume(t, c)
	num.Type = "Num"
	fmt.Println(num)
	return num
}

func parseOp(t []string, c *int) *Node {
	var node *Node
	node = new(Node)

	var exp []interface{}

	node.Val = consume(t, c)
	node.Type = "Op"
	node.Exp = exp

	// not nesting nodes here because there will never be a time when error is nil
	_, err := peek(t, c)

	for err == nil {
		a, b := parseExp(t, c)
		if a == nil {
			node.Exp = append(node.Exp, a)
		} else {
			node.Exp = append(node.Exp, b)
		}
	}
	fmt.Println(err)
	return node
}

func parseExp(t []string, c *int) (*Num, *Node) {
	num, _ := peek(t, c)
	fmt.Println(num)
	match, _ := regexp.MatchString(`\d`, num)
	fmt.Println(match)
	if match {
		return parseNum(t, c), nil
	} else {
		return nil, parseOp(t, c)
	}

}

// *int means pointer to an int value
func peek(t []string, c *int) (string, error) {
	if *c >= len(t) {
		return "-1", errors.New("index out of bounds")
	}
	return t[*c], nil
}

func consume(t []string, c *int) string {
	*c++
	return t[*c]
}
