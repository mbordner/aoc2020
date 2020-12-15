package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	input = "6,4,12,1,20,0,16"
)

type memory struct {
	nums map[uint64][]uint64
}

func newMemory() *memory {
	m := memory{}
	m.nums = make(map[uint64][]uint64)
	return &m
}

func (m *memory) add(v uint64, t uint64) []uint64 {
	if _, ok := m.nums[v]; !ok {
		m.nums[v] = []uint64{t}
	} else {
		m.nums[v] = []uint64{t, m.nums[v][0]}
	}
	return m.nums[v]
}

func main() {

	nums := getStart(input)

	m := newMemory()

	var turn uint64
	var next uint64

	for i := range nums {
		turn = uint64(i + 1)
		next = nums[i]
		m.add(next, turn)
	}

	next = 0

	for turn < 2020 {
		turn++

		fmt.Println("turn ", turn, " is saying ", next)
		turns := m.add(next, turn)
		if len(turns) > 1 {
			next = turns[0] - turns[1]
		} else {
			next = 0
		}
	}

}

func getStart(s string) []uint64 {
	tokens := strings.Split(s, ",")
	start := make([]uint64, len(tokens), len(tokens))
	for i := range tokens {
		val, _ := strconv.ParseUint(tokens[i], 10, 32)
		start[i] = val
	}
	return start
}
