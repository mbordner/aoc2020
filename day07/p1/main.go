package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
)

type bagsT map[string]bagT

type bagT struct {
	name        string
	contains    map[string]bagAmount
	containedBy map[string]bagT
}

func NewBag(n string) bagT {
	b := bagT{}
	b.name = n
	b.contains = make(map[string]bagAmount)
	b.containedBy = make(map[string]bagT)
	return b
}

type bagAmount struct {
	count int
	name  string
}

var (
	reContains = regexp.MustCompile(`(?:(?:\s)(\d+)\s(\w+\s\w+)\sbags?(?:,|\.))`)
	reBag      = regexp.MustCompile(`(.*)\sbags contain`)
)

func main() {
	bags := make(bagsT)

	lines, _ := file.GetLines("../rules.txt")

	for i := range lines {
		var bagName string
		if matches := reBag.FindStringSubmatch(lines[i]); len(matches) == 2 {
			bagName = matches[1]
		}
		var bag bagT
		if b, ok := bags[bagName]; ok {
			bag = b
		} else {
			bag = NewBag(bagName)
			bags[bagName] = bag
		}

		matches := reContains.FindAllStringSubmatch(lines[i], -1)
		if len(matches) > 0 {
			for _, match := range matches {
				amount, _ := strconv.Atoi(match[1])
				name := match[2]

				var containedBag bagT
				if b, ok := bags[name]; ok {
					containedBag = b
				} else {
					containedBag = NewBag(name)
					bags[name] = containedBag
				}

				ba := bagAmount{
					count: amount,
					name:  name,
				}

				bag.contains[name] = ba
				containedBag.containedBy[bagName] = bag
			}
		}
	}

	sg := bags["shiny gold"]

	containedBy := make(bagsT)

	fmt.Println(getContainedByCount(sg, &containedBy))
	fmt.Println(len(containedBy))
}

func getContainedByCount(bag bagT, containedBy *bagsT) int {
	count := len(bag.containedBy)
	for n, b := range bag.containedBy {
		(*containedBy)[n] = b
		count += getContainedByCount(b, containedBy)
	}

	return count

}
