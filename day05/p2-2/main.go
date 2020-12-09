package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"strconv"
	"strings"
)

func main() {
	lines, _ := file.GetLines("../seats1.txt")
	/*
		strs := []string{"BFFFBBFRRR","FFFBBBFRRR","BBFFBBFRLL"}
		for i := range strs {
			id := bspToID(strs[i])

			fmt.Println(id)

			bsp := idToBSP(id)

			fmt.Println(bsp)
		}
	*/
	taken := make(map[int]string)

	for i := range lines {
		taken[bspToID(lines[i])] = lines[i]
	}

	for i := 1; i < 127; i++ {
		for j := 0; j < 8; j++ {
			id := i*8 + j
			if _, ok := taken[id]; !ok {
				if _, taken1 := taken[id-1]; taken1 {
					if _, taken2 := taken[id+1]; taken2 {
						fmt.Println(id)
					}
				}
			}

		}
	}
}

func bspToID(bsp string) int {
	binstr := strings.ReplaceAll(bsp, "B", "1")
	binstr = strings.ReplaceAll(binstr, "F", "0")
	binstr = strings.ReplaceAll(binstr, "R", "1")
	binstr = strings.ReplaceAll(binstr, "L", "0")
	id, _ := strconv.ParseInt(binstr, 2, 32)
	return int(id)
}

func idToBSP(id int) string {
	//binstr := strconv.FormatInt(int64(id), 2)
	binstr := fmt.Sprintf("%010b", id)
	s1 := binstr[0:7]
	s2 := binstr[7:]
	s1 = strings.ReplaceAll(s1, "1", "B")
	s1 = strings.ReplaceAll(s1, "0", "F")
	s2 = strings.ReplaceAll(s2, "1", "R")
	s2 = strings.ReplaceAll(s2, "0", "L")
	return s1 + s2
}
