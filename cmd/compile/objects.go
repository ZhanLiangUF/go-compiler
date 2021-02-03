package main

type Num struct {
	Val  string
	Type string
}

type Node struct {
	Type string
	Val  string
	Exp  []ast
}

func (n Num) getType() string {
	return n.Type
}

func (n Node) getType() string {
	return n.Type
}

type ast interface {
	getType() string
}
