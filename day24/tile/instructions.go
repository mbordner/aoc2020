package tile

import "bytes"

type Direction int
type Directions []Direction

// e, se, sw, w, nw, and ne -- note: no n or s
const (
	E Direction = iota
	SE
	SW
	W
	NW
	NE
)

func (d Direction) String() string {
	switch d {
	case E:
		return "e"
	case W:
		return "w"
	case NE:
		return "ne"
	case SE:
		return "se"
	case NW:
		return "nw"
	case SW:
		return "sw"
	}
	return ""
}

func (d Direction) Opposite() Direction {
	switch d {
	case E:
		return W
	case W:
		return E
	case NE:
		return SW
	case SE:
		return NW
	case NW:
		return SE
	case SW:
		return NE
	}
	return E
}

func (dirs Directions) String() string {
	var buf bytes.Buffer
	for _, d := range dirs {
		buf.WriteString(d.String())
	}
	return buf.String()
}

func Simplify(dirs []Direction) []Direction {
	counts := make(map[Direction]int)
	preferred := []Direction{E, W, NE, SE, NW, SW}
	for _, d := range preferred {
		counts[d] = 0
	}
	for i := range dirs {
		counts[dirs[i]]++
	}
	reduced := true
	for reduced {
		reduced = false
		for counts[NW] > 0 && counts[SE] > 0 {
			counts[NW]--
			counts[SE]--
			reduced = true
		}
		for counts[NE] > 0 && counts[SW] > 0 {
			counts[NE]--
			counts[SW]--
			reduced = true
		}
		for counts[NW] > 0 && counts[SW] > 0 {
			counts[NW]--
			counts[SW]--
			counts[W]++
			reduced = true
		}
		for counts[NE] > 0 && counts[SE] > 0 {
			counts[NE]--
			counts[SE]--
			counts[E]++
			reduced = true
		}
		for counts[E] > 0 && counts[W] > 0 {
			counts[E]--
			counts[W]--
			reduced = true
		}
		for counts[W] > 0 && counts[NE] > 0 {
			counts[W]--
			counts[NE]--
			counts[NW]++
			reduced = true
		}
		for counts[W] > 0 && counts[SE] > 0 {
			counts[W]--
			counts[SE]--
			counts[SW]++
			reduced = true
		}
		for counts[E] > 0 && counts[NW] > 0 {
			counts[E]--
			counts[NW]--
			counts[NE]++
			reduced = true
		}
		for counts[E] > 0 && counts[SW] > 0 {
			counts[E]--
			counts[SW]--
			counts[SE]++
			reduced = true
		}
	}

	simplified := make([]Direction, 0, len(dirs))
	for _, d := range preferred {
		for i := 0; i < counts[d]; i++ {
			simplified = append(simplified, d)
		}
	}
	return simplified
}

func Parse(s string) []Direction {
	ins := make([]Direction, 0, len(s))
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case 'e':
			ins = append(ins, E)
		case 's':
			switch bytes[i+1] {
			case 'e':
				ins = append(ins, SE)
			case 'w':
				ins = append(ins, SW)
			}
			i++
		case 'w':
			ins = append(ins, W)
		case 'n':
			switch bytes[i+1] {
			case 'e':
				ins = append(ins, NE)
			case 'w':
				ins = append(ins, NW)
			}
			i++
		}
	}

	return ins
}
