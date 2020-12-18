package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/expression"
	"github.com/mbordner/aoc2020/common/file"
)

var (
	o = map[string]int{"+": 2, "*": 1}
)

func main() {
	//t := expression.NewParser(`1 + 2 * 3 + 4 * 5 + 6`, func(op1, op2 string) int {
	//	return o[op1] - o[op2]
	//})
	//fmt.Println(t.Eval())

	var sum int64
	lines, _ := file.GetLines("../input.txt")
	for i := range lines {
		p := expression.NewParser(lines[i], func(op1, op2 string) int {
			return o[op1] - o[op2]
		})
		sum += p.Eval()
	}

	fmt.Println(sum)
}
