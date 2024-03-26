package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	day := os.Args[1]
	file := mustReadFile(os.Args[2])
	switch day {
	case "day1p1":
		fmt.Println(day1p1(file))
	case "day1p2":
		fmt.Println(day1p2(file))
	case "day2p1":
		fmt.Println(day2p1(file))
	case "day2p2":
		fmt.Println(day2p2(file))
	case "day3p1":
		fmt.Println(day3p1(file))
	case "day3p2":
		fmt.Println(day3p2(file))
	}
}

type mines [][]int

func (m *mines) mark1(i, j int) {
	mine := *m
	if i >= 0 && j >= 0 && i < len(mine) && j < len(mine[i]) {
		mine[i][j] = 1
	}
}
func (m *mines) mark2(i, j int) {
	mine := *m
	if i >= 0 && j >= 0 && i < len(mine) && j < len(mine[i]) {
		mine[i][j] = 2
	}
}

func day3p1(str string) string {
	nums := make([][]rune, 0)
	for _, s := range strings.Split(str, "\n") {
		nums = append(nums, []rune(s))
	}

	mine := mines{}
	for i := 0; i < len(nums); i++ {
		mine = append(mine, make([]int, len(nums[i])))
		for j := 0; j < len(nums[i]); j++ {
			if !unicode.IsDigit(nums[i][j]) && nums[i][j] != '.' {
				mine.mark1(i, j)
			}
		}
	}

	for i := 0; i < len(mine); i++ {
		for j := 0; j < len(mine[i]); j++ {
			if mine[i][j] == 1 {
				mine.mark2(i-1, j-1)
				mine.mark2(i-1, j)
				mine.mark2(i-1, j+1)

				mine.mark2(i, j-1)
				mine.mark2(i, j+1)

				mine.mark2(i+1, j-1)
				mine.mark2(i+1, j)
				mine.mark2(i+1, j+1)
			}
		}
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		numStr := ""
		started := false
		poisoned := false
		for j := 0; j < len(nums[i]); j++ {
			if unicode.IsDigit(nums[i][j]) {
				numStr += string(nums[i][j])
				started = true
				poisoned = poisoned || mine[i][j] > 0
			} else {
				if started && poisoned {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatal(err)
					}
					sum += num
				}
				numStr = ""
				started = false
				poisoned = false
			}
		}
		if started && poisoned {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}
		numStr = ""
		started = false
		poisoned = false
	}

	return fmt.Sprintf("%d", sum)
}

type nums [][]rune

func (n *nums) dig(i, j int) string {
	num := *n
	if i >= 0 && j >= 0 && i < len(num) && j < len(num[i]) {
		if unicode.IsDigit(num[i][j]) {
			return "1"
		}
	}
	return "0"
}

type cells [][]int

