package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
)

func main() {
	seats := getSeats()

	var changes int
	for {
		seats, changes = getNextState(seats)
		if changes == 0 {
			break
		}
	}

	occuppied := 0
	for r, row := range seats {
		for c := range row {
			if seats[r][c] == '#' {
				occuppied++
			}
		}
	}

	fmt.Println(occuppied)
}

func getNextState(seats [][]byte) (next [][]byte, changes int) {
	next = make([][]byte, len(seats), len(seats))
	for i := range seats {
		next[i] = make([]byte, len(seats[i]), len(seats[i]))
	}

	for r, row := range seats {
		for c := range row {
			n := seats[r][c]
			if seats[r][c] != '.' { // floors never change
				a := getAdjacent(seats, r, c)
				switch seats[r][c] {
				case 'L':
					if a == 0 {
						n = '#'
					}
				case '#':
					if a > 3 {
						n = 'L'
					}
				}
			}
			if n != seats[r][c] {
				changes++
			}
			next[r][c] = n
		}
	}

	return
}

func getAdjacent(seats [][]byte, row int, col int) int {
	count := 0
	if row > 0 {
		if col > 0 {
			if seats[row-1][col-1] == '#' { // upper left
				count++
			}
		}
		if seats[row-1][col] == '#' { // above
			count++
		}
		if col < len(seats[0])-1 {
			if seats[row-1][col+1] == '#' { // upper right
				count++
			}
		}
	}

	if col > 0 {
		if seats[row][col-1] == '#' { // left
			count++
		}
	}
	if col < len(seats[0])-1 {
		if seats[row][col+1] == '#' { // right
			count++
		}
	}
	if row < len(seats)-1 {
		if col > 0 {
			if seats[row+1][col-1] == '#' { // lower left
				count++
			}
		}
		if seats[row+1][col] == '#' { // below
			count++
		}
		if col < len(seats[0])-1 {
			if seats[row+1][col+1] == '#' { // lower right
				count++
			}
		}
	}
	return count
}

func getSeats() [][]byte {
	lines, _ := file.GetLines("../seats.txt")

	seats := make([][]byte, len(lines), len(lines))
	for i := range lines {
		seats[i] = []byte(lines[i])
	}

	return seats
}
