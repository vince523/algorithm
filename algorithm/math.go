package algorithm

// IsSquare 判断一个自然数是否是某个数的平方
func IsSquare(n int) bool {
	if n <= 0 {
		return false
	}

	var m int
	for i := 1; i < n; i++ {
		m = i * i
		if m == n {
			return true
		} else if m > n {
			return false
		}
	}
	return false
}

// Is2NPower 判断一个数是否为2的n次方
func Is2NPower(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}

func Division(m, n int) (int, int) {
	var i int

	for m >= n {
		m = m - n
		i++
	}
	return i, m
}
