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

type waypointT struct {
	x float64
	y float64
}

func (w waypointT) X() int64 {
	return int64(math.Round(w.x))
}

func (w waypointT) Y() int64 {
	return int64(math.Round(w.y))
}

type shipT struct {
	x int64
	y int64
	w waypointT
}

func (s *shipT) manhattanDistance() int64 {
	return int64(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *shipT) rotateWaypoint(d int64) {
	radians := math.Pi * float64(d) / float64(180)

	x := (s.w.x * math.Cos(radians)) -
		(s.w.y * math.Sin(radians))

	y := (s.w.y * math.Cos(radians)) +
		(s.w.x * math.Sin(radians))

	s.w.x = x
	s.w.y = y
}

func (s *shipT) process(c *cmdT) {
	switch c.cmd {
	case "N":
		s.w.y += float64(c.value)
	case "S":
		s.w.y -= float64(c.value)
	case "E":
		s.w.x += float64(c.value)
	case "W":
		s.w.x -= float64(c.value)
	case "L":
		s.rotateWaypoint(c.value)
	case "R":
		s.rotateWaypoint(360 - c.value)
	case "F":
		for i := int64(0); i < c.value; i++ {
			s.x += s.w.X()
			s.y += s.w.Y()
		}
	}
}

func NewShipT() *shipT {
	s := shipT{}
	s.w.x = 10
	s.w.y = 1
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
