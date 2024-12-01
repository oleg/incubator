package main

import (
	"bufio"
	"github.com/oleg/incubator/aoc2024/misc"
	"math"
	"sort"
)

func main() {
	reader := misc.MustOpen("day01/input.txt")
	scanner := bufio.NewScanner(reader)

	left := make([]int, 0)
	right := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		l, r := parseLine(text)
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)

	println(sumOfDiffs(left, right))
	println(similarityScore(left, right))
}

func similarityScore(left []int, right []int) int {
	counts := make(map[int]int)

	sum := 0
	for _, l := range left {
		c, ok := counts[l]
		if !ok {
			for _, r := range right {
				if l == r {
					c++
				}
			}
			counts[l] = c
		}
		sum += l * c
	}
	return sum
}

func sumOfDiffs(left []int, right []int) int {
	sum := 0.
	for i, l := range left {
		r := right[i]
		sum += math.Abs(float64(l - r))
	}
	return int(sum)
}

func parseLine(text string) (int, int) {
	return misc.MustAtoi(text[:5]), misc.MustAtoi(text[8:])
}
