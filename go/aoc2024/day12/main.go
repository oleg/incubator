package main

import (
	"github.com/oleg/incubator/aoc2024/misc"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day12/input.txt")
	m := newMatrix(content)
	println(m.countPrice())

}

type matrix struct {
	data    map[point]rune
	visited misc.Set[point]
	max     point
}

func newMatrix(content string) *matrix {
	m := &matrix{data: make(map[point]rune), visited: misc.NewSet[point]()}
	var p point
	for i, line := range strings.Split(content, "\n") {
		for j, r := range line {
			p = point{i, j}
			m.data[p] = r
		}
	}
	m.max = p
	return m
}

type region struct {
	name      rune
	area      int
	perimeter int
}

func (m *matrix) countPrice() int {
	count := 0
	for p, name := range m.data {
		r := &region{name: name}
		m.discoverRegion(p, r)
		count += r.area * r.perimeter
	}
	return count
}

func (m *matrix) discoverRegion(p point, r *region) {
	if m.visited.Contains(p) {
		return
	}
	m.visited.Add(p)
	if m.data[p] == r.name {
		r.area += 1
		for _, d := range directions {
			if m.data[p.move(d)] != r.name {
				r.perimeter += 1
			} else {
				m.discoverRegion(p.move(d), r)
			}
		}
	}
}

func (m *matrix) eachCell(f func(p point, r rune)) {
	for p, r := range m.data {
		f(p, r)
	}
}

type direction int

var directions = []direction{Up, Left, Down, Right}

const (
	Up direction = iota
	Right
	Left
	Down
)

type point struct {
	i, j int
}

func (p point) move(d direction) point {
	switch d {
	case Up:
		return p.add(point{-1, 0})
	case Right:
		return p.add(point{0, 1})
	case Down:
		return p.add(point{1, 0})
	case Left:
		return p.add(point{0, -1})
	}
	return p
}

func (p point) add(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
}
