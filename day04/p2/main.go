package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"strconv"
)

var (
	reToken  = regexp.MustCompile(`((\w{3}):([^\s]+))`)
	reHeight = regexp.MustCompile(`^(\d+)(in|cm)$`)
	reHair   = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	reEye    = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	rePid    = regexp.MustCompile(`^\d{9}$`)
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

func isValid(p map[string]string) bool {
	for _, field := range required {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	if !validInt(p["byr"], 1920, 2002) {
		return false
	}

	if !validInt(p["iyr"], 2010, 2020) {
		return false
	}

	if !validInt(p["eyr"], 2020, 2030) {
		return false
	}

	if matches := reHeight.FindStringSubmatch(p["hgt"]); len(matches) == 3 {
		switch matches[2] {
		case "cm":
			if !validInt(matches[1], 150, 193) {
				return false
			}
		case "in":
			if !validInt(matches[1], 59, 76) {
				return false
			}
		default:
			return false
		}
	} else {
		return false
	}

	if !reHair.MatchString(p["hcl"]) {
		return false
	}

	if !reEye.MatchString(p["ecl"]) {
		return false
	}

	if !rePid.MatchString(p["pid"]) {
		return false
	}

	return true
}

func validInt(i string, min int, max int) bool {
	val, err := strconv.Atoi(i)
	if err == nil {
		if val >= min && val <= max {
			return true
		}
	}

	return false
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
