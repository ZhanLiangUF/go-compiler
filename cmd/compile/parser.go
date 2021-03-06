package main

import (
	"errors"
	"regexp"
)

func parse(t []string) (*Node, *Num) {
	var c *int
	c = new(int)
	*c = 0

	a, b := parseExp(t, c)
	return a, b
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

	var exp []ast

	node.Val = consume(t, c)
	node.Type = "Op"
	node.Exp = exp

	// not nesting nodes here because there will never be a time when error is nil

	for peekHelper(peek(t, c)) {
		a, b := parseExp(t, c)
		if a != nil {
			node.Exp = append(node.Exp, a)
		} else {
			node.Exp = append(node.Exp, b)
		}
	}
	return node
}

// helper function for recursive while loop
func peekHelper(s string, err error) bool {
	if err != nil {
		return false
	}
	return true
}

func parseExp(t []string, c *int) (*Node, *Num) {
	num, _ := peek(t, c)
	match, _ := regexp.MatchString(`\d`, num)
	if match {
		return nil, parseNum(t, c)
	} else {
		return parseOp(t, c), nil
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
	s := t[*c]
	*c++
	return s
}
