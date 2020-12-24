package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day20/tile"
)

func main() {
	lines, _ := file.GetLines("../input.txt")
	tiles := tile.NewTiles(lines)

	orders := tiles.GetOrders()

	fmt.Println(orders[0].GetID() * orders[1].GetID() * orders[2].GetID() * orders[3].GetID())
}
