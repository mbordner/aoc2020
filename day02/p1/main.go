package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

var (
	reRecord = regexp.MustCompile(`(\d+)\-(\d+)\s*(\w):\s*(.*)`)
)

type passwordData struct {
	min      int
	max      int
	char     byte
	password string
}

func main() {
	passwords, err := getPasswords("../data.txt")
	if err != nil {
		panic(err)
	}

	c := 0
	for _, p := range passwords {
		if isValidPassword(p) {
			c++
		}
	}

	fmt.Println("valid password count:", c)
}

func isValidPassword(p *passwordData) bool {
	c := 0
	for _, b := range []byte(p.password) {
		if b == p.char {
			c++
		}
	}
	if c >= p.min && c <= p.max {
		return true
	}
	return false
}

func getPasswords(filename string) ([]*passwordData, error) {
	lines, err := file.GetLines(filename)
	if err != nil {
		return nil, err
	}

	passwords := make([]*passwordData, len(lines), len(lines))

	for i, l := range lines {
		if matches := reRecord.FindStringSubmatch(l); len(matches) == 5 {
			p := passwordData{}
			p.char = matches[3][0]
			if tmp, err := strconv.Atoi(matches[1]); err == nil {
				p.min = tmp
			} else {
				return nil, errors.Errorf("invalid password at row %d", i)
			}
			if tmp, err := strconv.Atoi(matches[2]); err == nil {
				p.max = tmp
			} else {
				return nil, errors.Errorf("invalid password at row %d", i)
			}
			p.password = matches[4]
			passwords[i] = &p
		} else {
			return nil, errors.Errorf("invalid password at row %d", i)
		}
	}
	return passwords, nil
}
