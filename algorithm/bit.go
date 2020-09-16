package algorithm

// IsPowerOfTwo 是否是2的倍数
func IsPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

// HowManyOne 计算在一个 32 位的整数的二进制表式中有多少个1
func HowManyOne(x int) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
