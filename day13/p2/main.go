package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"sort"
	"strconv"
	"strings"
)

type input struct {
	arrival uint64
	buses   []uint64
}

type bus struct {
	id     uint64
	offset int64
}

func main() {
	data := getInputData()

	buses := make([]bus, 0, len(data.buses))

	for i := range data.buses {
		if data.buses[i] != uint64(0) {
			b := bus{id: data.buses[i], offset: int64(i)}
			buses = append(buses, b)
		}
	}

	sort.SliceStable(buses, func(i, j int) bool {
		return buses[i].id > buses[j].id
	})

	i := findRepeats(buses[0].id, buses[0].id, buses)

	fmt.Println(i)
}

func findRepeats(i uint64, skip uint64, buses []bus) uint64 {
	last := uint64(0)
	for ; ; i += skip {
		offset := buses[0].offset - buses[1].offset
		if uint64(int64(i)-offset)%buses[1].id == 0 {

			foundAll := true

			for j := 1; j < len(buses)-1; j++ {
				if uint64(int64(i)-buses[j].offset)%buses[j].id != 0 {
					foundAll = false
					break
				}
			}

			if foundAll {
				return uint64(int64(i) - buses[0].offset)
			} else {
				if last != uint64(0) {
					return findRepeats(uint64(int64(i)-offset), i-last, buses[1:])
				} else {
					last = i
				}
			}
		}

	}
}

func getInputData() input {
	lines, _ := file.GetLines("../input.txt")

	arrival, _ := strconv.ParseUint(lines[0], 10, 64)

	tokens := strings.Split(lines[1], ",")

	buses := make([]uint64, len(tokens), len(tokens))
	for i := range tokens {
		if tokens[i] != "x" {
			tmp, _ := strconv.ParseUint(tokens[i], 10, 64)
			buses[i] = tmp
		}
	}

	return input{arrival: arrival, buses: buses}
}
