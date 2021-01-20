package main

import "strings"

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
