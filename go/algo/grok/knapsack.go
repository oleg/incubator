package grok

import (
	"fmt"
	"strconv"
	"strings"
)

//todo use uint?

type item struct {
	name   string
	weight int
	value  int
}

type subsack []*item

func (s subsack) value() int {
	sum := 0
	for _, v := range s {
		sum += v.value
	}
	return sum
}
func (s subsack) weight() int {
	sum := 0
	for _, v := range s {
		sum += v.weight
	}
	return sum
}

func (s subsack) String() string {
	var b strings.Builder
	v := 0
	for _, it := range s {
		b.WriteString(it.name)
		b.WriteString(" ")
		v += it.value
	}
	b.WriteString(strconv.FormatInt(int64(v), 10))
	return b.String()
}

//todo: weight steps steps should be gcd
func choose(capacity int, items []item) []string {
	if len(items) == 0 {
		return nil
	}

	table := make([][]subsack, len(items))
	for i := range table {
		table[i] = make([]subsack, capacity)
	}

	fmt.Println(table)
	//best for 1, best for two, best for thre

	return nil
}
