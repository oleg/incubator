package main

import (
	"bufio"
	"github.com/oleg/incubator/aoc2024/misc"
	"slices"
	"strings"
)

func main() {
	reader := misc.MustOpen("day02/input.txt")
	scanner := bufio.NewScanner(reader)

	reports := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		report := parseLine(text)
		reports = append(reports, report)
	}

	println(numberOfSafeReports(reports))
	println(numberOfAlmostSafeReports(reports))
}

func numberOfSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func isSafe(report []int) bool {
	increasing := report[0] < report[1]

	for i := range report[1:] {
		if increasing && !safeIncrease(report[i], report[i+1]) {
			return false
		}
		if !increasing && !safeDecrease(report[i], report[i+1]) {
			return false
		}
	}
	return true
}

func numberOfAlmostSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		increasing := isIncreasing(report)
		if isAlmostSafe(increasing, report, false) {
			count++
		}
	}
	return count
}

func isIncreasing(report []int) bool {
	up := 0
	down := 0
	for i := range report[1:] {
		if report[i] < report[i+1] {
			up++
		}
		if report[i] > report[i+1] {
			down++
		}
	}
	return up > down
}

func isAlmostSafe(increasing bool, report []int, modified bool) bool {
	for i := 0; i < len(report)-1; i++ {
		if increasing {
			if safeIncrease(report[i], report[i+1]) {
				continue
			}
			if modified {
				return false
			}
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			return isAlmostSafe(increasing, slices.Delete(report, i, i+1), true) ||
				isAlmostSafe(increasing, slices.Delete(reportCopy, i+1, i+2), true)
		} else {
			if safeDecrease(report[i], report[i+1]) {
				continue
			}
			if modified {
				return false
			}
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			return isAlmostSafe(increasing, slices.Delete(report, i, i+1), true) ||
				isAlmostSafe(increasing, slices.Delete(reportCopy, i+1, i+2), true)
		}
	}
	return true
}

func safeIncrease(a, b int) bool {
	l := b - a
	return l >= 1 && l <= 3
}

func safeDecrease(a, b int) bool {
	l := a - b
	return l >= 1 && l <= 3
}

func parseLine(text string) []int {
	numStrs := strings.Split(text, " ")
	numbers := make([]int, 0, len(numStrs))
	for _, numStr := range numStrs {
		numbers = append(numbers, misc.MustAtoi(numStr))
	}
	return numbers
}
