package main

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
