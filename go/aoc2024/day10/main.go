package main

import (
	"github.com/oleg/incubator/aoc2024/misc"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day10/input.txt")
	m := newMatrix(content)
	println(m.countTrails())
	m.distinctTrail = true
	println(m.countTrails())
}

type matrix struct {
	data          [][]int
	visited       [][]bool
	h, w          int
	distinctTrail bool
}

func newMatrix(content string) *matrix {
	lines := strings.Split(content, "\n")
	m := &matrix{
		data:    make([][]int, 0, len(lines)),
		visited: make([][]bool, 0, len(lines)),
	}
	for _, line := range lines {
		cells := make([]int, 0, len(line))
		for _, r := range line {
			cells = append(cells, misc.MustAtoi(string(r)))
		}
		m.data = append(m.data, cells)
		m.visited = append(m.visited, make([]bool, len(line)))
	}
	m.h = len(m.data)
	m.w = len(m.data[0])
	return m
}

func (m *matrix) countTrails() int {
	count := 0
	m.eachCell(func(i, j int) {
		if m.data[i][j] == 0 {
			count += m.copyMe().countTrailsNext(0, i, j)
		}
	})
	return count
}

func (m *matrix) countTrailsNext(prev, i, j int) int {
	return m.countTrailsFrom(prev, i+1, j) +
		m.countTrailsFrom(prev, i, j+1) +
		m.countTrailsFrom(prev, i-1, j) +
		m.countTrailsFrom(prev, i, j-1)
}

func (m *matrix) countTrailsFrom(prev, i, j int) int {
	ok, h := m.at(i, j)
	if !ok || h-1 != prev {
		return 0
	}
	if !m.distinctTrail {
		if m.visited[i][j] {
			return 0
		}
		m.visited[i][j] = true
	}
	if h == 9 {
		return 1
	}
	return m.countTrailsNext(h, i, j)
}

func (m *matrix) eachCell(f func(int, int)) {
	for i := 0; i < m.h; i++ {
		for j := 0; j < m.w; j++ {
			f(i, j)
		}
	}
}

func (m *matrix) at(i, j int) (bool, int) {
	if i < 0 || i >= m.h || j < 0 || j >= m.w {
		return false, 0
	}
	return true, m.data[i][j]
}

func (m *matrix) copyMe() *matrix {
	duplicateData := make([][]int, len(m.data))
	duplicateVisited := make([][]bool, len(m.visited))
	for i := range m.data {
		duplicateData[i] = misc.Copy(m.data[i])
		duplicateVisited[i] = misc.Copy(m.visited[i])
	}
	return &matrix{
		data:          duplicateData,
		visited:       duplicateVisited,
		h:             m.h,
		w:             m.w,
		distinctTrail: m.distinctTrail,
	}
}
