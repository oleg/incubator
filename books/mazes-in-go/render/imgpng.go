package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"mazes/maze"
)

var ww = 4

//todo refactor
//todo pass options?
func ToPng(grid *maze.Grid, wr io.Writer) error {
	w := grid.Width * 5
	h := grid.Height * 5
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	grid.EachRow(func(n int, row []*maze.Cell) {
		v := n * ww
		if n == 0 {
			renderTop(grid, row, v+0, img)
		}
		renderMiddle(grid, row, v, img)
		renderBottom(grid, row, v+ww, img)
	})

	err := png.Encode(wr, img)
	if err != nil {
		return err
	}
	return nil
}

func renderTop(grid *maze.Grid, row []*maze.Cell, n int, img *image.RGBA) {
	img.Set(0, n, color.Black)
	for i, cell := range row {
		y := i * ww
		if !cell.Linked(grid.North(cell)) {
			img.Set(y+1, n, color.Black)
			img.Set(y+2, n, color.Black)
			img.Set(y+3, n, color.Black)
		}
		img.Set(y+ww, n, color.Black)
	}
}

func renderBottom(grid *maze.Grid, row []*maze.Cell, n int, img *image.RGBA) {
	img.Set(0, n, color.Black)
	for i, cell := range row {
		y := i * ww
		if !cell.Linked(grid.South(cell)) {
			img.Set(y+1, n, color.Black)
			img.Set(y+2, n, color.Black)
			img.Set(y+3, n, color.Black)
		}
		img.Set(y+ww, n, color.Black)
	}
}
func renderMiddle(grid *maze.Grid, row []*maze.Cell, n int, img *image.RGBA) {
	img.Set(0, n+1, color.Black)
	img.Set(0, n+2, color.Black)
	img.Set(0, n+3, color.Black)

	for i, cell := range row {
		if !cell.Linked(grid.East(cell)) {
			v := i*ww + ww
			img.Set(v, n+1, color.Black)
			img.Set(v, n+2, color.Black)
			img.Set(v, n+3, color.Black)
		}
	}

}
