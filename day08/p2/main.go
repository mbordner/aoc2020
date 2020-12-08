package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

type cmd struct {
	op    string
	value int
}

type computer struct {
	acc int
	ptr int
}

func (comp *computer) run(program []cmd) (int, error) {
	counts := make([]int, len(program), len(program))

	for {
		if comp.ptr == len(program) {
			return 0, nil
		} else if counts[comp.ptr] > 0 {
			return -1, errors.New("infinite loop")
		}
		counts[comp.ptr]++

		c := program[comp.ptr]

		switch c.op {
		case "nop":
		case "acc":
			comp.acc += c.value
		case "jmp":
			comp.ptr += c.value
			continue
		}

		comp.ptr++
	}
}

var (
	reLine = regexp.MustCompile(`(\w{3})\s([\+|\-]\d+)`)
)

func main() {

	lines, _ := file.GetLines("../program.txt")

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

	var c *computer

	for i := 0; i < len(program); i++ {
		if program[i].op == "jmp" || program[i].op == "nop" {
			last := program[i].op
			if program[i].op == "jmp" {
				program[i].op = "nop"
			} else {
				program[i].op = "jmp"
			}

			c = &computer{}
			code, _ := c.run(program)
			if code == 0 {
				fmt.Println("bug found at line ",i+1)
				break
			}
			program[i].op = last
		}
	}


	fmt.Println(c.acc)
}


