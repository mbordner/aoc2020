package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day24/tile"
)

func main() {
	tiles := tile.NewTiles()
	lines, _ := file.GetLines("../input.txt")

	black := 0

	for i := range lines {
		dirs := tile.Parse(lines[i])
		sdirs := tile.Simplify(dirs)
		t := tiles.GetTile(sdirs)
		t.Toggle()
		if t.GetState() == tile.Black {
			black++
		} else {
			black--
		}
	}

	fmt.Println(black)

}
