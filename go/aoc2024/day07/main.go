package main

import (
	"bufio"
	"github.com/oleg/incubator/aoc2024/misc"
	"strconv"
	"strings"
)

func main() {
	reader := misc.MustOpen("day07/input.txt")
	scanner := bufio.NewScanner(reader)

	equations := make([]equation, 0)
	for scanner.Scan() {
		e := parseLine(scanner.Text())
		equations = append(equations, e)
	}

	println(sum1(equations))
	println(sum2(equations))
}

type equation struct {
	arguments []int
	result    int
}

func sum1(equations []equation) int {
	sum := 0
	for _, e := range equations {
		if isCorrect1(e.result, e.arguments[0], e.arguments[1:]) {
			sum += e.result
		}
	}
	return sum
}
func sum2(equations []equation) int {
	sum := 0
	for _, e := range equations {
		if isCorrect2(e.result, e.arguments[0], e.arguments[1:]) {
			sum += e.result
		}
	}
	return sum
}

func isCorrect1(result, cur int, arguments []int) bool {
	if len(arguments) == 0 {
		return result == cur
	}
	return isCorrect1(result, cur+arguments[0], arguments[1:]) ||
		isCorrect1(result, cur*arguments[0], arguments[1:])
}

func isCorrect2(result, cur int, arguments []int) bool {
	if len(arguments) == 0 {
		return result == cur
	}
	return isCorrect2(result, cur+arguments[0], arguments[1:]) ||
		isCorrect2(result, cur*arguments[0], arguments[1:]) ||
		isCorrect2(result, concat(cur, arguments[0]), arguments[1:])
}

func concat(cur int, i int) int {
	return misc.MustAtoi(strconv.Itoa(cur) + strconv.Itoa(i))
}

func parseLine(text string) equation {
	before, after := misc.MustCut(text, ": ")
	strs := strings.Split(after, " ")

	e := equation{
		arguments: make([]int, 0, len(strs)),
		result:    misc.MustAtoi(before),
	}
	for _, s := range strs {
		e.arguments = append(e.arguments, misc.MustAtoi(s))
	}
	return e
}
