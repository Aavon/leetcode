package main

import (
	"fmt"
	"testing"
)

/*
* @lc app=leetcode.cn id=34 lang=golang
*
* [34] 在排序数组中查找元素的第一个和最后一个位置
 */

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left, right := 0, len(nums)-1
	hit := -1
	// 查找对应位置
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			left, right = left, mid-1
		}
		if nums[mid] < target {
			left, right = mid+1, right
		}
		if nums[mid] == target {
			hit = mid
			break
		}
	}

	if hit == -1 {
		return result
	}

	result[0], result[1] = hit, hit
	for r := hit + 1; r < len(nums); r++ {
		if nums[r] == target {
			result[1] = result[1] + 1
		} else {
			break
		}
	}
	for l := hit - 1; l >= 0; l-- {
		if nums[l] == target {
			result[0] = result[0] - 1
		} else {
			break
		}
	}
	return result
}

func Test_a(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
