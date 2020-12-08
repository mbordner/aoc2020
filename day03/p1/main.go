package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

var (
	x = 0
	y = 0
)

func main() {
	treeMap, _ := file.GetLines("../map1.txt")

	treeCount := 0
	rows := len(treeMap)
	cols := len(treeMap[0])

	for y < rows {
		x += 3
		if x >= cols {
			x = x - cols
		}
		y += 1
		if y < rows {
			if treeMap[y][x] == '#' {
				treeCount++
			}
		}
	}

	fmt.Println(treeCount)
}
