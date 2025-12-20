package hot100

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {
	// 先解析内层,然后在解析外层,可以使用栈
	// 先压入括号,然后碰到括号之后开始重复添加
	type pair struct {
		s string
		k int
	}

	// 定义栈
	stack := make([]pair, 0)
	res := ""
	num := 0

	for _, c := range s {
		if unicode.IsLetter(c) {
			res += string(c)
		} else if unicode.IsDigit(c) {
			num = num*10 + int(c)
		} else if c == '[' {
			stack = append(stack, pair{s: res, k: num})
			res = ""
			num = 0
		} else if c == ']' {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = p.s + strings.Repeat(res, p.k)
		}
	}

	return res

}
