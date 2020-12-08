package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
)

type cmd struct {
	op    string
	value int
}

var (
	reLine = regexp.MustCompile(`(\w{3})\s([\+|\-]\d+)`)
)

func main() {

	lines, _ := file.GetLines("../program.txt")
	ptr := 0
	acc := 0

	program := make([]cmd, 0, 200)
	for i := range lines {
		if matches := reLine.FindStringSubmatch(lines[i]); len(matches) == 3 {
			c := cmd{}
			c.op = matches[1]
			tmp, _ := strconv.Atoi(matches[2])
			c.value = tmp
			program = append(program, c)
		}
	}

	counts := make([]int, len(program), len(program))

	for {
		if counts[ptr] > 0 {
			break
		}
		counts[ptr]++

		c := program[ptr]

		switch c.op {
		case "nop":
		case "acc":
			acc += c.value
		case "jmp":
			ptr += c.value
			continue
		}

		ptr++
	}

	fmt.Println(acc)
}
