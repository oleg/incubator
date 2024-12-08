package main

import (
	"github.com/oleg/incubator/aoc2024/misc"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day08/input.txt")
	m := newMatrix(content)
	println(m.countAntinodes())
}

type matrix struct {
	data  [][]rune
	index map[rune][]point
	h, w  int
}

func newMatrix(content string) *matrix {
	lines := strings.Split(content, "\n")
	m := &matrix{
		data:  make([][]rune, 0, len(lines)),
		index: make(map[rune][]point),
	}
	for i, line := range lines {
		cells := make([]rune, 0, len(line))
		for j, r := range line {
			cells = append(cells, r)
			if r != '.' {
				m.index[r] = append(m.index[r], point{i, j})
			}
		}
		m.data = append(m.data, cells)
	}
	m.h = len(m.data)
	m.w = len(m.data[0])
	return m
}

func (m *matrix) countAntinodes() (int, int) {
	return len(m.antinodesAlg1()), len(m.antinodesAlg2())
}

func (m *matrix) antinodesAlg1() misc.Set[point] {
	s := misc.NewSet[point]()
	for _, points := range m.index {
		for _, p1 := range points {
			for _, p2 := range points {
				if p1 != p2 {
					s.Add(m.antinode1(p1, p2)...)
					s.Add(m.antinode1(p2, p1)...)
				}
			}
		}
	}
	return s
}

func (m *matrix) antinode1(p1 point, p2 point) []point {
	diff := p1.minus(p2)
	p := p1
	p = p.plus(diff)
	if m.isOnGrid(p) {
		return []point{p}
	}
	return nil
}

func (m *matrix) antinodesAlg2() misc.Set[point] {
	s := misc.NewSet[point]()
	for _, points := range m.index {
		for _, p1 := range points {
			for _, p2 := range points {
				if p1 != p2 {
					s.Add(p1, p2)
					s.AddSet(m.antinode2(p1, p2))
					s.AddSet(m.antinode2(p2, p1))
				}
			}
		}
	}
	return s
}
func (m *matrix) antinode2(p1 point, p2 point) misc.Set[point] {
	s := misc.NewSet[point]()
	diff := p1.minus(p2)
	p := p1
	for {
		p = p.plus(diff)
		if m.isOnGrid(p) {
			s.Add(p)
		} else {
			break
		}
	}
	return s
}

func (m *matrix) isOnGrid(p point) bool {
	return p.i >= 0 && p.i < m.h && p.j >= 0 && p.j < m.w
}

type point struct {
	i, j int
}

func (p point) plus(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
}

func (p point) minus(p2 point) point {
	return point{p.i - p2.i, p.j - p2.j}
}
