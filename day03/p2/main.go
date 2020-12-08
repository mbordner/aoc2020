package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

type slope struct {
	right int64
	down  int64
}

func main() {
	treeMap, _ := file.GetLines("../map1.txt")

	slopes := []slope{slope{1, 1}, slope{3, 1}, slope{5, 1}, slope{7, 1}, slope{1, 2}}
	treeCount := int64(0)

	for i := range slopes {
		c := getTreesForSlope(treeMap, slopes[i])
		if treeCount == int64(0) {
			treeCount = c
		} else {
			treeCount *= c
		}
	}

	fmt.Println(treeCount)
}

func getTreesForSlope(treeMap []string, s slope) int64 {
	treeCount := int64(0)
	rows := int64(len(treeMap))
	cols := int64(len(treeMap[0]))
	x := int64(0)
	y := int64(0)

	for y < rows {
		x += s.right
		if x >= cols {
			x = x - cols
		}
		y += s.down
		if y < rows {
			if treeMap[y][x] == '#' {
				treeCount++
			}
		}
	}

	return treeCount
}
