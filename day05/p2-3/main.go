package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"strconv"
	"strings"
)

func main() {
	r := strings.NewReplacer("B", "1", "F", "0", "R", "1", "L", "0")

	lines, _ := file.GetLines("../seats1.txt")

	taken := make(map[int]string)

	for i := range lines {
		binStr := r.Replace(lines[i])
		id, err := strconv.ParseInt(binStr, 2, 32)
		if err != nil {
			panic(err)
		}
		taken[int(id)] = lines[i]
	}

	for i := 1; i < 127; i++ { // skip first and last rows
		for j := 0; j < 8; j++ { // go through the columns
			id := i*8 + j
			if _, isTaken := taken[id]; !isTaken { // is this seat empty
				if _, isPreviousTaken := taken[id-1]; isPreviousTaken { // is the prev id taken
					if _, isNextTaken := taken[id+1]; isNextTaken { // is the next id taken
						fmt.Println(id) // my seat
					}
				}
			}
		}
	}
}
