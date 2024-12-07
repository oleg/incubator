package main

import (
	"github.com/oleg/incubator/aoc2024/misc"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day06/input.txt")
	m := NewMatrix(content)
	println(m.CountPositions())
	m = NewMatrix(content)
	println(m.CountObstacles())
}

type cell struct {
	visited     bool //X
	obstruction bool //#
}

type Matrix struct {
	data             [][]cell
	h, w             int
	currentPosition  point
	currentDirection direction
}

func NewMatrix(content string) *Matrix {
	lines := strings.Split(content, "\n")
	m := &Matrix{
		data: make([][]cell, 0, len(lines)),
	}
	for i, line := range lines {
		cells := make([]cell, 0, len(line))
		for j, r := range line {
			cells = append(cells, cell{obstruction: r == '#'})
			if r == '^' {
				m.currentPosition = point{i, j}
				m.currentDirection = Up
			}
		}
		m.data = append(m.data, cells)
	}
	m.h = len(m.data)
	m.w = len(m.data[0])
	return m
}

func (m *Matrix) CountPositions() int {
	m.moveUntilExit()
	return m.countVisited()
}
func (m *Matrix) CountObstacles() int {
	count := 0
	for i := 0; i < m.h; i++ {
		for j := 0; j < m.w; j++ {
			if i == m.currentPosition.i && j == m.currentPosition.j || m.data[i][j].obstruction {
				continue
			}
			m2 := m.copyMe()
			m2.data[i][j].obstruction = true
			if m2.canCycle() {
				count++
			}
		}
	}
	return count
}

func (m *Matrix) copyMe() *Matrix {
	duplicate := make([][]cell, len(m.data))
	for i := range m.data {
		duplicate[i] = make([]cell, len(m.data[i]))
		copy(duplicate[i], m.data[i])
	}
	return &Matrix{
		data:             duplicate,
		h:                m.h,
		w:                m.w,
		currentPosition:  m.currentPosition,
		currentDirection: m.currentDirection,
	}
}

func (m *Matrix) moveUntilExit() {
	d := Up
	for {
		for m.hasNextStep(d) {
			m.step(d)
		}
		if inbound, _ := m.at(m.currentPosition.move(d)); !inbound {
			return
		} else {
			d = d.turnRight()
		}
	}
}

func (m *Matrix) canCycle() bool {
	encounteredNew := false
	d := Up
	for {
		for m.hasNextStep(d) {
			if visited := m.stepAndIsVisitedBefore(d); !visited {
				encounteredNew = true
			}
		}
		if inbound, _ := m.at(m.currentPosition.move(d)); !inbound {
			return false
		}
		d = d.turnRight()
		if d == Up {
			if !encounteredNew {
				return true
			}
			encounteredNew = false
		}
	}
}

func (m *Matrix) stepAndIsVisitedBefore(d direction) (visited bool) {
	m.currentPosition = m.currentPosition.move(d)
	visited = m.data[m.currentPosition.i][m.currentPosition.j].visited
	m.data[m.currentPosition.i][m.currentPosition.j].visited = true
	return visited
}

func (m *Matrix) step(d direction) {
	m.currentPosition = m.currentPosition.move(d)
	m.data[m.currentPosition.i][m.currentPosition.j].visited = true
}

func (m *Matrix) hasNextStep(d direction) bool {
	inbound, obstructed := m.at(m.currentPosition.move(d))
	return inbound && !obstructed
}

func (m *Matrix) at(p point) (inbound, obstructed bool) {
	if p.i < 0 || p.i >= m.h || p.j < 0 || p.j >= m.w {
		return false, false
	}
	return true, m.data[p.i][p.j].obstruction
}

func (m *Matrix) inbound(p point) bool {
	return !(p.i < 0 || p.i >= m.h || p.j < 0 || p.j >= m.w)
}

func (m *Matrix) countVisited() int {
	count := 0
	for _, row := range m.data {
		for _, c := range row {
			if c.visited {
				count++
			}
		}
	}
	return count
}

type direction int

func (d direction) turnRight() direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("unexpected direction")
	}
}

const (
	Up direction = iota
	Right
	Left
	Down
)

type point struct {
	i, j int
}

func (p point) add(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
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
