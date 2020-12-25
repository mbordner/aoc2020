package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day22/combat"
)

func main() {
	lines, _ := file.GetLines("../input.txt")
	game := combat.NewGame(lines)

	fmt.Println(game.PlayGame())
}
