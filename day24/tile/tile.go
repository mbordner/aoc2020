package tile

type State int

const (
	White State = iota
	Black
)

type Tile struct {
	state    State
	adjacent map[Direction]*Tile
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

func (t *Tile) GetTile(dirs []Direction) *Tile {
	if len(dirs) == 0 {
		return t
	}
	var next *Tile
	if n, ok := t.adjacent[dirs[0]]; ok {
		next = n
	} else {
		next = NewTile()
		t.adjacent[dirs[0]] = next
	}
	return next.GetTile(dirs[1:])
}

func NewTile() *Tile {
	t := &Tile{}
	t.adjacent = make(map[Direction]*Tile)
	return t
}

type Tiles struct {
	ref *Tile
}

func (ts *Tiles) GetTile(dirs []Direction) *Tile {
	return ts.ref.GetTile(dirs)
}

func NewTiles() *Tiles {
	ts := &Tiles{}
	ts.ref = NewTile()
	return ts
}
