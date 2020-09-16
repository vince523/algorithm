package algorithm

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"空字符", args{s: ""}, ""},
		{"非空字符", args{s: "abc "}, " cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullArrange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"正常的测试", args{"abc"}, []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullArrange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonSubSeq(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"空字符串", args{"", ""}, 0},
		{"正常字符串", args{"abcde", "ace"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonSubSeq(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CommonSubSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonSubString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"abc", "bca"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonSubString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CommonSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}