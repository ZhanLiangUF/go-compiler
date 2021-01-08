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

	parseNum := parseNum(t, c)
	parseOp := parseOp(t, c)

	fmt.Println(parseNum)
	fmt.Println(parseOp)

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
		// if num then parseExpr
	}
	fmt.Println(err)
	return node
}

func isNum(t []string, c *int) bool {
	num, _ := peek(t, c)
	match, _ := regexp.MatchString(`\d`, num)
	return match
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
