package misc

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func MustOpen(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func MustReadFileToString(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func MustReadInts(reader io.Reader) []int {
	ints := make([]int, 0, 10)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		ints = append(ints, MustAtoi(scanner.Text()))
	}
	return ints
}

func MustAtoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func Copy[T any](slice []T) []T {
	aCopy := make([]T, len(slice))
	copy(aCopy, slice)
	return aCopy
}
