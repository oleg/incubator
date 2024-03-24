package leet

import "testing"

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
