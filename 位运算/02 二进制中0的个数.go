package main

// 输入一个整数，输出该二进制表示中0的个数，负数用补码表示。

// 每次右移一位，判断是否是0即可。

func findZero(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 0 {
			count++
		}
		n = n >> 1
	}
	return count
}
