package generator

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"mazes/maze"
	"mazes/render"
	"testing"
)

func Test_produces_sidewinder(t *testing.T) {
	grid := maze.NewGrid(5, 5)

	r := rand.New(rand.NewSource(42))
	Sidewinder(r, grid)

	expected := dedent.Dedent(`
	+---+---+---+---+---+
	|                   |
	+   +---+   +---+---+
	|       |           |
	+   +---+---+---+   +
	|   |               |
	+   +---+   +   +---+
	|   |       |       |
	+   +   +   +   +   +
	|   |   |   |   |   |
	+---+---+---+---+---+`)
	assert.Equal(t, expected, render.ToAscii(grid))
}
