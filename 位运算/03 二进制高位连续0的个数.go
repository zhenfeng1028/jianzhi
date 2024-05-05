package main

// 输入一个整数，输出该二进制中高位连续0的个数。

/*
	0x80000000的二进制是1000 0000 0000 0000 0000 0000 0000 0000，仅最高位为1，
	每次左移一位，与最高位为1的二进制进行&操作，可以判断高位连续为0的个数。
*/

func numberOfLeading0(n int) int {
	if n == 0 {
		return 32
	}
	count := 0
	mask := 0x80000000
	m := n & mask
	for m == 0 {
		count++
		n <<= 1
		m = n & mask
	}
	return count
}
