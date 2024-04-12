package lear

import (
	"errors"
)

var ErrNotEnoughElements = errors.New("less then two elements")

func twoLargest(arr []int) (int, int, error) {
	if len(arr) < 2 {
		return 0, 0, ErrNotEnoughElements
	}

	m1, m2 := arr[0], arr[1]

	if m1 < m2 {
		m1, m2 = m2, m1
	}

	for _, v := range arr[2:] {
		if v > m1 {
			m2 = m1
			m1 = v
		} else if v > m2 {
			m2 = v
		}
	}

	return m1, m2, nil
}
