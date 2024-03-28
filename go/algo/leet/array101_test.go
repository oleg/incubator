package leet

import (
	"reflect"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{}}, 0},
		{"single one", args{[]int{0}}, 0},
		{"single zero", args{[]int{1}}, 1},
		{"ones", args{[]int{1, 1, 1, 1}}, 4},
		{"example 1", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"example 2", args{[]int{1, 0, 1, 1, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"example 1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example 2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty", []int{}, 0},
		{"example 1", []int{12, 345, 2, 6, 7896}, 2},
		{"example 2", []int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
