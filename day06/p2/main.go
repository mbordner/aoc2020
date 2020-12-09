package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

type groupAnswers map[rune]int

func main() {
	lines, _ := file.GetLines("../answers.txt")

	groups := make([]groupAnswers, 0, 100)
	people := make([]int, 0, 100)

	current := make(groupAnswers)

	count := 0
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups, current)
			current = make(groupAnswers)
			people = append(people, count)
			count = 0
			continue
		}
		count++
		for _, r := range line {
			current[r]++
		}
	}
	groups = append(groups, current)
	people = append(people, count)

	totals := make([]int, len(people), len(people))
	for i, g := range groups {
		for _, t := range g {
			if t == people[i] {
				totals[i]++
			}
		}
	}

	total := 0
	for _, t := range totals {
		total += t
	}

	fmt.Print(total)

}
