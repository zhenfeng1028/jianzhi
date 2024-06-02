package main

/*
	给你一个字符串s和一个字符规律p，请你来实现一个支持'.'和'*'的正则表达式匹配。

	'.'匹配任意单个字符
	'*'匹配零个或多个前面的那一个元素
*/

// 动态规划（leetcode第10题）
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' { // 保证每次出现字符'*'时，前面都匹配到有效的字符，因此第一个字符不可能为'*'
				f[i][j] = f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
