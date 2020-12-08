package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

type seat struct {
	row int
	col int
}

func (s *seat) id() int {
	return s.row * 8 + s.col
}

func main() {
	seats := make(map[int]seat)

	lines, _ := file.GetLines("../seats1.txt")


	for i := range lines {
		s := getSeat(lines[i])
		seats[s.id()] = s
	}

	empty := make(map[int]seat)

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			id := i * 8 + j
			if !seatTaken(seats,id) {
				empty[id] = seat{row:i,col:j,}
			}
		}
	}

	for id, s := range empty {
		if s.row != 0 && s.row != 127 {
			if seatTaken(seats,id-1) && seatTaken(seats,id+1) {
				fmt.Println(id)
			}
		}
	}

}

func seatTaken(seats map[int]seat, id int) bool {
	if _, ok := seats[id]; ok {
		return true
	}
	return false
}

func getSeat(bsp string) seat {
	s := seat{}

	p1 := 0
	p2 := 128

	for _, l := range bsp[0:7] {
		t := (p2 - p1) / 2 + p1
		if l == 'F' {
			p2 = t
		} else if l == 'B' {
			p1 = t
		}
	}

	s.row = p1

	p1 = 0
	p2 = 8

	for _, l := range bsp[7:] {
		t := (p2 - p1) / 2 + p1
		if l == 'L' {
			p2 = t
		} else if l == 'R' {
			p1 = t
		}
	}

	s.col = p1

	return s
}
