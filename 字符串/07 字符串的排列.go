package main

// 输入一个字符串，按字典序打印出该字符串中字符的所有排列

var ans []string

func permute(s string) []string {
	permuteHelper(&s, 0, len(s)-1)
	return ans
}

func permuteHelper(s *string, l, r int) {
	if l == r {
		ans = append(ans, *s)
	} else {
		for i := l; i <= r; i++ {
			swap(s, l, i)
			permuteHelper(s, l+1, r)
			swap(s, l, i)
		}
	}
}

func swap(s *string, i, j int) {
	r := []rune(*s)
	r[i], r[j] = r[j], r[i]
	*s = string(r)
}
