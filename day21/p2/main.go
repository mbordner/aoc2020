package main

import (
	"fmt"
	arrstrings "github.com/mbordner/aoc2020/common/array/strings"
	"github.com/mbordner/aoc2020/common/file"
	"regexp"
	"sort"
	"strings"
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
				ti, te := arrstrings.Intersect(possibleIngredients, input[i].ingredients)
				extras = arrstrings.Union(extras, te)
				possibleAllergens[a] = ti
			} else {
				possibleAllergens[a] = input[i].ingredients
			}
		}
	}

	knownAllergens := make(map[string]string)

	for len(possibleAllergens) > 0 {
		for k, v := range possibleAllergens {
			if len(v) == 1 {
				knownAllergens[v[0]] = k
				delete(possibleAllergens, k)
				extras = arrstrings.Remove(extras, v[0])
			} else {
				for i := range knownAllergens {
					possibleAllergens[k] = arrstrings.Remove(v, i)
				}
			}
		}
	}

	fmt.Println(knownAllergens, extras, ingredients)

	count := 0
	for _, s := range extras {
		count += ingredients[s]
	}

	fmt.Println(count)

	knownAllergies := make(map[string]string)
	allergies := make([]string, 0, len(knownAllergens))
	for k, v := range knownAllergens {
		knownAllergies[v] = k
		allergies = append(allergies, v)
	}

	sort.Strings(allergies)

	allergens := make([]string, 0, len(allergies))
	for _, a := range allergies {
		allergens = append(allergens, knownAllergies[a])
	}

	fmt.Println(strings.Join(allergens, ","))

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
