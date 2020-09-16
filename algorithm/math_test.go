package algorithm

import (
	"testing"
)

func TestIsSquare(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"不是自然数", args{-1}, false},
		{"不是", args{2}, false},
		{"是", args{4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSquare(tt.args.n); got != tt.want {
				t.Errorf("IsSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs2NPower(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, true},
		{"2", args{2}, true},
		{"8", args{8}, true},
		{"9", args{9}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is2NPower(tt.args.n); got != tt.want {
				t.Errorf("Is2NPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"m == n", args{1, 1}, 1, 0},
		{"m > n", args{2, 1}, 2, 0},
		{"m > n", args{10, 3}, 3, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Division(tt.args.m, tt.args.n)
			if got != tt.want {
				t.Errorf("Division() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Division() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIs2NPower1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is2NPower(tt.args.n); got != tt.want {
				t.Errorf("Is2NPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSquare1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSquare(tt.args.n); got != tt.want {
				t.Errorf("IsSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}