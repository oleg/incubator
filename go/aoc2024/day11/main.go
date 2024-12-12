package main

import (
	"github.com/oleg/incubator/aoc2024/misc"
	"strconv"
	"strings"
)

func main() {
	content := misc.MustReadFileToString("day11/input.txt")
	var stones []int
	for _, s := range strings.Split(content, " ") {
		stones = append(stones, misc.MustAtoi(s))
	}
	for range 25 {
		var newStones []int
		for _, stone := range stones {
			children := blink(stone)
			newStones = append(newStones, children...)
		}
		stones = newStones
	}
	println(len(stones))
}

func blink(n int) []int {
	if n == 0 {
		return []int{1}
	}
	s := strconv.Itoa(n)
	if len(s) > 1 && len(s)%2 == 0 {
		return []int{misc.MustAtoi(s[:len(s)/2]), misc.MustAtoi(s[len(s)/2:])}
	}
	return []int{n * 2024}
}
