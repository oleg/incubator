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
	}

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
