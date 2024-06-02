package main

import (
	"regexp"
)

// 输入一个字符串，判断该字符串是否表示数值（包括整数、小数、正负数、科学计数法等）

func isNumeric(s string) bool {
	re := regexp.MustCompile(`[+-]?[0-9]{0,}(\.?[0-9]{1,})?([Ee][+-]?[0-9]{1,})?`)
	return re.MatchString(s)
}
