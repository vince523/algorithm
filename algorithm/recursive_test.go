package algorithm

import (
	"reflect"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"n = 1", args{1}, 1},
		{"n = 2", args{2}, 1},
		{"n = 5", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"n = 1", args{1}, 1},
		{"n = 2", args{2}, 2},
		{"n = 7", args{7}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hanota(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{2, 1, 0}, []int{}, []int{}}, []int{2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hanota(tt.args.A, tt.args.B, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hanota() = %v, want %v", got, tt.want)
			}
		})
	}
}
