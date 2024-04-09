package leet

import (
	"reflect"
	"slices"
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

func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"a", []int{1, 0, 2}, []int{1, 0, 0}},
		{"b", []int{1, 0, 0, 2}, []int{1, 0, 0, 0}},
		{"empty", []int{}, []int{}},
		{"zeros 1", []int{0}, []int{0}},
		{"zeros 2", []int{0, 0}, []int{0, 0}},
		{"zeros 3", []int{1, 0}, []int{1, 0}},
		{"zeros 4", []int{1, 0, 0}, []int{1, 0, 0}},
		{"zeros 5", []int{0, 0, 0}, []int{0, 0, 0}},
		{"zeros 6", []int{1, 0, 0, 0}, []int{1, 0, 0, 0}},
		{"no zeros", []int{1, 2, 3}, []int{1, 2, 3}},
		{"with zeros", []int{1, 0, 3, 4}, []int{1, 0, 0, 3}},
		{"with two zeros", []int{1, 0, 0, 4, 5, 6}, []int{1, 0, 0, 0, 0, 4}},
		{"with zeros at start", []int{0, 1, 2}, []int{0, 0, 1}},
		{"example 1", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"many zeros", []int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}},
		{"example 2", []int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.arr)
			if !slices.Equal(tt.arr, tt.want) {
				t.Errorf("duplicateZeros() = %v, want %v", tt.arr, tt.want)
			}

		})
	}
}
func Test_countZeros(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"a", []int{1, 0, 2}, 1},
		{"b", []int{1, 0, 0, 2}, 2},
		{"c", []int{1, 2, 0, 0}, 1},
		{"empty", []int{}, 0},
		{"zeros 1", []int{0}, 1},
		{"zeros 2", []int{0, 0}, 1},
		{"zeros 3", []int{1, 0}, 1},
		{"zeros 4", []int{1, 0, 0}, 1},
		{"zeros 5", []int{0, 0, 0}, 2},
		{"zeros 6", []int{1, 0, 0, 0}, 2},
		{"no zeros", []int{1, 2, 3}, 0},
		{"with zeros", []int{1, 0, 3, 4}, 1},
		{"with two zeros", []int{1, 0, 0, 4, 5, 6}, 2},
		{"with zeros at start", []int{0, 1, 2}, 1},
		{"example 1", []int{1, 0, 2, 3, 0, 4, 5, 0}, 2},
		{"many zeros", []int{0, 0, 0, 0, 0, 0, 0}, 4},
		{"example 2", []int{8, 4, 5, 0, 0, 0, 0, 7}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countZeros(tt.arr); got != tt.want {
				t.Errorf("countZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty",
			args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			"one element",
			args{[]int{1}, 1, []int{}, 0},
			[]int{1},
		},
		{
			"two elements",
			args{[]int{0}, 0, []int{1}, 1},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !slices.Equal(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	type want struct {
		answer  int
		content []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"example 1",
			args{nums: []int{3, 2, 2, 3}, val: 3},
			want{answer: 2, content: []int{2, 2}},
		},
		{
			"example 2",
			args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want{answer: 5, content: []int{0, 0, 1, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			content := tt.args.nums[0:got]
			slices.Sort(content)

			if got != tt.want.answer {
				t.Errorf("removeElement() = %v, want %v", got, tt.want.answer)
			}
			if !slices.Equal(content, tt.want.content) {
				t.Errorf("content %v, want %v", content, tt.want.content)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		answer  int
		content []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"example 1",
			args{nums: []int{1, 1, 2}},
			want{answer: 2, content: []int{1, 2}},
		},
		{
			"example 2",
			args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want{answer: 5, content: []int{0, 1, 2, 3, 4}},
		},
		{
			"example 3",
			args{nums: []int{1}},
			want{answer: 1, content: []int{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			content := tt.args.nums[0:got]
			slices.Sort(content)

			if got != tt.want.answer {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want.answer)
			}
			if !slices.Equal(content, tt.want.content) {
				t.Errorf("content %v, want %v", content, tt.want.content)
			}
		})
	}
}

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example 1", args{arr: []int{10, 2, 5, 3}}, true},
		{"example 2", args{arr: []int{3, 1, 7, 11}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
