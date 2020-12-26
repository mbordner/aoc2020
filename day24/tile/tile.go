package tile

type State int

const (
	White State = iota
	Black
)

var (
	tilesMap map[string]*Tile
	ts       *Tiles
)

func init() {
	tilesMap = make(map[string]*Tile)
}

type Tile struct {
	id       string
	state    State
	adjacent map[Direction]*Tile
}

func (t *Tile) GetID() string {
	return t.id
}

func (t *Tile) Toggle() {
	if t.state == White {
		t.state = Black
	} else {
		t.state = White
	}
}

func (t *Tile) GetState() State {
	return t.state
}

func (t *Tile) GetAdjacentStateTiles(s State) []*Tile {
	directions := []Direction{E, W, NE, SE, NW, SW}
	adjacentStateTiles := make([]*Tile, 0, len(directions))
	for _, d := range directions {
		tmp := t.GetTile([]Direction{d})
		if tmp.GetState() == s {
			adjacentStateTiles = append(adjacentStateTiles, tmp)
		}
	}
	return adjacentStateTiles
}

func (t *Tile) GetTile(dirs []Direction) *Tile {
	dirs = Simplify(dirs)
	if len(dirs) == 0 {
		return t
	}
	var next *Tile
	if n, ok := t.adjacent[dirs[0]]; ok {
		next = n
	} else {
		suffix := dirs[0].String()
		nid := t.id + suffix

		ndirs := Directions(Simplify(Parse(nid)))
		tmp := ndirs.String()
		if tmp == nid {
			next = NewTile(t.id + suffix)
		} else {
			next = ts.GetTile(ndirs)
		}

		t.adjacent[dirs[0]] = next
		next.adjacent[dirs[0].Opposite()] = t
	}
	return next.GetTile(dirs[1:])
}

func NewTile(id string) *Tile {
	t := &Tile{id: id}
	t.adjacent = make(map[Direction]*Tile)
	tilesMap[id] = t
	return t
}

type Tiles struct {
	ref *Tile
}

func (ts *Tiles) GetTile(dirs []Direction) *Tile {
	return ts.ref.GetTile(dirs)
}

func (ts *Tiles) GetTiles() []*Tile {
	tiles := make([]*Tile, 0, len(tilesMap))
	for _, t := range tilesMap {
		tiles = append(tiles, t)
	}
	return tiles
}

func NewTiles() *Tiles {
	if ts != nil {
		return ts
	}
	ts = &Tiles{}
	ts.ref = NewTile("")
	return ts
}
