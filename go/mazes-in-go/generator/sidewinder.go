package generator

import (
	"math/rand"
	"mazes/maze"
)

func Sidewinder(r *rand.Rand, grid *maze.Grid) {
	grid.EachRow(func(n int, row []*maze.Cell) {
		var run []*maze.Cell
		for _, cell := range row {
			run = append(run, cell)

			shouldCloseOut := grid.East(cell) == nil ||
				(grid.North(cell) != nil && r.Intn(2) == 0)

			if shouldCloseOut {
				member := run[r.Intn(len(run))]
				north := grid.North(member)
				if north != nil {
					member.Link(north)
				}
				run = []*maze.Cell{}
			} else {
				cell.Link(grid.East(cell))
			}
		}

	})
}
