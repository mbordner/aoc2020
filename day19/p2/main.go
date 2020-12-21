package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day19/parser"
)

func main() {
	lines, _ := file.GetLines("../input.txt")

	i := 0
	for len(lines[i]) > 0 {
		i++
	}

	g := parser.NewGrammar(lines[0:i])

	g.UpdateRule("8", "42 | 42 8")
	g.UpdateRule("11", "42 31 | 42 11 31")

	g.OptimizeTree()

	maxLength := 0
	for l := i + 1; l < len(lines); l++ {
		if len(lines[l]) > maxLength {
			maxLength = len(lines[l])
		}
	}

	//fmt.Println(g.IsValidFor("0", "babbaaaaababbbababbbbbbb"))

	count := 0

	matches := make([]string, 0, 400)

	for l := i + 1; l < len(lines); l++ {
		if g.IsValidFor("0", lines[l]) {
			matches = append(matches, lines[l])
			count++
		}
	}


	fmt.Println(count)

}
