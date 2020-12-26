package tile_test

import (
	"github.com/mbordner/aoc2020/day24/tile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Simplify(t *testing.T) {
	dirs := tile.Parse("nesww")
	assert.Equal(t, 3, len(dirs))
	dirs = tile.Simplify(dirs)
	assert.Equal(t, 1, len(dirs))
	assert.Equal(t, "w", dirs[0].String())

	dirs = tile.Simplify(tile.Parse("nesew"))
	str := tile.Directions(dirs).String()
	assert.Equal(t, "", str)
}
