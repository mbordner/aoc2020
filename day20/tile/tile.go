package tile

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array/bytes"
	"github.com/mbordner/aoc2020/common/array/ints"
	"github.com/pkg/errors"
	"math"
	"regexp"
	"sort"
	"strconv"
)

var (
	reHeader = regexp.MustCompile(`^Tile\s(\d+):$`)
)

type Type int

const (
	Corner Type = iota
	Outside
	Inside
)

type Side int

const (
	Unknown Side = iota
	Top
	Right
	Bottom
	Left
)

type Orientation struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

type Details struct {
	Outside    []int
	Inside     []int
	Alignments map[int]*Tile
	Type       Type
}

type Tile struct {
	id      int
	t       int
	r       int
	b       int
	l       int
	bitmap  [][]byte
	details Details
	lines   []string
}

func (t *Tile) GetDetails() Details {
	return t.details
}

func (t *Tile) GetID() int {
	return t.id
}

func (t *Tile) GetType() Type {
	return t.details.Type
}

func (t *Tile) GetSide(s Side) int {
	switch s {
	case Top:
		return t.t
	case Right:
		return t.r
	case Bottom:
		return t.b
	case Left:
		return t.l
	}
	return int(Unknown)
}

func (t *Tile) GetDimensions() (int, int) {
	return len(t.bitmap), len(t.bitmap[0])
}

func (t *Tile) GetAlignmentTile(side int) *Tile {
	if v, ok := t.details.Alignments[side]; ok {
		return v
	}
	return nil
}

func (t *Tile) updateBorderKeys() {
	top := make([]byte, 10, 10)
	right := make([]byte, 10, 10)
	bottom := make([]byte, 10, 10)
	left := make([]byte, 10, 10)

	for i, j := 0, len(t.bitmap)-1; i < len(t.bitmap[0]); i, j = i+1, j-1 {
		if t.bitmap[0][i] == '#' {
			top[i] = '1'
		} else {
			top[i] = '0'
		}
		if t.bitmap[len(t.bitmap)-1][j] == '#' {
			bottom[i] = '1'
		} else {
			bottom[i] = '0'
		}
		if t.bitmap[i][0] == '#' {
			left[i] = '1'
		} else {
			left[i] = '0'
		}
		if t.bitmap[j][len(t.bitmap[0])-1] == '#' {
			right[i] = '1'
		} else {
			right[i] = '0'
		}
	}

	genId := func(side []byte) int {
		bits := make([]byte, len(side)+len(side), len(side)+len(side))
		tmp, _ := strconv.ParseInt(string(side), 2, 32)
		t1 := int(tmp)
		reversed := bytes.Reverse(side)
		tmp, _ = strconv.ParseInt(string(reversed), 2, 32)
		t2 := int(tmp)
		if t1 > t2 {
			copy(bits, side)
			copy(bits[len(side):], reversed)
		} else {
			copy(bits, reversed)
			copy(bits[len(side):], side)
		}
		tmp, _ = strconv.ParseInt(string(bits), 2, 32)
		return int(tmp)
	}

	t.t = genId(top)
	t.r = genId(right)
	t.b = genId(bottom)
	t.l = genId(left)
}

func (t *Tile) String() string {
	return fmt.Sprintf("{ [%d]: %d, %d, %d, %d }", t.id, t.t, t.r, t.b, t.l)
}

func (t *Tile) Rotate() {
	t.bitmap = bytes.Rotate(t.bitmap)
	t.updateBorderKeys()
}

func (t *Tile) Flip(direction bytes.Direction) {
	t.bitmap = bytes.Flip(direction, t.bitmap)
	t.updateBorderKeys()
}

