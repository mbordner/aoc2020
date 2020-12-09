package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"sort"
	"strconv"
)

func main() {

	numbers := getNumbers()
	for i := 25; i < len(numbers); i++ {
		if !isValid(numbers[i], getPrev(i, 25, numbers)) {
			fmt.Println(numbers[i])
			break
		}
	}
}

func getPrev(i int, c int, numbers []uint64) []uint64 {
	n := make([]uint64, c, c)
	for i, j := i-c, 0; j < c; i, j = i+1, j+1 {
		n[j] = numbers[i]
	}
	return n
}

func isValid(n uint64, p []uint64) bool {
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })

	for h := len(p) - 1; h > 0; h-- {
		if p[h] < n {
		loopj:
			for j := h; j > 0; j-- {
				for i := 0; i < j; i++ {
					sum := p[i] + p[j]
					if sum == n {
						return true
					} else if sum > n {
						continue loopj
					}
				}
			}
		}
	}

	return false
}

func getNumbers() []uint64 {
	lines, _ := file.GetLines("../xmas.txt")
	numbers := make([]uint64, len(lines), len(lines))

	for i := range lines {
		n, e := strconv.ParseUint(lines[i], 10, 64)
		if e != nil {
			panic(e)
		}
		numbers[i] = n
	}

	return numbers
}
