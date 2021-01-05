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
	c := 0

	peek := peek(t, c)
	consume := consume(t, c)

	parseNum := parseNum(t, c)
}

func parseNum(t []string, c int) {
	var num *Num
	num = new(Num)
	numHolder, err := strconv.Atoi(consume(t, c))
	num.Val = numHolder
	num.Type = "Num"
	if err != nil {
		log.Fatal("Error returned from filepath.Walk:", err)
	}
	return
}

func peek(t []string, c int) string {
	return t[c]
}

func consume(t []string, c int) string {
	return t[c]
}

func main() {
	str := "mul 3 sub 2 sum 1 3 4"
	lex := lex(str)
	fmt.Println(lex)
}
