package main

import (
	//"fmt"
	"github.com/oleg/incubator/aoc2024/misc"
)

func main() {
	content := misc.MustReadFileToString("day09/input.txt")
	data := parseData(content)
	println(part1(misc.Copy(data)), part2(data))

}

func parseData(content string) []int {
	var data []int
	id := 0
	for i, r := range content {
		n := misc.MustAtoi(string(r))
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				data = append(data, id)
			}
			id++
		} else {
			for j := 0; j < n; j++ {
				data = append(data, -1) //empty
			}
		}
	}
	return data
}

func part1(data []int) int {
	for i, j := 0, len(data)-1; i < j; {
		if data[i] != -1 {
			i++
			continue
		}
		if data[j] == -1 {
			j--
			continue
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}

	sum := 0
	for i := 0; i < len(data) && data[i] != -1; i++ {
		sum += data[i] * i
	}
	return sum
}

func part2(data []int) int {
	for i, j := 0, len(data)-1; i < j; {
		if data[j] == -1 {
			j--
			continue
		}
		jf := j
		v := data[j]
		for jf >= 0 && data[jf] == v {
			jf--
		}
		l := j - jf
		for i < j {
			if data[i] != -1 {
				i++
				continue
			}
			it := i
			for it-i < l && data[it] == -1 {
				it++
			}
			if it-i == j-jf {
				copy(data[i:it], data[jf+1:j+1])
				copy(data[jf+1:j+1], make([]int, j-jf))
				break
			} else {
				i++
				continue
			}
		}
		j = jf
		i = 0
	}
	for i := range data {
		if data[i] == -1 {
			data[i] = 0
		}
	}
	sum := 0
	for i := 0; i < len(data) && data[i] != -1; i++ {
		sum += data[i] * i
	}
	return sum
}