func (t *Tile) Transform(o Orientation) bool {
	sideCount := 0
	if o.Top > 0 {
		sideCount++
	}
	if o.Right > 0 {
		sideCount++
	}
	if o.Bottom > 0 {
		sideCount++
	}
	if o.Left > 0 {
		sideCount++
	}
	if sideCount <= 1 {
		panic(errors.New("can't transform to 1 side only"))
	}
	if sideCount == 2 {
		if o.Top > 0 && o.Bottom > 0 {
			panic(errors.New("can't transform with only 2 opposite sides"))
		}
		if o.Left > 0 && o.Right > 0 {
			panic(errors.New("can't transform with only 2 opposite sides"))
		}
	}
	for i := 0; i < 3; i++ {
		t.Reset()

		switch i {
		case 1:
			t.Flip(bytes.Horizontal)
		case 2:
			t.Flip(bytes.Vertical)
		}

		if t.checkOrientation(o) {
			return true
		}

		for j := 0; j < 3; j++ {
			t.Rotate()
			if t.checkOrientation(o) {
				return true
			}
		}
	}

	return false
}

func (t *Tile) checkOrientation(orientation Orientation) bool {
	if orientation.Top > 0 {
		if t.t != orientation.Top {
			return false
		}
	}
	if orientation.Right > 0 {
		if t.r != orientation.Right {
			return false
		}
	}
	if orientation.Bottom > 0 {
		if t.b != orientation.Bottom {
			return false
		}
	}
	if orientation.Left > 0 {
		if t.l != orientation.Left {
			return false
		}
	}
	return true
}

func (t *Tile) Reset() {
	t.constructBitmap()
}

func (t *Tile) constructBitmap() {
	t.bitmap = make([][]byte, len(t.lines), len(t.lines))
	for i := range t.lines {
		t.bitmap[i] = []byte(t.lines[i])
	}
	t.updateBorderKeys()
}

func (t *Tile) Print() {
	fmt.Printf("Tile %d:\n", t.id)
	for i := range t.bitmap {
		fmt.Println(string(t.bitmap[i]))
	}
}

func (t *Tile) RemoveBorders() [][]byte {
	l := len(t.bitmap) - 2
	b := make([][]byte, l, l)
	for i := range b {
		b[i] = make([]byte, l, l)
	}
	bytes.Copy2D(b, t.bitmap, 0, 0, 1, 1, l, l)
	return b
}

func NewTile(lines []string) *Tile {
	t := &Tile{}

	matches := reHeader.FindStringSubmatch(lines[0])
	id, _ := strconv.ParseInt(matches[1], 10, 32)
	t.id = int(id)

	t.details.Inside = make([]int, 0, 4)
	t.details.Outside = make([]int, 0, 2)
	t.details.Alignments = make(map[int]*Tile)

	t.lines = lines[1:]
	t.constructBitmap()
	return t
}

type Tiles struct {
	tiles      []*Tile
	tilesMap   map[int]*Tile
	alignments map[int][]int
	orders     []*Tile
	bitmap     [][]*Tile
}

