package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/expression"
	"github.com/mbordner/aoc2020/common/file"
)

func main() {
	//p := expression.NewParser(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, expression.CompareOperator)
	//fmt.Println(p.Eval())

	var sum int64
	lines, _ := file.GetLines("../input.txt")
	for i := range lines {
		p := expression.NewParser(lines[i], expression.CompareOperator)
		sum += p.Eval()
	}

	fmt.Println(sum)
}
