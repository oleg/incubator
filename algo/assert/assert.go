package assert

import (
	"testing"
)

func EqualSlice[T comparable](t *testing.T, actual, expected []T) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Errorf("got %#v, expected %#v", actual, expected)
		return
	}

	for i, a := range actual {
		Equal(t, a, expected[i])
	}
}

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %#v, expected %#v", actual, expected)
	}
}
