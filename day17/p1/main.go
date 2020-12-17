package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array"
	"github.com/mbordner/aoc2020/common/file"
)

type activeState map[string]bool

func main() {

	active := make(activeState)

	lines, _ := file.GetLines("../input.txt")
	size := len(lines)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if lines[y][x] == '#' {
				active[getKey(x, y, 0)] = true
			}
		}
	}

	cycles := 6

	for c := 1; c <= cycles; c++ {
		next := cloneState(active)

		for x := -c; x <= size+c; x++ {
			for y := -c; y <= size+c; y++ {
				for z := -c; z <= c; z++ {
					current := getKey(x, y, z)
					neighbors := getNeighbors(current)
					activeNeighbors := getActiveNeighborsCount(active, neighbors)

					if isActive(active, current) {
						if activeNeighbors < 2 || activeNeighbors > 3 {
							setInactive(&next, current)
						}
					} else {
						if activeNeighbors == 3 {
							setActive(&next, current)
						}
					}

				}
			}
		}

		active = next
	}

	fmt.Println(len(active))
}

func cloneState(s activeState) activeState {
	n := make(activeState)
	for k := range s {
		n[k] = true
	}
	return n
}

func getActiveNeighborsCount(s activeState, neighbors []string) int {
	count := 0
	for i := range neighbors {
		if isActive(s, neighbors[i]) {
			count++
		}
	}
	return count
}

func setActive(s *activeState, key string) {
	(*s)[key] = true
}

func setInactive(s *activeState, key string) {
	delete(*s, key)
}

func isActive(s activeState, key string) bool {
	if _, ok := s[key]; ok {
		return true
	}
	return false
}

func getKey(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func getNeighbors(s string) []string {
	a := array.ToIntArray(s)

	n := make([][]int, 0, 26)

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x == 0 && y == 0 && z == 0) {
					t := []int{a[0] + x, a[1] + y, a[2] + z}
					n = append(n, t)
				}
			}
		}
	}

	ss := make([]string, len(n), len(n))

	for i := range n {
		ss[i] = getKey(n[i][0], n[i][1], n[i][2])
	}

	return ss
}
