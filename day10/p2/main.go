package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/pkg/errors"
	"sort"
	"strconv"
)

func main() {
	adapters := getAdapters()
	sort.Ints(adapters)

	diff1, diff3 := 0, 0

	last := 0
	for i := 0; i < len(adapters); i++ {
		diff := adapters[i] - last
		last = adapters[i]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		} else {
			panic(errors.New("wtf"))
		}
	}

	diff3++

	fmt.Println(diff1, diff3, diff1*diff3)

	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	shortest := make([]int, 0, len(adapters))
	shortest = append(shortest, 0)
	for i := 0; i < len(adapters)-1; {
		max := i
		for j := i + 1; j < len(adapters) && j < i+4; j++ {
			diff := adapters[j] - adapters[i]
			if diff <= 3 {
				max = j
			} else {
				break
			}
		}
		shortest = append(shortest, adapters[max])
		if i == max {
			i++
		} else {
			i = max
		}
	}
	fmt.Println(adapters)
	fmt.Println(shortest)

	pathsFromMap := make(map[int]int)
	fmt.Println(pathsFrom(0, adapters, pathsFromMap))

}

func pathsFrom(i int, adapters []int, pathsFromMap map[int]int) int {
	if i == len(adapters)-1 {
		return 1
	}
	if p, ok := pathsFromMap[i]; ok {
		return p
	}
	paths := 0
	for j := i + 1; j < len(adapters) && j < i+4; j++ {
		diff := adapters[j] - adapters[i]
		if diff <= 3 {
			paths += pathsFrom(j, adapters, pathsFromMap)
		}
	}
	pathsFromMap[i] = paths
	return paths
}

func getAdapters() []int {
	lines, _ := file.GetLines("../adapters.txt")

	adapters := make([]int, 0, 200)
	for i := range lines {
		tmp, _ := strconv.Atoi(lines[i])
		adapters = append(adapters, tmp)
	}

	return adapters
}
