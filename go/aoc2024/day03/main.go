package main

import (
	"fmt"
	"github.com/oleg/incubator/aoc2024/misc"
	"regexp"
	"strings"
)

var mulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var doRe = regexp.MustCompile(`do\(\)(?:[^d]|d[^o]|do[^n]|don[^']|don'[^t]|don't[^(]|don't\([^)])*don't\(\)`)

func main() {
	content := misc.MustReadFileToString("day03/input.txt")

	println(calculateSum(content))
	println(calculateAccurateSum(content))
	println(calculateAccurateSumRe(content))
}

func calculateAccurateSum(content string) int {
	split := strings.Split(content, "don't()")
	sum := calculateSum(split[0])

	for _, s := range split[1:] {
		_, after, found := strings.Cut(s, "do()")
		if found {
			sum += calculateSum(after)
		}
	}
	return sum
}

func calculateAccurateSumRe(content string) int {
	content = "do()" + content
	sum := 0
	if do := doRe.FindAllStringSubmatch(content, -1); len(do) > 0 {
		for _, m := range do {
			s := strings.TrimPrefix(m[0], "do()")
			s = strings.TrimSuffix(s, "don't()")
			sum += calculateSum(s)
		}
	}
	return sum
}

func calculateSum(content string) int {
	sum := 0
	if mul := mulRe.FindAllStringSubmatch(content, -1); len(mul) > 0 {
		for _, m := range mul {
			if len(m) != 3 {
				panic(fmt.Sprintf("unexpected match: %v", m))
			}
			sum += misc.MustAtoi(m[1]) * misc.MustAtoi(m[2])
		}
	}
	return sum
}
