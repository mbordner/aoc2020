package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array/bytes"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/mbordner/aoc2020/day20/tile"
)

var (
	seaMonster = [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '#', 0},
		{'#', 0, 0, 0, 0, '#', '#', 0, 0, 0, 0, '#', '#', 0, 0, 0, 0, '#', '#', '#'},
		{0, '#', 0, 0, '#', 0, 0, '#', 0, 0, '#', 0, 0, '#', 0, 0, '#', 0, 0, 0},
	}
)

func main() {
	lines, _ := file.GetLines("../input.txt")
	tiles := tile.NewTiles(lines)

	bitmapArrangement := tiles.GetBitmapArrangement()
	tileSqLen, _ := bitmapArrangement[0][0].GetDimensions()

	tileSqLen -= 2

	bitmapLen := tileSqLen * len(bitmapArrangement)

	bitmap := make([][]byte, bitmapLen, bitmapLen)
	for i := range bitmap {
		bitmap[i] = make([]byte, bitmapLen, bitmapLen)
	}

	for j := 0; j < len(bitmapArrangement); j++ {
		for i := 0; i < len(bitmapArrangement); i++ {
			bytes.Copy2D(bitmap, bitmapArrangement[j][i].RemoveBorders(), j*tileSqLen, i*tileSqLen, 0, 0, tileSqLen, tileSqLen)
		}
	}

	var smPos []bytes.Pos
	var smBM [][]byte

	smCount := 0

	for i := 0; i < 3; i++ {
		bm := bytes.Clone2D(bitmap)

		switch i {
		case 1:
			bm = bytes.Flip(bytes.Horizontal, bm)
		case 2:
			bm = bytes.Flip(bytes.Vertical, bm)
		}

		pos := bytes.FindMasked(bm, seaMonster)
		if len(pos) > smCount {
			smCount = len(pos)
			smBM = bm
			smPos = pos
		}

		for j := 0; j < 3; j++ {
			bm = bytes.Rotate(bm)
			pos = bytes.FindMasked(bm, seaMonster)
			if len(pos) > smCount {
				smCount = len(pos)
				smBM = bm
				smPos = pos
			}
		}
	}

	for j := 0; j < len(bitmap); j++ {
		fmt.Print("\n")
		for i := 0; i < len(bitmap[0]); i++ {
			if bitmap[j][i] == 0 {
				fmt.Print("O")
			} else {
				fmt.Print(string(bitmap[j][i]))
			}
		}
	}
	fmt.Print("\n\n----\n\n")

	for _, pos := range smPos {
		bytes.ApplyMask(smBM, seaMonster, pos)
	}

	count := 0
	for j := 0; j < len(smBM); j++ {
		for i := 0; i < len(smBM[0]); i++ {
			if smBM[j][i] == '#' {
				count++
			}
		}
	}

	for j := 0; j < len(smBM); j++ {
		fmt.Print("\n")
		for i := 0; i < len(smBM[0]); i++ {
			if smBM[j][i] == 0 {
				fmt.Print("O")
			} else {
				fmt.Print(string(smBM[j][i]))
			}
		}
	}
	fmt.Print("\n")

	fmt.Println(count)

}
