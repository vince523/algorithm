package algorithm

import (
	"github.com/BUGLAN/kit/compare"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)

	i, j := 0, len(s)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

// FullArrange 全排列
func FullArrange(s string) []string {
	res := make([]string, 0)
	fullArrange([]rune(s), 0, &res)
	return res
}

func fullArrange(runes []rune, start int, res *[]string) {
	if runes == nil {
		return
	}

	if start == len(runes)-1 {
		*res = append(*res, string(runes))
	} else {
		for i := start; i < len(runes); i++ {
			// 交换
			if i != start {
				runes[i], runes[start] = runes[start], runes[i]
			}
			fullArrange(runes, start+1, res)
			if i != start {
				runes[start], runes[i] = runes[i], runes[start]
			}

		}
	}
}

// CommonSubSeq 最长公共子序列长度 动态规划
func CommonSubSeq(s1, s2 string) int {
	// 定义dp矩阵
	dp := make([][]int, len(s1)+1)

	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = compare.MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// CommonSubString  公共最长子串长度
func CommonSubString(s1, s2 string) int {
	// 定义dp矩阵
	dp := make([][]int, len(s1)+1)

	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	var m int

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			m = compare.MaxInt(m, dp[i][j])
		}
	}
	return m
}
