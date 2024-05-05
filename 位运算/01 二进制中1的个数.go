package main

// 输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。

// n&(n-1)操作相当于把二进制表示中最右边的1变成0。

func numberOf1(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
