package main

import (
	"fmt"
	"github.com/oleg/incubator/aoc2024/misc"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day04/input.txt")
	m := New(content)
	println(m.CountXmas())
	println(m.CountMas())
}

type Matrix struct {
	data [][]rune
	h, w int
}

func New(content string) *Matrix {
	lines := strings.Split(content, "\n")
	matrix := make([][]rune, 0, len(lines))
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}
	m := Matrix{data: matrix, h: len(matrix), w: len(matrix[0])}
	return &m
}

func (m *Matrix) CountXmas() int {
	count := 0
	for i, row := range m.data {
		for j := range row {
			count += m.xmasCountAt(point{i, j})
		}
	}
	return count
}

func (m *Matrix) xmasCountAt(p point) int {
	if m.at(p) != 'X' && m.at(p) != 'S' {
		return 0
	}
	count := 0
	if m.match("XMAS", p, p.mv(R), p.mv(R, R), p.mv(R, R, R)) {
		count++
	}
	if m.match("SAMX", p, p.mv(R), p.mv(R, R), p.mv(R, R, R)) {
		count++
	}
	if m.match("XMAS", p, p.mv(D), p.mv(D, D), p.mv(D, D, D)) {
		count++
	}
	if m.match("SAMX", p, p.mv(D), p.mv(D, D), p.mv(D, D, D)) {
		count++
	}
	if m.match("XMAS", p, p.mv(D, R), p.mv(D, D, R, R), p.mv(D, D, D, R, R, R)) {
		count++
	}
	if m.match("SAMX", p, p.mv(D, R), p.mv(D, D, R, R), p.mv(D, D, D, R, R, R)) {
		count++
	}
	if m.match("XMAS", p, p.mv(D, L), p.mv(D, D, L, L), p.mv(D, D, D, L, L, L)) {
		count++
	}
	if m.match("SAMX", p, p.mv(D, L), p.mv(D, D, L, L), p.mv(D, D, D, L, L, L)) {
		count++
	}
	return count

}

func (m *Matrix) CountMas() int {
	count := 0
	for i, row := range m.data {
		for j := range row {
			count += m.masCountAt(point{i, j})
		}
	}
	return count
}

func (m *Matrix) masCountAt(p point) int {
	if m.at(p) != 'A' {
		return 0
	}
	if m.isMasOrSamAt(p.mv(U, L), p, p.mv(D, R)) &&
		m.isMasOrSamAt(p.mv(D, L), p, p.mv(U, R)) {
		return 1
	}
	return 0

}

func (m *Matrix) isMasOrSamAt(p1, p2, p3 point) bool {
	return m.match("MAS", p1, p2, p3) || m.match("SAM", p1, p2, p3)
}
func (m *Matrix) at(p point) rune {
	if p.i < 0 || p.i >= m.h || p.j < 0 || p.j >= m.w {
		return '.'
	}
	return m.data[p.i][p.j]
}

func (m *Matrix) match(s string, ps ...point) bool {
	if len(s) != len(ps) {
		panic(fmt.Sprintf("len(s) != len(ps): %v != %v", len(s), len(ps)))
	}
	for i, r := range s {
		if m.at(ps[i]) != r {
			return false
		}
	}
	return true
}

type point struct {
	i, j int
}

var R = point{0, 1}
var L = point{0, -1}
var D = point{1, 0}
var U = point{-1, 0}

func (p point) add(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
}
func (p point) mv(ps ...point) point {
	t := p
	for _, p2 := range ps {
		t = t.add(p2)
	}
	return t
}
