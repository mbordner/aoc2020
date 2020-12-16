package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
	"strings"
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

func (c *class) rangeNotCovered(min, max int) []int {
	r := make([]int, 0, max-min+1)
	for i := min; i <= max; i++ {
		covered := false
		for _, rs := range c.ranges {
			if i >= rs.min && i <= rs.max {
				covered = true
				break
			}
		}
		if !covered {
			r = append(r, i)
		}
	}
	return r
}

func newClass(s string) *class {
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

	return &c
}

type classes struct {
	min         int
	max         int
	classValues map[int][]*class
	classes     map[string]*class
}

func (cs *classes) getNotCovered() map[int][]*class {
	values := make(map[int][]*class)

	for _, c := range cs.classes {
		notCovered := c.rangeNotCovered(cs.min, cs.max)
		for _, i := range notCovered {
			if _, ok := values[i]; !ok {
				values[i] = make([]*class, 0, 25)
			}
			values[i] = append(values[i], c)
		}
	}

	return values
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
	c := newClass(s)
	cs.classes[c.name] = c

	for i := 0; i < len(c.ranges); i++ {
		r := c.ranges[i]
		for j := r.min; j <= r.max; j++ {
			cs.recordPossibleForClass(j, c)
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
	cs, ticket, others := getInput()

	tmp := others

	others = make([][]int, 0, len(others))

	for i := range tmp {
		valid := true
		for _, j := range tmp[i] {
			if !cs.isValidForClass(j) {
				valid = false
				break
			}
		}
		if valid {
			others = append(others, tmp[i])
		}
	}

	numTicketFields := len(others[0])

	possibles := make([]map[string]*class, numTicketFields, numTicketFields)
	notpossibles := make([]map[string]*class, numTicketFields, numTicketFields)
	for j := range possibles {
		possibles[j] = make(map[string]*class)
		notpossibles[j] = make(map[string]*class)
	}

	notCovered := cs.getNotCovered()

	for j := 0; j < numTicketFields; j++ {
		for i := range others {
			possibleClasses := cs.classValues[others[i][j]]
			for _, c := range possibleClasses {
				possibles[j][c.name] = c
			}
			if _, ok := notCovered[others[i][j]]; ok {
				for _, c := range notCovered[others[i][j]] {
					notpossibles[j][c.name] = c
				}
			}
		}
	}

	p, np := getClassKeys(possibles, notpossibles)

	for !allFieldsIdentified(p) {
		for i := range p {
			if len(p[i]) == 1 {
				for j := range p {
					if i != j {
						np[j] = addField(np[j], p[i][0])
					}
				}
			}
		}
		p = removeImpossibles(p, np)
	}

	answer := 1

	for i := range p {
		if strings.HasPrefix(p[i][0], "departure") {
			answer *= ticket[i]
		}
	}

	fmt.Println(answer)
}

func addField(fs []string, f string) []string {
	if searchFields(fs, f) == -1 {
		fs = append(fs, f)
	}
	return fs
}

func searchFields(fs []string, f string) int {
	result := -1
	for i := range fs {
		if fs[i] == f {
			result = i
			break
		}
	}
	return result
}

func allFieldsIdentified(p [][]string) bool {
	for i := range p {
		if len(p[i]) > 1 {
			return false
		}
	}
	return true
}

func removeImpossibles(p, np [][]string) [][]string {
	ps := make([][]string, len(p), len(p))

	for i := 0; i < len(p); i++ {
		ps[i] = make([]string, 0, len(p[i]))
		for j := 0; j < len(p[i]); j++ {
			index := searchFields(np[i], p[i][j])
			if index == -1 {
				ps[i] = append(ps[i], p[i][j])
			}
		}
	}
	return ps
}

func getClassKeys(possibles, notpossibles []map[string]*class) ([][]string, [][]string) {
	p := make([][]string, len(possibles), len(possibles))
	np := make([][]string, len(notpossibles), len(notpossibles))

	for f := range possibles {
		p[f] = make([]string, 0, len(possibles[f]))
		for k := range possibles[f] {
			p[f] = append(p[f], k)
		}
		np[f] = make([]string, 0, len(notpossibles[f]))
		for k := range notpossibles[f] {
			np[f] = append(np[f], k)
		}
	}

	return p, np
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
