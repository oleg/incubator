package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"mazes/maze"
)

var pw = 4
var ww = 1

var red color.Color = color.RGBA{R: 255, A: 255}
var black color.Color = color.Black

//todo refactor
//todo pass options?
func ToPng(grid *maze.Grid, wr io.Writer) error {
	w := grid.Width * (pw + ww*2)
	h := grid.Height * (pw + ww*2)
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	grid.EachRow(func(y int, row []*maze.Cell) {
		if y == 0 {
			renderTop(grid, row, y*pw, img)
		}
		renderMiddle(grid, row, y*pw+(ww-1)+(ww-1)*y, img)
		renderBottom(grid, row, y*pw+pw+(ww-1)+(ww-1)*y, img)
	})

	err := png.Encode(wr, img)
	if err != nil {
		return err
	}
	return nil
}

func renderTop(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	for a := 0; a < ww; a++ {
		img.Set(0, y+a, black)
	}
	for i, cell := range row {
		if !cell.Linked(grid.North(cell)) {
			for a := 0; a < ww; a++ {
				img.Set(pw*i+1, y+a, black)
				img.Set(pw*i+2, y+a, black)
				img.Set(pw*i+3, y+a, black)
			}
		}
		for a := 0; a < ww; a++ {
			img.Set(i*pw+pw, y+a, black)
		}
	}
}

func renderBottom(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	for a := 0; a < ww; a++ {
		img.Set(0, y+a, black)
	}
	for i, cell := range row {
		if !cell.Linked(grid.South(cell)) {
			for a := 0; a < ww; a++ {
				img.Set(i*pw+1, y+a, black)
				img.Set(i*pw+2, y+a, black)
				img.Set(i*pw+3, y+a, black)
			}
		}
		for a := 0; a < ww; a++ {
			img.Set(i*pw+pw, y+a, black)
		}
	}
}
func renderMiddle(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	img.Set(0, y+1, red)
	img.Set(0, y+2, red)
	img.Set(0, y+3, red)

	for i, cell := range row {
		if !cell.Linked(grid.East(cell)) {
			img.Set(i*pw+pw, y+1, red)
			img.Set(i*pw+pw, y+2, red)
			img.Set(i*pw+pw, y+3, red)
		}
	}

}
