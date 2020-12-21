package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"sort"
	"strings"
)

var (
	reLine = regexp.MustCompile(`^(\d+):\s(.*)$`)
)

type rule struct {
	id     string
	value  *[]string
	chains *[][]string
}

func (r *rule) IsLeaf() bool {
	if r.chains == nil {
		return true
	}
	return false
}

func (r *rule) GetValues() []string {
	return *r.value
}

func (r *rule) GetChains() [][]string {
	return *r.chains
}

func newRule(id string, def string) *rule {
	r := rule{}
	r.id = id
	if def[0] == '"' && def[len(def)-1] == '"' {
		v := []string{def[1 : len(def)-1]}
		r.value = &v
	} else {
		tokens := strings.Split(def, " | ")
		cs := make([][]string, len(tokens), len(tokens))
		for i := range tokens {
			cs[i] = strings.Split(tokens[i], " ")
		}
		r.chains = &cs
	}

	return &r
}

type Grammar struct {
	rules         map[string]*rule
	comboMemo     map[string][]string
	sequencesMemo map[string][][]string
	optimized     bool
}

func (g *Grammar) getCombinations(subValues [][]string) []string {
	key := fmt.Sprintf("%v", subValues)
	if vs, ok := g.comboMemo[key]; ok {
		return vs
	}
	if len(subValues) == 1 {
		return subValues[0]
	}

	suffixes := g.getCombinations(subValues[1:])

	values := make([]string, 0, len(suffixes)*len(subValues[0]))

	for i := range subValues[0] {
		for j := range suffixes {
			values = append(values, subValues[0][i]+suffixes[j])
		}
	}

	g.comboMemo[key] = values

	return values
}

func (g *Grammar) GetValues(id string) []string {
	values := make([]string, 0, 10)
	if r, ok := g.rules[id]; ok {
		if r.IsLeaf() {
			values = append(values, r.GetValues()...)
		} else {
			if r.value != nil {
				values = append(values, *r.value...)
			}
			chains := r.GetChains()
			for _, chain := range chains {
				chainValues := make([][]string, len(chain), len(chain))
				for i, cid := range chain {
					chainValues[i] = g.GetValues(cid)
				}
				values = append(values, g.getCombinations(chainValues)...)
			}
		}
	}
	return values
}

func (g *Grammar) getMinLength(values [][]string) int {
	length := 0
	for _, vs := range values {
		minLength := len(vs[0])
		for _, v := range vs[1:] {
			if len(v) < minLength {
				minLength = len(v)
			}
		}
		length += minLength
	}

	return length
}

func (g *Grammar) dedup(v1 []string) []string {
	values := make([]string, 0, len(v1))
	vs := make(map[string]bool)
	for _, v := range v1 {
		vs[v] = true
	}
	for v := range vs {
		values = append(values, v)
	}
	return values
}

func (g *Grammar) intersectValues(v1, v2 []string) []string {
	vs := make(map[string]bool)
	values := make([]string, 0, len(v1)+len(v2))
	for _, v := range v1 {
		vs[v] = true
		values = append(values, v)
	}
	for _, v := range v2 {
		if _, ok := vs[v]; !ok {
			vs[v] = true
			values = append(values, v)
		}
	}
	return values
}

func (g *Grammar) getMinPrefix(values []string, value string) int {
	minLength := -1
	for _, v := range values {
		if strings.HasPrefix(value, v) {
			if minLength == -1 {
				minLength = len(v)
			} else if len(v) == len(value) {
				return len(value)
			} else if len(v) < minLength {
				minLength = len(v)
			}
		}
	}
	return minLength
}

