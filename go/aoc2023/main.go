package main

import (
	"fmt"
	"log"
	"os"
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
	}
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

func parseGame(str string) game {
	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
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

func day2p2(file string) string {
	return ""
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
