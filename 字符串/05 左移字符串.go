package main

// 对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”，要求输出循环左移3位后的结果，即“XYZdefabc”。

func leftShiftString(s string, k int) string {
	doubleStr := s + s
	return doubleStr[k : k+len(s)]
}