func (ts *Tiles) analyzeAlignments() {
	ts.alignments = make(map[int][]int)

	for id, t := range ts.tilesMap {
		if _, ok := ts.alignments[t.t]; ok {
			ts.alignments[t.t] = append(ts.alignments[t.t], id)
		} else {
			ts.alignments[t.t] = []int{id}
		}
		if _, ok := ts.alignments[t.r]; ok {
			ts.alignments[t.r] = append(ts.alignments[t.r], id)
		} else {
			ts.alignments[t.r] = []int{id}
		}
		if _, ok := ts.alignments[t.b]; ok {
			ts.alignments[t.b] = append(ts.alignments[t.b], id)
		} else {
			ts.alignments[t.b] = []int{id}
		}
		if _, ok := ts.alignments[t.l]; ok {
			ts.alignments[t.l] = append(ts.alignments[t.l], id)
		} else {
			ts.alignments[t.l] = []int{id}
		}
	}

	for border, tiles := range ts.alignments {
		if len(tiles) == 2 {
			ts.tilesMap[tiles[0]].details.Inside = append(ts.tilesMap[tiles[0]].details.Inside, border)
			ts.tilesMap[tiles[1]].details.Inside = append(ts.tilesMap[tiles[1]].details.Inside, border)
			ts.tilesMap[tiles[0]].details.Alignments[border] = ts.tilesMap[tiles[1]]
			ts.tilesMap[tiles[1]].details.Alignments[border] = ts.tilesMap[tiles[0]]
		} else if len(tiles) == 1 {
			ts.tilesMap[tiles[0]].details.Outside = append(ts.tilesMap[tiles[0]].details.Outside, border)
		} else {
			fmt.Println("unexpected Alignments issue")
		}
	}

	ts.orders = make([]*Tile, 0, len(ts.tilesMap))
	for _, tile := range ts.tilesMap {
		ts.orders = append(ts.orders, tile)
	}

	sort.SliceStable(ts.orders, func(i, j int) bool {
		if len(ts.orders[i].details.Outside) < len(ts.orders[j].details.Outside) {
			return false
		}
		if len(ts.orders[i].details.Inside) < len(ts.orders[j].details.Inside) {
			return true
		}
		return true
	})
}

func (ts *Tiles) Add(lines []string) *Tile {
	t := NewTile(lines)
	ts.tilesMap[t.id] = t
	ts.tiles = append(ts.tiles, t)
	return t
}

func (ts *Tiles) SquareLength() int {
	sl := math.Sqrt(float64(len(ts.tiles)))
	return int(sl)
}

func (ts *Tiles) checkAlignments() bool {
	sqLen := ts.SquareLength()
	expectedCorners := 4
	expectedOutside := (sqLen - 1) * 4
	expectedOutsideNotCorners := expectedOutside - 4
	expectedInside := sqLen*sqLen - expectedOutside

	unexpected := make([]*Tile, 0, len(ts.tiles))
	corners := make([]*Tile, 0, len(ts.tiles))
	outside := make([]*Tile, 0, len(ts.tiles))
	inside := make([]*Tile, 0, len(ts.tiles))

	for _, t := range ts.tiles {
		if len(t.details.Outside) == 0 && len(t.details.Inside) == 4 {
			inside = append(inside, t)
			t.details.Type = Inside
		} else if len(t.details.Outside) == 1 && len(t.details.Inside) == 3 {
			outside = append(outside, t)
			t.details.Type = Outside
		} else if len(t.details.Outside) == 2 && len(t.details.Inside) == 2 {
			corners = append(corners, t)
			t.details.Type = Corner
		} else {
			unexpected = append(unexpected, t)
		}
	}

	if len(corners) != expectedCorners {
		return false
	}
	if len(outside) != expectedOutsideNotCorners {
		return false
	}
	if len(inside) != expectedInside {
		return false
	}

	return true
}

func (ts *Tiles) Print() {
	for i := range ts.tiles {
		if i > 0 {
			fmt.Println()
		}
		ts.tiles[i].Print()
	}
}

func (ts *Tiles) GetOrders() []*Tile {
	return ts.orders
}

