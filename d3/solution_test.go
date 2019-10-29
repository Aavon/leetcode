package d3

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (31.68%)
 * Likes:    2568
 * Dislikes: 0
 * Total Accepted:    246.5K
 * Total Submissions: 776.1K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

func lengthOfLongestSubstring(s string) int {
	var maxLen, curLen int = 0, 0
	var left int = -1
	l := len(s)
	var c uint8
	exChars := make(map[uint8]int,l/2)
	for i:=0; i<l; i++{
		c = s[i]
		if before, ok := exChars[c]; ok {
			if before > left {
				left = before
			}
		}
		curLen = i - left
		if curLen > maxLen {
			maxLen = curLen
		}
		exChars[c] = i
	}
	return maxLen
}


func lengthOfLongestSubstring2(s string) int {
	var maxLen, curLen int = 0, 0
	var left int = -1
	exChars := make(map[int32]int, len(s)/2)
	// 居然会转成rune...
	for i, c := range s {
		if before, ok := exChars[c]; ok {
			if before > left {
				left = before
			}
		}
		curLen = i - left
		if curLen > maxLen {
			maxLen = curLen
		}
		exChars[c] = i
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int {
	var maxLen, curLen int = 0, 0
	var left int = -1
	bs := []byte(s)
	exChars := make(map[byte]int, len(bs)/2)
	for i, c := range bs {
		if before, ok := exChars[c]; ok {
			if before > left {
				left = before
			}
		}
		curLen = i - left
		if curLen > maxLen {
			maxLen = curLen
		}
		exChars[c] = i
	}
	return maxLen
}

func Test_a(t *testing.T) {
	str := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(str))
	str = "aab"
	fmt.Println(lengthOfLongestSubstring(str))
	str = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str))
}