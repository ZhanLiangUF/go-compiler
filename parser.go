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

	return parseExp(t, c)
}

func parseNum(t []string, c *int) *Num {
	var num *Num
	num = new(Num)
	num.Val = consume(t, c)
	num.Type = "Num"
	return num
}

func parseOp(t []string, c *int) *Node {
	var node *Node
	node = new(Node)

	node.Val = consume(t, c)
	node.Type = "Op"
	node.Exp = []Num{}

	_, err := peek(t, c)

	for err == nil {
		node.Exp.add(parseExp(t, c))
	}
	fmt.Println(err)
	return node
}

func parseExp(t []string, c *int) (*Num, *Node) {
	num, _ := peek(t, c)
	match, _ := regexp.MatchString(`\d`, num)
	if match {
		return parseNum(t, c), nil
	} else {
		return nil, parseOp(t, c)
	}

}

// *int means pointer to an int value
func peek(t []string, c *int) (string, error) {
	if *c > len(t) {
		return "-1", errors.New("index out of bounds")
	}
	return t[*c], nil
}

func consume(t []string, c *int) string {
	*c++
	return t[*c]
}
