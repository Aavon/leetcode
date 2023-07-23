package d121

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start

// 可以买卖无数次
// MP[i][0] = MAX(MP[i-1][1] + a[i], MP[i-1][0])
// MP[i][1] = MAX(MP[i-1][1], MP[i-1][0] - a[i])
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profits := make([][2]int, 2)
	for i := 0; i < 2; i++ {
		profits[i] = [2]int{}
	}

	profits[0][0], profits[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		next := i % 2
		before := 1 - next
		if profits[before][1]+prices[i] > profits[before][0] {
			profits[next][0] = profits[before][1] + prices[i]
		} else {
			profits[next][0] = profits[before][0]
		}

		if profits[before][1] > profits[before][0]-prices[i] {
			profits[next][1] = profits[before][1]
		} else {
			profits[next][1] = profits[before][0] - prices[i]
		}
	}

	lastDay := (len(prices) - 1) % 2
	if profits[lastDay][0] > profits[lastDay][1] {
		return profits[lastDay][0]
	} else {
		return profits[lastDay][1]
	}
}

// 可以买卖N次
// MP[i][k][0] = MAX(MP[i-1][k][0], MP[i-1][k-1][1] + a[i])
// MP[i][k][1] = MAX(MP[i-1][k][1], MP[i-1][k-1][0] - a[i])
func maxProfitN(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	N := 2 + 1 // 买卖次数
	profits := make([][][2]int, 2)
	for i := 0; i < 2; i++ {
		profits[i] = make([][2]int, N)
		for j := 0; j < N; j++ {
			profits[i][j] = [2]int{}
		}
	}

	for k := 1; k < N; k++ {
		if k%2 > 0 {
			profits[0][k][1] = -prices[0] // 一手买入都是负收益
		}
	}
	result := 0
	for i := 1; i < len(prices); i++ {
		next := i % 2
		before := 1 - next
		for k := 0; k < N; k++ {
			if k == 0 {
				profits[next][k][0] = 0
			} else {
				if profits[before][k][0] > profits[before][k-1][1]+prices[i] {
					profits[next][k][0] = profits[before][k][0]
				} else {
					profits[next][k][0] = profits[before][k-1][1] + prices[i]
				}

				if profits[before][k][1] > profits[before][k-1][0]-prices[i] {
					profits[next][k][1] = profits[before][k][1]
				} else {
					profits[next][k][1] = profits[before][k-1][0] - prices[i]
				}
			}
			if result < profits[next][k][0] {
				result = profits[next][k][0]
			}
			if result < profits[next][k][1] {
				result = profits[next][k][1]
			}
		}
	}
	return result
}

func Test_A(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test_N(t *testing.T) {
	// N := 2
	//fmt.Println(maxProfitN([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfitN([]int{7, 1, 5, 3, 6, 4}))
}

func Test_Two(t *testing.T) {
	// N:=3
	fmt.Println(maxProfitN([]int{1, 2, 3, 4, 5}))
	//fmt.Println(maxProfitN([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
