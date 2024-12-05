package main

import (
	"bufio"
	"fmt"
	"github.com/oleg/incubator/aoc2024/misc"
	"slices"
	"strings"
)

func main() {
	reader := misc.MustOpen("day05/input.txt")
	scanner := bufio.NewScanner(reader)

	r, ir := parseRelations(scanner)
	ps := &PrintSystem{
		relation:        r,
		inverseRelation: ir,
		updates:         parseUpdates(scanner),
	}

	println(ps.checksumOfCorrectUpdates())
	println(ps.fixUpdatesAndGetChecksum())
}

type PrintSystem struct {
	relation        map[int][]int
	inverseRelation map[int][]int
	updates         [][]int
}

func (ps *PrintSystem) checksumOfCorrectUpdates() int {
	checksum := 0
	for _, update := range ps.updates {
		if ps.isCorrectUpdate(update) {
			checksum += update[len(update)/2]
		}
	}
	return checksum
}

func (ps *PrintSystem) isCorrectUpdate(update []int) bool {
	return slices.IsSortedFunc(update, ps.comparePages)
}

func (ps *PrintSystem) fixUpdatesAndGetChecksum() int {
	checksum := 0
	for _, update := range ps.updates {
		if !ps.isCorrectUpdate(update) {
			//updateCopy := misc.Copy(update)
			slices.SortFunc(update, ps.comparePages)
			checksum += update[len(update)/2]
		}
	}
	return checksum
}

func (ps *PrintSystem) comparePages(a, b int) int {
	if next, ok := ps.relation[a]; ok {
		if slices.Contains(next, b) {
			return -1
		}
	}
	if prev, ok := ps.inverseRelation[a]; ok {
		if slices.Contains(prev, b) {
			return 1
		}
	}
	return 0
}

func parseUpdates(scanner *bufio.Scanner) [][]int {
	var updates [][]int
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		strs := strings.Split(text, ",")
		update := make([]int, 0, len(strs))
		for _, s := range strs {
			update = append(update, misc.MustAtoi(s))
		}
		updates = append(updates, update)
	}
	return updates
}

func parseRelations(scanner *bufio.Scanner) (map[int][]int, map[int][]int) {
	relation := map[int][]int{}
	inverseRelation := map[int][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		b, a := parseRelation(text)
		relation[b] = append(relation[b], a)
		inverseRelation[a] = append(inverseRelation[a], b)
	}
	return relation, inverseRelation
}

func parseRelation(text string) (int, int) {
	before, after, found := strings.Cut(text, "|")
	if !found {
		panic(fmt.Sprintf("no | found in %s", text))
	}
	return misc.MustAtoi(before), misc.MustAtoi(after)
}