func (ts *Tiles) arrangeBitmap() error {
	sqLen := ts.SquareLength()
	ts.bitmap = make([][]*Tile, sqLen, sqLen)

	if len(ts.orders) != sqLen*sqLen {
		return errors.New("invalid state")
	}

	for i := range ts.bitmap {
		ts.bitmap[i] = make([]*Tile, sqLen, sqLen)
	}

	firstCorner := ts.orders[0]
	var firstSide *Tile
	var alignmentSide int
	for side, tile := range firstCorner.details.Alignments {
		firstSide = tile
		alignmentSide = side
		break
	}

	if firstSide == nil {
		return errors.New("couldn't locate first side")
	}

	var otherSide int
	for s := 0; s < len(firstCorner.details.Inside); s++ {
		if firstCorner.details.Inside[s] != alignmentSide {
			otherSide = firstCorner.details.Inside[s]
			break
		}
	}

	if !firstCorner.Transform(Orientation{Right: alignmentSide, Bottom: otherSide}) {
		return errors.New("first corner doesn't align")
	}
	if !firstSide.Transform(Orientation{Left: alignmentSide, Top: firstSide.details.Outside[0]}) {
		return errors.New("first side doesn't align")
	}

	ts.bitmap[0][0] = firstCorner
	ts.bitmap[0][1] = firstSide

	if firstSide.t != firstSide.details.Outside[0] {
		return errors.New("first side doesn't align correctly")
	}

	if !ints.Contains(firstCorner.details.Outside, firstCorner.t) || !ints.Contains(firstCorner.details.Outside, firstCorner.l) {
		return errors.New("first corner doesn't align correctly")
	}

	for j := 0; j < sqLen; j++ {
		for i := 0; i < sqLen; i++ {
			if j == 0 && i <= 1 {
				continue // already did these
			}
			var refTile *Tile
			var refSide Side
			if i == 0 {
				refTile = ts.bitmap[j-1][0]
				refSide = Bottom
			} else {
				refTile = ts.bitmap[j][i-1]
				refSide = Right
			}
			if refTile == nil {
				return errors.New("couldn't find ref tile")
			}
			side := refTile.GetSide(refSide)
			alignmentTile := refTile.GetAlignmentTile(side)
			if alignmentTile == nil {
				return errors.New("couldn't find alignment tile")
			}
			var o Orientation
			switch refSide {
			case Bottom:
				o = Orientation{Top: side}
			case Left:
				o = Orientation{Right: side}
			case Top:
				o = Orientation{Bottom: side}
			case Right:
				o = Orientation{Left: side}
			}
			switch alignmentTile.GetType() {
			case Outside:
				if j == 0 {
					o.Top = alignmentTile.details.Outside[0]
				} else if j == sqLen-1 {
					o.Bottom = alignmentTile.details.Outside[0]
				} else if i == 0 {
					o.Left = alignmentTile.details.Outside[0]
				} else if i == sqLen-1 {
					o.Right = alignmentTile.details.Outside[0]
					o.Top = ts.bitmap[j-1][i].GetSide(Bottom)
				}
			case Inside:
				o.Top = ts.bitmap[j-1][i].GetSide(Bottom)
			case Corner:
				var otherSide int
				for s := 0; s < len(alignmentTile.details.Inside); s++ {
					if alignmentTile.details.Inside[s] != side {
						otherSide = alignmentTile.details.Inside[s]
						break
					}
				}
				if j == 0 {
					o.Bottom = otherSide
				} else if i == 0 {
					o.Right = otherSide
				} else {
					o.Top = otherSide
				}
			}
			if !alignmentTile.Transform(o) {
				return errors.New("couldn't align alignment tile")
			}
			ts.bitmap[j][i] = alignmentTile
		}
	}

	return nil
}

func (ts *Tiles) GetTile(id int) *Tile {
	if t, ok := ts.tilesMap[id]; ok {
		return t
	}
	return nil
}

func (ts *Tiles) GetBitmapArrangement() [][]*Tile {
	return ts.bitmap
}

func NewTiles(lines []string) *Tiles {
	ts := &Tiles{}
	ts.tilesMap = make(map[int]*Tile)
	ts.tiles = make([]*Tile, 0, 25)

	var tileLines []string
	for i := range lines {
		if len(tileLines) == 0 {
			tileLines = make([]string, 0, 11)
		}
		if lines[i] != "" {
			tileLines = append(tileLines, lines[i])
		} else {
			ts.Add(tileLines)
			tileLines = make([]string, 0, 11)
		}
		i++
	}
	ts.Add(tileLines)

	ts.analyzeAlignments()
	if !ts.checkAlignments() {
		return nil
	}

	err := ts.arrangeBitmap()
	if err != nil {
		return nil
	}

	return ts
}
