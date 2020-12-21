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
	g.OptimizeTree()

	valid := make(map[string]bool)

	for _, s := range g.GetValues("0") {
		valid[s] = true
	}

	count := 0

	for l := i + 1; l < len(lines); l++ {
		if _, ok := valid[lines[l]]; ok {
			count++
		}
	}

	fmt.Println(count)

}
