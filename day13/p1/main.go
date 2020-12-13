package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/file"
	"strconv"
	"strings"
)

type input struct {
	arrival int64
	buses   []int64
}

func main() {
	data := getInputData()

	var preferredID, departure int64

	for i := range data.buses {
		id := data.buses[i]
		d := id
		for j := int64(2); d < data.arrival; j++ {
			d = j * id
		}
		if preferredID == 0 {
			preferredID = id
			departure = d

		} else if d <= departure {
			departure = d
			preferredID = id
		}
	}

	fmt.Println(preferredID, (departure-data.arrival)*preferredID)

}

func getInputData() input {
	lines, _ := file.GetLines("../input.txt")

	arrival, _ := strconv.ParseInt(lines[0], 10, 64)

	tokens := strings.Split(lines[1], ",")

	buses := make([]int64, 0, len(tokens))
	for i := range tokens {
		if tokens[i] != "x" {
			tmp, _ := strconv.ParseInt(tokens[i], 10, 64)
			buses = append(buses, tmp)
		}
	}

	return input{arrival: arrival, buses: buses}
}
