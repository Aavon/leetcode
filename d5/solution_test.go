package d5

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	/**
	状态转移：当前串回文，字串（去掉两头）也是回文
	dp[i, j] = dp[i+1, j-1] && s[i] == s[j]
	**/
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	max := 0
	max_i, max_j := 0, 0
	// i 降序遍历、 j 增序遍历
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
				fmt.Println(i, j, dp[i][j])
			}
			if dp[i][j] && j-i > max {
				fmt.Println(i, j, "set")
				max_i, max_j = i, j
				max = j - i
			}
		}
	}
	return s[max_i : max_j+1]
}

func Test_a(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}
