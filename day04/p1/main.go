package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
)

var (
	reToken  = regexp.MustCompile(`((\w{3}):([^\s]+))`)
	required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // cid is not required, it's not in North Pole Creds
)

func main() {
	lines, _ := file.GetLines("../passports1.txt")

	passports := getPassports(lines)

	valid := 0

	for i := range passports {
		if isValid(passports[i]) {
			valid++
		}
	}

	fmt.Println(valid)
}

func isValid(passport map[string]string) bool {
	for _, field := range required {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func getPassports(lines []string) []map[string]string {

	passports := make([]map[string]string, 0, 100)

	current := make(map[string]string)

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			passports = append(passports, current)
			current = make(map[string]string)
		}

		matches := reToken.FindAllStringSubmatch(lines[i], -1)

		for _, match := range matches {
			current[match[2]] = match[3]
		}

	}

	passports = append(passports, current)

	return passports
}
