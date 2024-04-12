package lear

import (
	"errors"
	"testing"
)

func Test_twoLargest(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		want1 int
		want2 int
		err   error
	}{
		{"empty", []int{}, 0, 0, ErrNotEnoughElements},
		{"one", []int{1}, 0, 0, ErrNotEnoughElements},
		{"two", []int{1, 2}, 2, 1, nil},
		{"three", []int{2, 3, 1}, 3, 2, nil},
		{"four", []int{-1, 5, 2, 1}, 5, 2, nil},
		{"same", []int{4, 3, 4, 2}, 4, 4, nil},
		{"ex1", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9, 6, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := twoLargest(tt.arr)
			if got1 != tt.want1 {
				t.Errorf("twoLargest() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("twoLargest() got2 = %v, want %v", got2, tt.want2)
			}
			if !errors.Is(err, tt.err) {
				t.Errorf("twoLargest() err = %v, want %v", err, tt.err)
			}
		})
	}
}