func (c *cells) ids(i, j int) []int {
	m := make(map[int]bool)

	m[c.d(i-1, j-1)] = true
	m[c.d(i-1, j)] = true
	m[c.d(i-1, j+1)] = true

	m[c.d(i, j-1)] = true
	m[c.d(i, j)] = true
	m[c.d(i, j+1)] = true

	m[c.d(i+1, j-1)] = true
	m[c.d(i+1, j)] = true
	m[c.d(i+1, j+1)] = true

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (c *cells) d(i, j int) int {
	cs := *c
	if i >= 0 && j >= 0 && i < len(cs) && j < len(cs[i]) {
		return cs[i][j]
	}
	return 0
}

func day3p2(str string) string {
	num := nums{}
	for _, s := range strings.Split(str, "\n") {
		num = append(num, []rune(s))
	}

	mapping := make(map[int]int)
	cs := cells{}
	id := 1
	digits := false
	digStr := ""
	for i := 0; i < len(num); i++ {
		cs = append(cs, make([]int, len(num[i])))

		for j := 0; j < len(num[i]); j++ {
			if unicode.IsDigit(num[i][j]) {
				cs[i][j] = id
				digStr += string(num[i][j])
				digits = true
			} else {
				cs[i][j] = 0
				if digits {
					num, err := strconv.Atoi(digStr)
					if err != nil {
						log.Fatal(err)
					}
					mapping[id] = num
					id++
				}
				digits = false
				digStr = ""
			}
		}
		if digits {
			num, err := strconv.Atoi(digStr)
			if err != nil {
				log.Fatal(err)
			}
			mapping[id] = num
			id++
		}
		digits = false
		digStr = ""
	}

	sum := 0
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num[i]); j++ {
			if num[i][j] == '*' {
				ids := cs.ids(i, j)
				if len(ids) == 3 {
					slices.Sort(ids)
					a := mapping[ids[1]]
					b := mapping[ids[2]]
					sum += a * b
				}
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

func day2p1(str string) string {
	sum := 0
	for _, s := range strings.Split(str, "\n") {
		g := parseGame(s)
		if g.satisfies(round{red: 12, green: 13, blue: 14}) {
			sum += g.id
		}
	}
	return fmt.Sprintf("%d", sum)
}

type round struct {
	red   int
	green int
	blue  int
}

type game struct {
	id     int
	rounds []round
}

func (g game) satisfies(max round) bool {
	for _, r := range g.rounds {
		if r.red > max.red || r.green > max.green || r.blue > max.blue {
			return false
		}
	}
	return true
}

func (g game) min() round {
	m := round{
		red:   0,
		green: 0,
		blue:  0,
	}
	for _, r := range g.rounds {
		if r.red > m.red {
			m.red = r.red
		}
		if r.green > m.green {
			m.green = r.green
		}
		if r.blue > m.blue {
			m.blue = r.blue
		}
	}
	return m
}

func parseGame(str string) game {
	gameStr, roundsStr, found := strings.Cut(str, ":")
	if !found {
		log.Fatal("no colon found")
	}
	id := parseId(gameStr)
	g := game{id: id}
	roundsSlice := strings.Split(roundsStr, ";")
	for _, roundStr := range roundsSlice {
		r := round{}
		colors := strings.Split(roundStr, ",")
		for _, colorStr := range colors {
			numAndName := strings.Split(strings.TrimSpace(colorStr), " ")
			num, err := strconv.Atoi(strings.TrimSpace(numAndName[0]))
			if err != nil {
				log.Fatal(err)
			}
			color := strings.TrimSpace(numAndName[1])
			switch color {
			case "blue":
				r.blue = num
			case "red":
				r.red = num
			case "green":
				r.green = num
			default:
				log.Fatalf("unknown color %s", color)
			}
		}
		g.rounds = append(g.rounds, r)
	}
	return g
}

func parseId(gameStr string) int {
	idStr := strings.TrimPrefix(gameStr, "Game ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func day2p2(str string) string {
	sum := 0
	for _, s := range strings.Split(str, "\n") {
		g := parseGame(s)
		r := g.min()
		sum += r.red * r.green * r.blue
	}
	return fmt.Sprintf("%d", sum)
}

func day1p1(str string) string {
	sum := 0
	for _, s := range strings.Split(str, "\n") {
		var f byte
		for i := 0; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				f = s[i]
				break
			}
		}

		var l byte
		for i := len(s) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(s[i])) {
				l = s[i]
				break
			}
		}
		str := fmt.Sprintf("%c%c", f, l)
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	return fmt.Sprintf("%d", sum)
}

func day1p2(str string) string {
	sum := 0
	for _, s := range strings.Split(str, "\n") {
		sum += numOf(s)
	}
	return fmt.Sprintf("%d", sum)
}

func numOf(str string) int {
	var nums []int
	for i := 0; i < len(str); i++ {
		if num := decode(str[i:]); num > 0 {
			nums = append(nums, num)
		}
	}
	return nums[0]*10 + nums[len(nums)-1]
}

func decode(str string) int {
	switch {
	case strings.HasPrefix(str, "1") || strings.HasPrefix(str, "one"):
		return 1
	case strings.HasPrefix(str, "2") || strings.HasPrefix(str, "two"):
		return 2
	case strings.HasPrefix(str, "3") || strings.HasPrefix(str, "three"):
		return 3
	case strings.HasPrefix(str, "4") || strings.HasPrefix(str, "four"):
		return 4
	case strings.HasPrefix(str, "5") || strings.HasPrefix(str, "five"):
		return 5
	case strings.HasPrefix(str, "6") || strings.HasPrefix(str, "six"):
		return 6
	case strings.HasPrefix(str, "7") || strings.HasPrefix(str, "seven"):
		return 7
	case strings.HasPrefix(str, "8") || strings.HasPrefix(str, "eight"):
		return 8
	case strings.HasPrefix(str, "9") || strings.HasPrefix(str, "nine"):
		return 9
	}
	return 0
}

func mustReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
