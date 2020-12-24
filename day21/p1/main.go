package main

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array/strings"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
)

var (
	reWords = regexp.MustCompile(`(\w+)`)
)

type inputLine struct {
	ingredients []string
	allergens   []string
}

func main() {
	input := getInput()

	possibleAllergens := make(map[string][]string)

	extras := make([]string, 0, 10)

	ingredients := make(map[string]int)

	for i := range input {
		for _, s := range input[i].ingredients {
			if _, ok := ingredients[s]; !ok {
				ingredients[s] = 0
			}
			ingredients[s]++
		}
		for _, a := range input[i].allergens {
			if possibleIngredients, ok := possibleAllergens[a]; ok {
				ti, te := strings.Intersect(possibleIngredients, input[i].ingredients)
				extras = strings.Union(extras, te)
				possibleAllergens[a] = ti
			} else {
				possibleAllergens[a] = input[i].ingredients
			}
		}
	}

	allergens := make(map[string]string)

	for len(possibleAllergens) > 0 {
		for k, v := range possibleAllergens {
			if len(v) == 1 {
				allergens[v[0]] = k
				delete(possibleAllergens, k)
				extras = strings.Remove(extras, v[0])
			} else {
				for i := range allergens {
					possibleAllergens[k] = strings.Remove(v, i)
				}
			}
		}
	}

	fmt.Println(allergens, extras, ingredients)

	count := 0
	for _, s := range extras {
		count += ingredients[s]
	}

	fmt.Println(count)
}

func getInput() []inputLine {
	lines, _ := file.GetLines("../input.txt")
	input := make([]inputLine, len(lines), len(lines))
	for i := range lines {
		matches := reWords.FindAllStringSubmatch(lines[i], -1)
		if len(matches) > 0 {
			split := 0
			tmp := make([]string, 0, len(matches))
			for j := range matches {
				if matches[j][0] == "contains" {
					split = j
				} else {
					tmp = append(tmp, matches[j][0])
				}
			}
			input[i] = inputLine{ingredients: tmp[0:split], allergens: tmp[split:]}
		}
	}
	return input
}