func (g *Grammar) isValidFor(id, value string) (matched string) {
	if r, ok := g.rules[id]; ok {
		if r.IsLeaf() == false {
			matchedValue := ""
			// check if this rule has leaf values that satisfy before checking chains
			if r.value != nil {
				values := *r.value
				for _, v := range values {
					if strings.HasSuffix(value, v) {
						matchedValue = v // found match at end, so we want to return what we matched
						break
					}
				}
			}

			// check chains backwards
		nextchain:
			for ci := len(*r.chains) - 1; ci >= 0; ci-- {
				chainValue := value
				chain := (*r.chains)[ci]
				for cri := len(chain) - 1; cri >= 0; cri-- {
					if !(chain[cri] == id && cri == len(chain)-1) {
						if matchedValue != "" {
							return matchedValue
						}
						tmp := g.isValidFor(chain[cri], chainValue)
						if tmp == "" {
							matched = ""
							continue nextchain
						} else {
							matched = tmp + matched
						}
						if matched == value && cri > 0 {
							matched = ""
							continue nextchain
						}
						chainValue = chainValue[0 : len(chainValue)-len(tmp)]
					} else {
						// loop backwards here... we're taking advantage of the problem
						// since 8: 42 | 42 8
						// and 11: 42 31 | 42 11 31 ....   we can continue to go back and consume the 42s
						tmpMatched := ""

						for {
							tmp := ""
							values := *r.value
							for _, v := range values {
								if strings.HasSuffix(chainValue, v+tmpMatched) {
									tmp = v
									break
								}
							}
							if tmp == "" {
								break
							} else {
								tmpMatched = tmp + tmpMatched
							}
						}

						if tmpMatched == "" {
							matched = ""
							continue nextchain
						} else {
							matched = tmpMatched + matched
						}
						if matched == value {
							return
						}
						chainValue = chainValue[0 : len(chainValue)-len(tmpMatched)]
					}
				}
			}
		} else {
			if r.value != nil {
				values := *r.value
				for _, v := range values {
					if strings.HasSuffix(value, v) {
						return v // found match at end, so we want to return what we matched
					}
				}
			}
		}

	}
	return
}

func (g *Grammar) IsValidFor(id, value string) bool {
	if r, ok := g.rules[id]; ok {
		if r.IsLeaf() == false {
			if !g.optimized {
				panic(errors.New("tree must be optimized"))
			}
			return g.isValidFor(id, value) == value
		} else {
			values := g.GetValues(id)
			index := sort.SearchStrings(values, value)
			if index < len(values) && values[index] == value {
				return true
			}
		}
	}
	return false
}

func (g *Grammar) UpdateRule(id string, rule string) {
	g.rules[id] = newRule(id, rule)
	g.optimized = false
}

func (g *Grammar) GetRule(id string) *rule {
	if r, ok := g.rules[id]; ok {
		return r
	}
	return nil
}

func (g *Grammar) optimizeRule(r *rule) bool {
	optimized := false
	if r.IsLeaf() == false {
		remainingChains := make([][]string, 0, len(*r.chains))
		for _, chain := range *r.chains {
			allLeaves := true
			leaves := make([][]string, 0, 100)
			for _, cid := range chain {
				if cid == r.id || g.GetRule(cid).IsLeaf() == false {
					allLeaves = false
					remainingChains = append(remainingChains, chain)
					break
				} else {
					leaves = append(leaves, g.GetRule(cid).GetValues())
				}
			}
			if allLeaves {
				combos := g.getCombinations(leaves)
				var values []string
				if r.value != nil {
					values = *r.value
				} else {
					values = make([]string, 0, 100)
				}
				values = g.intersectValues(values, combos)
				sort.Strings(values)
				r.value = &values
				optimized = true
			}
		}
		if len(remainingChains) == 0 {
			r.chains = nil
		} else {
			r.chains = &remainingChains
		}
	}
	return optimized
}

func (g *Grammar) OptimizeTree() {
	count := 1
	for count > 0 {
		count = 0
		for _, r := range g.rules {
			if g.optimizeRule(r) {
				count++
			}
		}
	}
	leaves := 0
	for _, r := range g.rules {
		if r.IsLeaf() {
			leaves++
		}
	}
	fmt.Println("tree optimized> leaves: ", leaves, " total rules: ", len(g.rules))
	g.optimized = true
}

func NewGrammar(rules []string) *Grammar {
	g := Grammar{}
	g.optimized = false
	g.rules = make(map[string]*rule)
	g.comboMemo = make(map[string][]string)
	g.sequencesMemo = make(map[string][][]string)

	for i := range rules {
		matches := reLine.FindStringSubmatch(rules[i])
		g.rules[matches[1]] = newRule(matches[1], matches[2])
	}

	return &g
}
