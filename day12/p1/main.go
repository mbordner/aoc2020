package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"math"
	"regexp"
	"strconv"
)

var (
	reCmd = regexp.MustCompile(`([ENWSLRF])(\d+)`)
)

type cmdT struct {
	cmd   string
	value int64
}

type shipT struct {
	x int64
	y int64
	w int64
}

func (s *shipT) process(c *cmdT) {
	switch c.cmd {
	case "N":
		s.y += c.value
	case "S":
		s.y -= c.value
	case "E":
		s.x += c.value
	case "W":
		s.x -= c.value
	case "L":
		s.w += c.value
		s.w = s.w % 360
	case "R":
		s.w += 360 - c.value
		s.w = s.w % 360
	case "F":
		nc := cmdT{value: c.value}
		switch s.w {
		case 0:
			nc.cmd = "E"
		case 90:
			nc.cmd = "N"
		case 180:
			nc.cmd = "W"
		case 270:
			nc.cmd = "S"
		}
		s.process(&nc)
	}
}

func (s *shipT) manhattanDistance() int64 {
	return int64(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func NewShipT() *shipT {
	s := shipT{}
	return &s
}

func main() {
	cmds := getCmds()

	s := NewShipT()

	for _, c := range cmds {
		s.process(c)
	}

	fmt.Println(s.x, s.y, s.manhattanDistance())
}

func getCmds() []*cmdT {
	lines, _ := file.GetLines("../route.txt")

	cmds := make([]*cmdT, len(lines), len(lines))

	for i := range lines {
		matches := reCmd.FindStringSubmatch(lines[i])
		c := cmdT{}
		c.cmd = matches[1]

		tmp, _ := strconv.ParseInt(matches[2], 10, 64)
		c.value = tmp

		cmds[i] = &c
	}

	return cmds
}
