package main

import (
	"bufio"
	"github.com/oleg/incubator/aoc2024/misc"
	"io"
	"regexp"
)

func main() {
	reader := misc.MustOpen("day13/input.txt")
	data := parseInput(reader)
	println(calculate(data))
}

type game struct {
	a, b  button
	prize point
}
type button struct {
	x, y int
}

type point struct {
	x, y int
}

func calculate(games []*game) int {
	sum := 0
	for _, g := range games {
		sum += calculateOne(g)
	}
	return sum
}

func calculateOne(g *game) int {
	for a := range 100 {
		for b := range 100 {
			if optimal(a, b, g) {
				return a*3 + b*1
			}
		}
	}
	return 0
}

func optimal(aCount, bCount int, g *game) bool {
	x := g.prize.x - g.a.x*aCount - g.b.x*bCount
	y := g.prize.y - g.a.y*aCount - g.b.y*bCount
	if x == 0 && y == 0 {
		return true
	}
	return false
}

var aRe = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
var bRe = regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
var pRe = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

func parseInput(reader io.Reader) []*game {
	scanner := bufio.NewScanner(reader)
	games := make([]*game, 0)
	for scanner.Scan() {
		a := button{}
		a.x, a.y = getXY(scanner.Text(), aRe)

		scanner.Scan()
		b := button{}
		b.x, b.y = getXY(scanner.Text(), bRe)

		scanner.Scan()
		p := point{}
		p.x, p.y = getXY(scanner.Text(), pRe)

		scanner.Scan() //new line
		games = append(games, &game{a: a, b: b, prize: p})
	}
	return games
}

func getXY(str string, re *regexp.Regexp) (int, int) {
	match := re.FindAllStringSubmatch(str, -1)
	return misc.MustAtoi(match[0][1]), misc.MustAtoi(match[0][2])
}
