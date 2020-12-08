package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

type groupAnswers map[rune]bool

func main() {
	lines, _ := file.GetLines("../answers.txt")

	groups := make([]groupAnswers,0,100)

	current := make(groupAnswers)
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups,current)
			current = make(groupAnswers)
		}
		for _, r := range line {
			current[r] = true
		}
	}
	groups = append(groups,current)

	total := 0
	for _, g := range groups {
		total += len(g)
	}

	fmt.Print(total)

}
