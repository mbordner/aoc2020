package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
)

var (
	reClass = regexp.MustCompile(`^(.*):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)$`)
)

type classRange struct {
	min int
	max int
}
type class struct {
	name   string
	ranges []classRange
}

type classes struct {
	min         int
	max         int
	classValues map[int][]*class
	classes     map[string]*class
}

func (cs *classes) isValidForClass(n int) bool {
	if _, ok := cs.classValues[n]; ok {
		return true
	}
	return false
}

func (cs *classes) recordPossibleForClass(n int, c *class) {
	if cs.min == 0 {
		cs.min = n
	} else if n < cs.min {
		cs.min = n
	}

	if cs.max == 0 {
		cs.max = n
	} else if n > cs.max {
		cs.max = n
	}

	if _, ok := cs.classValues[n]; !ok {
		cs.classValues[n] = make([]*class, 0, 25)
	}
	cs.classValues[n] = append(cs.classValues[n], c)
}

func (cs *classes) addClass(s string) {
	matches := reClass.FindStringSubmatch(s)

	c := class{name: matches[1]}
	c.ranges = make([]classRange, 0, 2)

	for i := 2; i < len(matches); i += 2 {
		min, _ := strconv.ParseInt(matches[i], 10, 32)
		max, _ := strconv.ParseInt(matches[i+1], 10, 32)
		cr := classRange{
			min: int(min),
			max: int(max),
		}
		c.ranges = append(c.ranges, cr)
	}

	cs.classes[c.name] = &c

	for i := 0; i < len(c.ranges); i++ {
		r := c.ranges[i]
		for j := r.min; j <= r.max; j++ {
			cs.recordPossibleForClass(j, &c)
		}
	}
}

func newClasses() *classes {
	cs := classes{}
	cs.classValues = make(map[int][]*class)
	cs.classes = make(map[string]*class)
	return &cs
}

func main() {
	cs, _, others := getInput()

	errorRate := 0

	for i := range others {
		for _, j := range others[i] {
			if !cs.isValidForClass(j) {
				errorRate += j
			}
		}
	}

	fmt.Println(errorRate)
}

func getInput() (*classes, []int, [][]int) {
	lines, _ := file.GetLines("../input.txt")

	cs := newClasses()

	var i int
	for lines[i] != "" {
		cs.addClass(lines[i])
		i++
	}

	i += 2
	ticket := array.ToIntArray(lines[i])

	i += 3
	others := make([][]int, 0, len(lines)-i+1)

	for i < len(lines) {
		others = append(others, array.ToIntArray(lines[i]))
		i++
	}

	return cs, ticket, others
}
