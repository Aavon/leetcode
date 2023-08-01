package main

import (
	"fmt"
	"testing"
)

func jump(nums []int) int {
	/**
	0. 一次遍历,记录不同步数可以走的最远距离
	1. 记录经过当前节点，可以跳跃的最远距离
	2. 当遍历超过已知的最远距离，增步数+1，说明需要额外再跳一次
	*/
	maxPosition := 0
	LastMaxPosition := 0 // 初始值
	step := 0
	l := len(nums)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 初始边界值判断
	if l == 1 {
		return 0
	}
	for i := 0; i < l; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == LastMaxPosition {
			LastMaxPosition = maxPosition
			step++
			if LastMaxPosition >= l-1 {
				break //  已经找到最小步数
			}
		}
	}
	return step
}

func Test_a(t *testing.T) {
	fmt.Println(jump([]int{0}))
}
