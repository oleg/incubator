package generator

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"mazes/maze"
	"mazes/render"
	"testing"
)

func Test_produces_binary_tree(t *testing.T) {
	grid := maze.NewGrid(5, 5)

	r := rand.New(rand.NewSource(42))
	BinaryTree(r, grid)

	expected := dedent.Dedent(`
	+---+---+---+---+---+
	|                   |
	+---+---+   +   +   +
	|           |   |   |
	+---+---+---+   +   +
	|               |   |
	+   +---+---+---+   +
	|   |               |
	+---+   +   +---+   +
	|       |   |       |
	+---+---+---+---+---+`)
	assert.Equal(t, expected, render.ToAscii(grid))
}

func Test_returns_nil_if_north_and_east_are_nil(t *testing.T) {
	grid := maze.NewGrid(5, 5)
	cell := grid.Cell(0, 4)

	neighbor := chooseNeighbor(nil, grid, cell)

	assert.Nil(t, neighbor)
}

func Test_returns_north_if_east_is_nil(t *testing.T) {
	grid := maze.NewGrid(5, 5)
	cell := grid.Cell(1, 4)

	neighbor := chooseNeighbor(nil, grid, cell)

	assert.Equal(t, grid.Cell(0, 4), neighbor)
}

func Test_returns_east_if_north_is_nil(t *testing.T) {
	grid := maze.NewGrid(5, 5)
	cell := grid.Cell(0, 3)

	neighbor := chooseNeighbor(nil, grid, cell)

	assert.Equal(t, grid.Cell(0, 4), neighbor)
}

//todo hangs on compare assert.Equal(t, grid.Cell(3, 3), grid.Cell(2, 3))

func Test_returns_rand_if_east_and_north_are_not_nil(t *testing.T) {
	grid := maze.NewGrid(5, 5)
	cell := grid.Cell(3, 3)

	r := rand.New(rand.NewSource(42))
	neighbor := chooseNeighbor(r, grid, cell)

	assert.Equal(t, grid.Cell(3, 4), neighbor)
}
