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

	occupied := 0
	for r, row := range seats {
		for c := range row {
			if seats[r][c] == '#' {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
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
					if a > 4 {
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
	// left
	for c := col - 1; c >= 0; c-- {
		if seats[row][c] != '.' {
			if seats[row][c] == '#' {
				count++
			}
			break
		}
	}

	// right
	for c := col + 1; c <= len(seats[row])-1; c++ {
		if seats[row][c] != '.' {
			if seats[row][c] == '#' {
				count++
			}
			break
		}
	}

	// above
	for r := row - 1; r >= 0; r-- {
		if seats[r][col] != '.' {
			if seats[r][col] == '#' {
				count++
			}
			break
		}
	}

	// below
	for r := row + 1; r <= len(seats)-1; r++ {
		if seats[r][col] != '.' {
			if seats[r][col] == '#' {
				count++
			}
			break
		}
	}

	// upper left
	for c, r := col-1, row-1; c >= 0 && r >= 0; c, r = c-1, r-1 {
		if seats[r][c] != '.' {
			if seats[r][c] == '#' {
				count++
			}
			break
		}
	}

	// lower left
	for c, r := col-1, row+1; c >= 0 && r <= len(seats)-1; c, r = c-1, r+1 {
		if seats[r][c] != '.' {
			if seats[r][c] == '#' {
				count++
			}
			break
		}
	}

	// upper right
	for c, r := col+1, row-1; c <= len(seats[row])-1 && r >= 0; c, r = c+1, r-1 {
		if seats[r][c] != '.' {
			if seats[r][c] == '#' {
				count++
			}
			break
		}
	}

	// lower right
	for c, r := col+1, row+1; c <= len(seats[row])-1 && r <= len(seats)-1; c, r = c+1, r+1 {
		if seats[r][c] != '.' {
			if seats[r][c] == '#' {
				count++
			}
			break
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
