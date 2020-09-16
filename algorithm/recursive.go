package algorithm

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return n
	}
	return numWays(n-2) + numWays(n-1)
}

// htps://leetcode-cn.com/problems/hanota-lcci/
func hanota(A []int, B []int, C []int) []int {
	hanotaaa(len(A), &A, &B, &C)
	return C
}

func hanotaaa(n int, a, b, c *[]int) {
	if n == 1 {
		*c = append(*c, (*a)[len(*a)-1])
		*a = (*a)[:len(*a)-1]
	} else {
		hanotaaa(n-1, a, c, b)
		*c = append(*c, (*a)[len(*a)-1])
		*a = (*a)[:len(*a)-1]
		hanotaaa(n-1, b, a, c)
	}
}
