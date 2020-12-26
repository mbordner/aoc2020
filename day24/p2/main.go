package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day24/tile"
)

func main() {
	floor := tile.NewTiles()
	lines, _ := file.GetLines("../input.txt")

	black := 0

	for i := range lines {
		dirs := tile.Parse(lines[i])
		sdirs := tile.Simplify(dirs)
		t := floor.GetTile(sdirs)
		t.Toggle()
		if t.GetState() == tile.Black {
			black++
		} else {
			black--
		}
	}

	fmt.Println("Initial State Black Count: ", black)

	tiles := floor.GetTiles()
	if len(tiles) > 0 {

	}

	days := 100

	for d := 0; d < days; d++ {
		black = process(floor)
		fmt.Println("Day ", d+1, ": ", black)
	}

}

func process(floor *tile.Tiles) int {
	tiles := floor.GetTiles()
	for _, t := range tiles {
		if t.GetState() == tile.Black {
			t.GetAdjacentStateTiles(tile.White) // activate adjacent white tiles to black
		}
	}
	tiles = floor.GetTiles() // collect any newly activated white tiles
	toggles := make([]string, 0, len(tiles))
	for _, t := range tiles {
		atiles := t.GetAdjacentStateTiles(tile.Black)
		switch t.GetState() {
		case tile.White:
			if len(atiles) == 2 {
				toggles = append(toggles, t.GetID())
			}
		case tile.Black:
			if len(atiles) == 0 || len(atiles) > 2 {
				toggles = append(toggles, t.GetID())
			}
		}
	}
	for _, toggle := range toggles {
		t := floor.GetTile(tile.Simplify(tile.Parse(toggle)))
		t.Toggle()
	}
	black := 0
	tiles = floor.GetTiles()
	for _, t := range tiles {
		if t.GetState() == tile.Black {
			black++
		}
	}
	return black
}
