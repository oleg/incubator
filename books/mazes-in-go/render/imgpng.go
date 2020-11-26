package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"mazes/maze"
)

var pw = 3
var ww = 1

//var top color.Color = color.RGBA{R: 150, A: 255}
//var middle color.Color = color.RGBA{G: 150, A: 255}
//var bottom color.Color = color.RGBA{B: 150, A: 255}

var top color.Color =    color.Black
var middle color.Color = color.Black
var bottom color.Color = color.Black

//todo refactor
//todo pass options?
func ToPng(grid *maze.Grid, wr io.Writer) error {
	w := grid.Width*(pw+ww) + ww
	h := grid.Height*(pw+ww) + ww
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	grid.EachRow(func(y int, row []*maze.Cell) {
		if y == 0 {
			renderTop(grid, row, 0, img)
		}
		renderMiddle(grid, row, y*pw+y*ww+ww, img)
		renderBottom(grid, row, y*pw+pw+y*ww+ww, img)
	})

	err := png.Encode(wr, img)
	if err != nil {
		return err
	}
	return nil
}

func renderTop(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	for a := 0; a < ww; a++ {
		for b := 0; b < ww; b++ {
			img.Set(b, y+a, color.Black)
		}
	}
	for i, cell := range row {
		if !cell.Linked(grid.North(cell)) {
			for a := 0; a < ww; a++ {
				img.Set(pw*i+ww*i+ww+0, y+a, top)
				img.Set(pw*i+ww*i+ww+1, y+a, top)
				img.Set(pw*i+ww*i+ww+2, y+a, top)
			}
		}
		for a := 0; a < ww; a++ {
			for b := 0; b < ww; b++ {
				img.Set((i*pw+pw)+(i*ww+ww)+b, y+a, color.Black)
			}
		}
	}
}

func renderBottom(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	for a := 0; a < ww; a++ {
		for b := 0; b < ww; b++ {
			img.Set(b, y+a, color.Black)
		}
	}
	for i, cell := range row {
		if !cell.Linked(grid.South(cell)) {
			for a := 0; a < ww; a++ {
				img.Set(i*pw+i*ww+ww+0, y+a, bottom)
				img.Set(i*pw+i*ww+ww+1, y+a, bottom)
				img.Set(i*pw+i*ww+ww+2, y+a, bottom)
			}
		}
		for a := 0; a < ww; a++ {
			for b := 0; b < ww; b++ {
				img.Set((i*pw+pw)+(i*ww+ww)+b, y+a, color.Black)
			}
		}
	}
}
func renderMiddle(grid *maze.Grid, row []*maze.Cell, y int, img *image.RGBA) {
	for a := 0; a < ww; a++ {
		img.Set(a, y+0, middle)
		img.Set(a, y+1, middle)
		img.Set(a, y+2, middle)
	}
	for i, cell := range row {
		if !cell.Linked(grid.East(cell)) {
			for a := 0; a < ww; a++ {
				img.Set(i*pw+pw+i*ww+ww+a, y+0, middle)
				img.Set(i*pw+pw+i*ww+ww+a, y+1, middle)
				img.Set(i*pw+pw+i*ww+ww+a, y+2, middle)
			}
		}
	}

}
