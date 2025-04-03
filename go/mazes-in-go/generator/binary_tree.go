package generator

import (
	"math/rand"
	"mazes/maze"
)

func BinaryTree(r *rand.Rand, grid *maze.Grid) {
	grid.EachCells(func(cell *maze.Cell) {
		if neighbor := chooseNeighbor(r, grid, cell); neighbor != nil {
			cell.Link(neighbor)
		}
	})
}

func chooseNeighbor(r *rand.Rand, grid *maze.Grid, cell *maze.Cell) *maze.Cell {
	north := grid.North(cell)
	east := grid.East(cell)

	if north != nil && east != nil {
		if r.Intn(2) == 0 {
			return north
		}
		return east
	}
	if north != nil {
		return north
	}
	if east != nil {
		return east
	}
	return nil
}
