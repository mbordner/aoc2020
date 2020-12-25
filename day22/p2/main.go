package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day22/rcombat"
)

func main() {
	lines, _ := file.GetLines("../input.txt")
	game := rcombat.NewGame(lines)

	fmt.Println(game.PlayGame())
}
