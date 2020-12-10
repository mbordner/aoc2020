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
