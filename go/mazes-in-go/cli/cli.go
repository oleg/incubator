package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"mazes/generator"
	"mazes/maze"
	"mazes/render"
	"os"
	"time"
)

type Algo func(r *rand.Rand, grid *maze.Grid)

var algorithms = map[string]Algo{
	"binary":     generator.BinaryTree,
	"sidewinder": generator.Sidewinder,
}
var renderers = map[string]Algo{
	"ascii": renderAscii,
	"png":   renderPng,
}

func main() {
	height := flag.Int("height", 10, "height of the maze")
	width := flag.Int("width", 10, "width of the maze")
	algorithm := flag.String("algorithm", "binary", "algorithm to generate maze")
	renderFlag := flag.String("render", "ascii", "algorithm to generate maze")
	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	grid := maze.NewGrid(*height, *width)
	algorithms[*algorithm](r, grid)
	renderers[*renderFlag](r, grid)
}

// todo move it
func renderAscii(r *rand.Rand, grid *maze.Grid) {
	ascii := render.ToAscii(grid)
	fmt.Println(ascii)
}

// todo move it
// todo enable output to file
func renderPng(r *rand.Rand, grid *maze.Grid) {
	err := render.ToPng(grid, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
