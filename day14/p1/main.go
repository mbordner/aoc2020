package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
)

var (
	reMask = regexp.MustCompile(`^mask\s*=\s([X10]{36})$`)
	reMem  = regexp.MustCompile(`^mem\[(\d+)\]\s*=\s*(\d+)$`)
)

type memval uint64

func (v *memval) set(pos uint) {
	val := uint64(*v) | (1 << pos)
	*v = memval(val)
}

func (v *memval) clear(pos uint) {
	mask := uint64(*v) ^ (1 << pos)
	val := uint64(*v) & uint64(mask)
	*v = memval(val)
}

func (v memval) has(pos uint) bool {
	val := int64(v) & (1 << pos)
	return val > 0
}

func main() {
	input := getInputData()

	var mask string
	memory := make(map[int64]memval)

	for i := range input {
		if reMask.MatchString(input[i]) {
			matches := reMask.FindStringSubmatch(input[i])
			mask = matches[1]

			runes := []rune(mask)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			mask = string(runes)
		} else if reMem.MatchString(input[i]) {
			matches := reMem.FindStringSubmatch(input[i])
			address, _ := strconv.ParseInt(matches[1], 10, 64)
			value, _ := strconv.ParseInt(matches[2], 10, 64)

			m := memval(value)

			if mask != "" {
				for p, b := range mask {
					switch b {
					case '1':
						m.set(uint(p))
					case '0':
						m.clear(uint(p))
					}
				}
			}

			memory[address] = m
		}
	}

	total := uint64(0)
	for _, v := range memory {
		total += uint64(v)
	}
	fmt.Println(total)
}

func getInputData() []string {
	lines, _ := file.GetLines("../input.txt")
	return lines
}
