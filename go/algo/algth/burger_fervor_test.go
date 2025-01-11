package main

import (
	"algo/assert"
	"testing"
)

func solve(m, n, t int) (int, int) {
	result := solveT(m, n, t)
	if result >= 0 {
		return result, 0
	}
	i := t - 1
	result = solveT(m, n, i)
	for result == -1 {
		i--
		result += solveT(m, n, i)
	}
	return result, t - i
}

func solveT(m, n, t int) int {
	if t == 0 {
		return 0
	}

	var first int
	if t >= m {
		first = solveT(m, n, t-m)
	} else {
		first = -1
	}

	var second int
	if t >= n {
		first = solveT(m, n, t-n)
	} else {
		second = -1
	}

	if first == -1 && second == -1 {
		return -1
	}

	if first > second {
		return first + 1
	} else {
		return second + 1
	}
}

func Test_Calc(t *testing.T) {
	type expected struct {
		bur  int
		beer int
	}
	tests := []struct {
		name     string
		m, n, t  int
		expected expected
	}{
		{"t1", 4, 9, 22, expected{3, 0}},
		{"t2", 4, 9, 54, expected{6, 0}},
		{"t3", 4, 9, 15, expected{1, 0}},
		{"t4", 4, 9, 36, expected{4, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bur, beer := solve(test.m, test.n, test.t)

			assert.Equal(t, bur, test.expected.bur)
			assert.Equal(t, beer, test.expected.beer)
		})
	}
}
