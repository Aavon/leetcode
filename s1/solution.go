// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
)

func main() {
	s0 := ""
	fmt.Scan(&s0)
	fmt.Println(s0)
	fmt.Println(lengestSunString(s0))
}

// assfaggsjduuttqwqr
func lengestSunString(s string) int {
	l := len(s)
	hits := make(map[byte]int, l)
	// 之前出现的位置
	var before int
	var ok bool
	// 子串开始的位置
	var left int = -1
	// 子串长度
	var maxLen, curLen int
	for i := 0; i < l; i++ {
		c := s[i]
		if before, ok = hits[c]; ok && before > left {
			left = before
		}
		curLen = i - left
		if curLen > maxLen {
			maxLen = curLen
		}
		hits[c] = i
	}
	return maxLen
}
