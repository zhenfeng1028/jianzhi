package main

import "sort"

// 在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回-1（需要区分大小写）

func firstNotRepeatingChar(s string) int {
	m := make(map[rune][]int) // char <-> indices
	for i, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = make([]int, 0)
		}
		m[c] = append(m[c], i)
	}
	once := []int{} // 只出现一次的下标
	for _, v := range m {
		if len(v) == 1 {
			once = append(once, v[0])
		}
	}
	sort.Ints(once)
	if len(once) > 0 {
		return once[0]
	}
	return -1
}
