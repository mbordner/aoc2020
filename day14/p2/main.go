package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reMask = regexp.MustCompile(`^mask\s*=\s([X10]{36})$`)
	reMem  = regexp.MustCompile(`^mem\[(\d+)\]\s*=\s*(\d+)$`)
)

type memval uint64
type memvals []memval

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

func (m memvals) toValues() []uint64 {
	values := make([]uint64, len(m), len(m))
	for i := range m {
		values[i] = uint64(m[i])
	}
	return values
}

func main() {
	input := getInputData()

	var mask string
	memory := make(map[uint64]uint64)

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

			addresses := getAddresses(mask, uint64(address))

			for _, v := range addresses {
				memory[v] = uint64(value)
			}
		}
	}

	total := uint64(0)
	for _, v := range memory {
		total += uint64(v)
	}
	fmt.Println(total)
}

func getAddresses(mask string, address uint64) []uint64 {
	xs := strings.Count(mask, "X")

	if xs > 0 {
		variants := 1
		for ; xs > 0; xs-- {
			variants = variants << 1
		}
		addresses := make(memvals, variants, variants)

		for i := range addresses {
			addresses[i] = memval(address)
		}

		x := 1

		for p, b := range mask {
			switch b {
			case '1':
				for i := range addresses {
					addresses[i].set(uint(p))
				}
			case 'X':
				i := 0
				for i < len(addresses) {
					for j := 0; j < x; j++ {
						addresses[i].set(uint(p))
						i++
					}
					for j := 0; j < x; j++ {
						addresses[i].clear(uint(p))
						i++
					}
				}
				x <<= 1
			}
		}

		return addresses.toValues()
	}

	return []uint64{address}
}

func getInputData() []string {
	lines, _ := file.GetLines("../input.txt")
	return lines
}
