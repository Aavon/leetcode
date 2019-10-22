package d20

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.80%)
 * Likes:    1132
 * Dislikes: 0
 * Total Accepted:    141.5K
 * Total Submissions: 354.9K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
func isValid(s string) bool {
	md := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	l := len(s)
	if l % 2 == 1 {
		return false
	}
	if l == 0 {
		return true
	}
	stack := make([]byte, 0, l)
	for i:=0; i<l; i++ {
		ls := len(stack)
		if ls == 0 || md[stack[ls-1]] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:ls-1]
		}
	}
	return len(stack) == 0
}


func Test_a(t *testing.T) {
	str := "[{}]"
	fmt.Println(isValid(str))
	str = "()[]{}"
	fmt.Println(isValid(str))
	str = "(([]){})"
	fmt.Println(isValid(str))
}