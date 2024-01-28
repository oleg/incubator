package grok

import (
	"algo/assert"
	"testing"
)

func TestKnapsack(t *testing.T) {
	cases := map[string]struct {
		items    []item
		capacity int
		expected []string
	}{
		"empty": {},
		"one item, same size": {
			items:    []item{{"one", 10, 100}},
			capacity: 10,
			expected: []string{"one"},
		},
		"one item, bigger": {
			items:    []item{{"one", 10, 100}},
			capacity: 5,
			expected: nil,
		},
		"one item, smaller": {
			items:    []item{{"one", 10, 100}},
			capacity: 20,
			expected: []string{"one"},
		},
		"two items, one bigger": {
			items:    []item{{"one", 10, 100}, {"two", 20, 200}},
			capacity: 10,
			expected: []string{"one"},
		},
		"two items, same size different value": {
			items:    []item{{"one", 10, 100}, {"two", 10, 200}},
			capacity: 10,
			expected: []string{"two"},
		},
		"two items, both fit": {
			items:    []item{{"one", 5, 100}, {"two", 5, 200}},
			capacity: 10,
			expected: []string{"one", "two"},
		},
		"two items, both bigger": {
			items:    []item{{"one", 20, 100}, {"two", 30, 200}},
			capacity: 10,
			expected: nil,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {

			res := choose(c.capacity, c.items)

			assert.EqualSlice(t, res, c.expected)
		})
	}
}

func TestKnapsackExample1(t *testing.T) {
	guitar := item{name: "guitar", value: 1500, weight: 1}
	stereo := item{name: "stereo", value: 3000, weight: 4}
	laptop := item{name: "laptop", value: 2000, weight: 3}

	bestValue := choose(4, []item{guitar, stereo, laptop})

	assert.Equal(t, len(bestValue), 2)
	assert.EqualSlice(t, bestValue, []string{"laptop", "guitar"})
}
