package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Num struct {
	Val  int
	Type string
	expr []Exp
}

type Exp struct {
	Type string
	Val  int
}

type Node struct {
	Type string
	Val  int
	Exp  []Num
}

func lex(s string) []string {
	entries := strings.Split(s, " ")
	lex := make([]string, 0)
	for _, e := range entries {
		eTrim := strings.Trim(e, " ")
		if len(eTrim) > 0 {
			lex = append(lex, eTrim)
		}
	}
	return lex
}

func parse(t []string) {
	var c *int
	c = new(int)
	*c = 0

	// peek := peek(t, c)
	consume := consume(t, c)

	parseNum := parseNum(consume)
	fmt.Println(parseNum)
}

func parseNum(consume string) *Num {
	var num *Num
	num = new(Num)
	numHolder, err := strconv.Atoi(consume)
	num.Val = numHolder
	num.Type = "Num"
	if err != nil {
		log.Fatal("Error returned from strconv.Atoi:", err)
	}
	return num
}

// func parseOp(t []string, c *int) *Node {
// 	var node *Node
// 	node = new(Node)
// 	numHolder, err := strconv.Atoi(consume(t, c))
// }

// *int means pointer to an int value
func peek(t []string, c *int) string {
	return t[*c]
}

func consume(t []string, c *int) string {
	*c++
	return t[*c]
}

func main() {
	str := "mul 3 sub 2 sum 1 3 4"
	lex := lex(str)
	parse(lex)
	fmt.Println(lex)
}
